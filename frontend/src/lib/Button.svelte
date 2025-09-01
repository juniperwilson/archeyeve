<script lang="ts">
    import '@fontsource/roboto'
    	
	let { 
        text,
        onclick,
        isSelected,
        isSquare
    } = $props()
</script>

{#if isSelected}
	{#if isSquare}
		<button class="selected square" {onclick}>{text}</button>
	{:else}
		<button class="selected" {onclick}>{text}</button>
	{/if}
{:else if isSquare}
	<button class="square" {onclick}>{text}</button>
{:else}
	<button {onclick}>{text}</button>
{/if}

<style>
	button {
		font-size: 1em;
		height: 100%;
		border-radius: 0px;
		border: 1px solid black;
		background-color: #f2eee2;
		font-family: 'Roboto';
		font-weight: bold;
		text-wrap: nowrap;
	}
	button.selected {
		font-size: 1em;
		height: 100%;
		border-radius: 0px;
		border: 1px solid black;
		background-color: #6dae47;
		color: white;
		font-family: 'Roboto';
		font-weight: bold;
		text-wrap: nowrap;
	}

    button.square {
        aspect-ratio: 1 / 1;
        width: 10vw;
		max-width: 10vw;
		max-height: 10vw;
		background-color: white;
		font-size: 1em;
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
