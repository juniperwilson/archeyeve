<script lang="ts">
	import type { Style } from "$lib";
    import { FILE_PATH } from '.';
	import Polaroid from "./Polaroid.svelte";

    let { 
        style,
        big, 
        mini,
    } : {
        style: Style,
        big?: boolean,
        mini?: boolean
    } = $props()

    let src = FILE_PATH
    let alt = ""

</script>

{#snippet photo()}
    <a data-sveltekit-preload-data="tap" href="/styles/{style.name}">
        <img {src} {alt} />
    </a>
{/snippet}

{#snippet description()}
    <a data-sveltekit-preload-data="tap" href="/buildings/{style.name}">
        <h1>{style.name}</h1>
    </a>
    <h2>{style.era}</h2>
{/snippet}

{#if big}
    <Polaroid {photo} {description} big={true}/>
{:else if mini}
    <Polaroid {photo} {description} mini={true}/>
{:else}
    <Polaroid {photo} {description}/>
{/if}

