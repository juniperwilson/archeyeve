<script lang="ts">
	import { doSearch, type Search } from "$lib/api";
	import BuildingPolaroid from "$lib/BuildingPolaroid.svelte";
	import GridView from "$lib/GridView.svelte";
	import Menu from "$lib/Menu.svelte";
	import SearchTab from "$lib/SearchTab.svelte";
    import type { PageData } from "../view/$types";

    let { data }: { data: PageData } = $props();
    let buildings = $state(data.buildings)

    async function advSearch(s: Search) {
        buildings = await doSearch(s)
    }
</script>

{#snippet polaroids()}
    {#each buildings as building}
        <BuildingPolaroid {building}/>
    {/each}
{/snippet}

<div class="pagecontainer">
        <Menu />
        <div class="content"> 
            <GridView {polaroids} />
            <SearchTab search={advSearch} />
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