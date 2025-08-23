DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

CREATE EXTENSION IF NOT EXISTS ltree;


DROP TABLE IF EXISTS paths CASCADE;
DROP TABLE IF EXISTS styles CASCADE;
DROP TABLE IF EXISTS observations CASCADE;
DROP TABLE IF EXISTS users CASCADE;



CREATE TABLE IF NOT EXISTS users (
    id SERIAL UNIQUE PRIMARY KEY NOT NULL
);

CREATE TABLE IF NOT EXISTS observations (
    id SERIAL UNIQUE PRIMARY KEY NOT NULL,
    userid INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL,
    lng NUMERIC(9,6) NOT NULL,
    lat NUMERIC(9,6) NOT NULL,
    style_ids VARCHAR(100)[3],
    year NUMERIC(4),
    imgcount INTEGER NOT NULL
);


CREATE TABLE IF NOT EXISTS styles (
    id SERIAL UNIQUE PRIMARY KEY NOT NULL,
    name VARCHAR(100) UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS paths (
    id SERIAL UNIQUE PRIMARY KEY NOT NULL,
    style_id SERIAL NOT NULL references styles(id),
    path ltree not null,
    CONSTRAINT path_unique UNIQUE (path) DEFERRABLE INITIALLY IMMEDIATE
);








CREATE OR REPLACE FUNCTION add_edge(parent_id int, child_id int)
    /* Adds an edge between ``parent`` and ``child``.

    Designed to work with Directed Acyclic Graphs (DAG)
    (or said in another way, trees with multiple parents without cycles).

    This method will raise an exception if:
    - Parent is the same as child.
    - Child contains the parent.
    - Edge parent - child already exists.

    Influenced by:
    - https://www.codeproject.com/Articles/22824/A-Model-to-Represent-Directed-Acyclic-Graphs-DAG
    - http://patshaughnessy.net/2017/12/12/installing-the-postgres-ltree-extension
    - https://en.wikipedia.org/wiki/Directed_acyclic_graph
     */
    RETURNS void AS
$$
DECLARE
    parent text := CAST(parent_id as text);
    child  text := CAST(child_id as text);
BEGIN
    if parent = child
    then
        raise exception 'Cannot create edge: the parent is the same as the child.';
    end if;

    if EXISTS(
            SELECT 1
            FROM paths
            where paths.path ~ CAST('*.' || child || '.*.' || parent || '.*' as lquery)
        )
    then
        raise exception 'Cannot create edge: child already contains parent.';
    end if;

    -- We have two subgraphs: the parent subgraph that goes from the parent to the root,
    -- and the child subgraph, going from the child (which is the root of this subgraph)
    -- to all the leafs.
    -- We do the cartesian product from all the paths of the parent subgraph that end in the parent
    -- WITH all the paths that start from the child that end to its leafs.
    insert into paths (style_id, path) (
        select distinct style_id, fp.path || subpath(paths.path, index(paths.path, text2ltree(child)))
        from paths,
             (select paths.path from paths where paths.path ~ CAST('*.' || parent AS lquery)) as fp
        where paths.path ~ CAST('*.' || child || '.*' AS lquery)
    );
    -- Cleanup: old paths that start with the child (that where used above in the cartesian product)
    -- have became a subset of the result of the cartesian product, thus being redundant.
    delete from paths where paths.path ~ CAST(child || '.*' AS lquery);
END
$$
    LANGUAGE plpgsql;










CREATE OR REPLACE FUNCTION delete_edge(parent_id int, child_id int)
    /* Deletes an edge between ``parent`` and ``child``.

    Designed to work with DAG (See ``add_edge`` function).

    This method will raise an exception if the relationship does not
    exist.
     */
    RETURNS void AS
$$
DECLARE
    parent text := CAST(parent_id as text);
    child  text := CAST(child_id as text);
    number int;
BEGIN
    -- to delete we remove from the path of the descendants of the child
    -- (and the child) any ancestor coming from this edge.
    -- When we added the edge we did a cartesian product. When removing
    -- this part of the path we will have duplicate paths.

    -- don't check uniqueness for path key until we delete duplicates
    SET CONSTRAINTS path_unique DEFERRED;

    -- remove everything above the child lot_id in the path
    -- this creates duplicates on path and lot_id
    update path
    set paths = subpath(path, index(path, text2ltree(child)))
    where path ~ CAST('*.' || parent || '.' || child || '.*' AS lquery);

    -- remove duplicates
    -- we need an id field exclusively for this operation
    -- from https://wiki.postgresql.org/wiki/Deleting_duplicates
    DELETE
    FROM paths
    WHERE id IN (SELECT id
                 FROM (SELECT id, ROW_NUMBER() OVER (partition BY style_id, path) AS rnum
                       FROM paths) t
                 WHERE t.rnum > 1);

    -- re-activate uniqueness check and perform check
    -- todo we should put this in a kind-of finally clause
    SET CONSTRAINTS path_unique IMMEDIATE;

    -- After the update the one of the paths of the child will be
    -- containing only the child.
    -- This can only be when the child has no parent at all.
    -- In case the child has more than one parent, remove this path
    -- (note that we want it to remove it too from descendants of this
    -- child, ex. 'child_id'.'desc1')
    select COUNT(1) into number from paths where style_id = child_id;
    IF number > 1
    THEN
        delete from paths where path <@ text2ltree(child);
    end if;

END
$$
    LANGUAGE plpgsql;


