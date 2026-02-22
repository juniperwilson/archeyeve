<script lang="ts">
	import type { Observation } from "$lib";
	import Menu from "$lib/Menu.svelte";
	import type { PageData } from "../[obsid]/$types";
	import BuildingPolaroid from "$lib/BuildingPolaroid.svelte";
    import ObservationPolaroid from "$lib/ObservationPolaroid.svelte";
    import GridView from "$lib/GridView.svelte";

    let { data }: { data: PageData } = $props();
    let observation = $state(data.observation)
    let observations = [observation, observation, observation]
</script>

{#snippet polaroids()}
    {#each observations as observation}
        <ObservationPolaroid {observation} mini={true}/>
    {/each}
{/snippet}

<div class="pagecontainer">
        <Menu />
        <div class="content">
            <div class="info">
                <BuildingPolaroid building={observation} big={true}/>
                <div class="description">
                    <h1 class="underlined">{observation.name}</h1>
                    <h2>Year: {observation.year}</h2>
                    <h2>Styles: {observation.styles}</h2>
                    <h2>Type: {observation.type}</h2>
                    <h2>Condition: {observation.condition}</h2>
                    <h2>Address: {observation.address}</h2>
                    <br>
                    <h1 class="underlined">Observations</h1>
                    <div class="minipolaroidsbar">
                        <GridView {polaroids}></GridView>
                    </div>
                </div>
            </div>
            <div class="comments">
                <h1 class="underlined">Comments</h1>
            </div>
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
        flex-direction: column;
        gap: 30px;
        align-items: center;
        justify-content: center;
    }

    .info {
        display: flex;
        flex-direction: row;
        gap: 30px;
        /* width: 33vw; */

    }

    .underlined {
        border-bottom: solid 2px black;
    }

    .comments {
        width: 62vw;
        background-color: lavender;
    }

</style>