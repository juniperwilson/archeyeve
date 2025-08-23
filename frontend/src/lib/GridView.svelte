<script lang="ts">
	import Polaroid from "./Polaroid.svelte";
    import type { Observation } from "$lib";

    let { observations }: { observations: Observation[] } = $props();
</script>

<div class="container">
    {#each observations as o} 
    <Polaroid observation={o}/>
    {/each}  
</div>

<style>
    .container {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: center;

        gap: 20px;
        max-width: 100%;
        max-height: 100%;
        padding-right: 10px;
    }
</style>