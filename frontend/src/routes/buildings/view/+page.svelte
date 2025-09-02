<script lang="ts">
	import type { Observation } from "$lib";
	import { doSearch, type Search } from "$lib/api";
	import GridView from "$lib/GridView.svelte";
	import Menu from "$lib/Menu.svelte";
	import SearchTab from "$lib/SearchTab.svelte";
    import type { PageData } from "../view/$types";

    let { data }: { data: PageData } = $props();
    let observations = data.observations

    async function advSearch(s: Search) {
        observations = await doSearch(s)
    }
</script>

<div class="pagecontainer">
        <Menu />
        <div class="content"> 
            <GridView tall={true} {observations} />
            <SearchTab search={advSearch}/>
        </div>
</div>

<style>
    .pagecontainer {
        display: flex;
        flex-direction: column;
        gap: 10px;
        overflow: hidden;
    }

    .content {
        /* display: grid;
        grid-template-columns: 75fr 25fr ;
        grid-template-rows: 1fr; */
        display: flex;
        flex-direction: row;
        justify-content: center;
        gap: 10px;
        height: 100%;
    }

</style>