<script lang="ts">
    import GridView from "$lib/GridView.svelte";
	import Menu from "$lib/Menu.svelte";
	import SearchTab from "$lib/SearchTab.svelte";
    import { doSearch, type Search } from "$lib/api";
	import StylesBar from "$lib/SelectableButtonsBar.svelte";
	import type { PageData } from "./$types";
	import { selectedButton } from "$lib/shared.svelte";
	import Polaroid from "$lib/Polaroid.svelte";
	import type { Observation } from "$lib";
	import ObservationPolaroid from "$lib/ObservationPolaroid.svelte";

    const titles = ["manueline", "mudejar", "romanesque", "brutalism", "neo-mudejar", "neo-romanesque", "castro", "português suave"]
    let { data }: { data: PageData } = $props();
    let observations = $state(data.observations)

    async function styleSearch(style: string) {
        observations = await doSearch({style: style})
    }

    async function advSearch(s: Search) {
        observations = await doSearch(s)
    }
</script>

{#snippet polaroids()}
    {#each observations as observation}
        <ObservationPolaroid {observation}/>
    {/each}
{/snippet}

<div class="pagecontainer">
        <Menu />
        <StylesBar selected={selectedButton} {titles} onclick={styleSearch} />
    <div class="content"> 
        <GridView {polaroids} />
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
        display: flex;
        flex-direction: row;
        justify-content: center;
        gap: 10px;
        height: 100%;
    }
</style>