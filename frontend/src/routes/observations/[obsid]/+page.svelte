<script lang="ts">
	import type { Observation } from "$lib";
	import Menu from "$lib/Menu.svelte";
	import type { PageData } from "../[obsid]/$types";
	import Polaroid from "$lib/Polaroid.svelte";
	import ObservationPolaroid from "$lib/ObservationPolaroid.svelte";

    let { data }: { data: PageData } = $props();
    let observation = $state(data.observation)
</script>

<div class="pagecontainer">
        <Menu />
        <div class="content">
            <ObservationPolaroid {observation} big={true}/>
            <div>
                <div class="description">
                    <h1>Year: {observation.year}</h1>
                    <h1>Styles: {observation.styles}</h1>
                    <h2>Type: {observation.type}</h2>
                    <h2>Condition: {observation.condition}</h2>
                    <h2>Address: {observation.address}</h2>
                </div>
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
        flex-direction: row;
        gap: 30px;
        height: 60vh;
        align-items: center;
        justify-content: center;
    }

    .description {
        height: 25vw;
        width: 25vw;
        border: 1px solid black;
		padding: 6.5%;
    }

</style>