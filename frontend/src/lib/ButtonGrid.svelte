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
        <Button isSquare={true} isSelected={selected.name == title? true : false } text={title.toUpperCase()} onclick={() => clickbutton(title)}></Button>
    {/each}
</div>
    
<style>
    .container {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        gap: 5px;
        width: 62vw;
    }
    
</style>