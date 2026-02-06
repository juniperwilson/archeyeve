<script lang="ts">
	import type { Observation } from "$lib";
    import { BASE_PATH } from '.';
	import Polaroid from "./Polaroid.svelte";

    let { 
        building,
        big = false, 
        mini = false,
    } : {
        building: Observation,
        big?: boolean,
        mini?: boolean
    } = $props()

    let src = BASE_PATH
    let alt = 'An observation of a ' + building.condition + ' ' + building.type + 
		' building in the  ' + building.styles[0] + 
		' style found at ' + building.address +'.'
    alt = ""

</script>

{#snippet photo()}
    <a data-sveltekit-preload-data="tap" href="/buildings/{building.id}">
        <img src="{BASE_PATH}/{building.id}/0.png" {alt} />
    </a>
{/snippet}

{#snippet description()}
    <a data-sveltekit-preload-data="tap" href="/buildings/{building.id}">
        <h1>{building.name}</h1>
    </a>
    <h2>
        {#each building.styles as style}
			<a data-sveltekit-preload-data="tap" href="/styles/{style}">{style}, </a>
		{/each}
        {building.year}
    </h2>
{/snippet}

{#if big}
    <Polaroid {photo} {description} big={true}/>
{:else if mini}
    <Polaroid {photo} {description} mini={true}/>
{:else}
    <Polaroid {photo} {description}/>
{/if}

