<script lang="ts">
	import type { Observation } from "$lib";
    import { BASE_PATH } from '.';
	import Polaroid from "./Polaroid.svelte";

    let { 
        observation,
        big, 
        mini,
    } : {
        observation: Observation,
        big?: boolean,
        mini?: boolean
    } = $props()

    // let alt = 'An observation of a ' + observation.condition + ' ' + observation.type + 
	// 	' building in the  ' + observation.styles[0] + 
	// 	' style found at ' + observation.address +'.'
    // alt = ""

</script>

{#snippet photo()}
    <a data-sveltekit-preload-data="tap" href="/observations/{observation.id}">
        <img src="{BASE_PATH}/{observation.id}/0.png" />
    </a>
{/snippet}

{#snippet description()}
    <a data-sveltekit-preload-data="tap" href="/observations/{observation.id}">
        <h1>{observation.name}</h1>
    </a>
{/snippet}

{#if big}
    <Polaroid {photo} {description} tall={false} big={true}/>
{:else if mini}
    <Polaroid {photo} {description} tall={false} mini={true}/>
{:else}
    <Polaroid {photo} {description} tall={false}/>
{/if}

