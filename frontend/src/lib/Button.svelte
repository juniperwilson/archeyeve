<script lang="ts">
    import '@fontsource/roboto'
    	
	let { 
        text,
        onclick,
        isSelected
    } = $props()
</script>
    {#if isSelected}
	<button class="selected" {onclick}>{text}</button>
    {:else}
    <button {onclick}>{text}</button>
    {/if}

<style>
	button {
		font-size:1em;
        height: 100%;
        border-radius: 0px;
        border: 1px solid black;
        background-color: #f2eee2;
        font-family: "Roboto";
        font-weight: bold;
        text-wrap: nowrap;
	}
    button.selected {
        font-size:1em;
        height: 100%;
        border-radius: 0px;
        border: 1px solid black;
        background-color: #6dae47;
        color: white;
        font-family: "Roboto";
        font-weight: bold;
        text-wrap: nowrap;
    }
    button:hover {
        /* text-shadow: 1px 1px 1px grey; */
        background-color: #8bb174;
    }
    
    button:active {
        background-color: #6dae47;
        color: white;
    }


</style>