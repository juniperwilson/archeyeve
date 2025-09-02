<script lang="ts">
	import GridView from "$lib/GridView.svelte";
	import Menu from "$lib/Menu.svelte";
	import StylePolaroid from "$lib/StylePolaroid.svelte";
    import type { PageData } from "../view/$types";

    let { data }: { data: PageData } = $props();
    let styles = data.styles

</script>

{#snippet polaroids()}
    {#each styles as style}
        <StylePolaroid {style}/>
    {/each}
{/snippet}

<div class="pagecontainer">
        <Menu />
        <div class="content"> 
            <GridView {polaroids} />
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