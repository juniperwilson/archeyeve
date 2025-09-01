<script lang="ts">
	import Polaroid from "./Polaroid.svelte";
    import type { Observation } from "$lib";
	import SmallPolaroid from "./SmallPolaroid.svelte";

    let { observations, small }: { observations: Observation[], small: boolean } = $props();
</script>

<div class="container">

    {#each observations as o} 
    {#if !small}
    <Polaroid observation={o}/>
    {:else}
    <SmallPolaroid observation={o}/>
    {/if}
    {/each}  
</div>

<style>
    .container {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: flex-end;
        gap: 20px;
        max-height: 100%;
        padding-right: 10px;
    }
</style>