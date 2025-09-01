<script lang="ts">
	import type { Observation } from "$lib";
	import GridView from "$lib/GridView.svelte";
	import Menu from "$lib/Menu.svelte";
	import Polaroid from "$lib/Polaroid.svelte";
    import type { PageData } from "../view/$types";

    let { data }: { data: PageData } = $props();
    let observations = data.observations
</script>

<div class="pagecontainer">
        <Menu />
        <div class="content"> 
            <GridView small={false} {observations} />
            <!-- <SearchTab search={advSearch}/> -->
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