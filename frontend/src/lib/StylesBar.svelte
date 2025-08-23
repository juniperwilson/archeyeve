<script lang="ts">
	import Button from "./Button.svelte";

    let { styles, styleSearch } = $props();
    let selected = $state(false)
</script>

<div class="container">
    {#each styles as style}
        <Button class={["primary lg", { selected }]} on:click={() => {styleSearch(style, selected); selected = !selected;}}>{style.toUpperCase()}</Button>
    {/each}
</div>
    
<style>
    .container {
        display: flex;
        flex-direction: row;
        /* width: 100%; */
        max-height: 100%;
        gap: 5px;
        overflow-x: auto;
    }
</style>