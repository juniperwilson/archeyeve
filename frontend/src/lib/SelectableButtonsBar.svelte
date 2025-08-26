<script lang="ts">
	import Button from "./Button.svelte";

    let { titles, onclick, selected = null } = $props();

    function clickbutton(title: string) {
        if (title == selected.name) {
            onclick("")
            selected.name = ""
        } else {
            onclick(title)
            selected.name = title
        }
    }
</script>

<div class="container">
    {#each titles as title }
        <Button isSelected={selected.name == title? true : false } text={title.toUpperCase()} onclick={() => clickbutton(title)}></Button>
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