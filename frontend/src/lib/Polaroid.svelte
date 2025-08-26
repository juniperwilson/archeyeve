<script lang="ts">
	import type { Observation } from '.';
	import '@fontsource/roboto';

	let {
		observation
	}: {
		observation: Observation;
	} = $props();

	$inspect(observation)

	// let src = `${FILE_PATH}${observation.id}/1.jpeg`
	let alt = 'An observation of a building in the  ' + observation.styles[0] + ' style found at ' + observation.address +'.';
</script>

<div class="gridcontainer">
	<div class="photo">
		<a data-sveltekit-preload-data="tap" href="/observations/{observation.id}">
			<img {alt} />
		</a>
		<!-- <p>{observation.imgcount}</p> -->
	</div>
	<div class="description">
		<a data-sveltekit-preload-data="tap" href="/observations/{observation.id}">
			<h1>{observation.name}</h1>
		</a>

		{#if !observation.year}
			<h2>
				{#each observation.styles as style}
					<a data-sveltekit-preload-data="tap" href="/styles/{style}">{style}</a>
				{/each}
			</h2>
		{:else}
			<h2>
				{#each observation.styles as style}
					<a data-sveltekit-preload-data="tap" href="/styles/{style}"> {style},</a>
				{/each}
				{observation.year}
			</h2>
		{/if}
	</div>
</div>

<style>
	.gridcontainer {
		display: grid;
		grid-template-rows: 79fr 21fr;
		aspect-ratio: 610 / 510;
		max-width: 20vw;
		min-width: 20vw;
		height: auto;
		display: block;
		max-height: 100%;
		background: #ffffff;
		border: 1px solid black;
		padding: 2%;
	}

	.photo {
		max-height: 79%;
		height: 100%;
	}

	img {
		width: 100%;
		aspect-ratio: 1;
		object-fit: cover;
	}

	.description {
		max-height: 100%;
		height: 100%;
		/* max-width: 100%; */
	}

	h1 {
		margin: 0%;
		margin-top: 2%;

		font-size: 1.4em;
		line-height: 1;
		/* word-break: break-all; */
		hyphens: auto;
		overflow: hidden;
	}

	h2 {
		margin: 0%;
		margin-top: 2%;
		font-size: 1em;
		line-height: 1;
		/* word-break: break-all; */
		hyphens: auto;
		overflow: hidden;
		font-family: 'Roboto';
	}
</style>
