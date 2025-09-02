<script lang="ts">
	import { FILE_PATH, type Observation } from '.';
	import '@fontsource/roboto';
	import type { ObservationID } from './api';

	let {
		observation,
		big
	}: {
		observation: Observation,
		big: Boolean;
	} = $props();

	$inspect(observation)

	let src = "/static/0/1.jpeg"
	let alt = 'An observation of a ' + observation.condition + ' ' + observation.type + 
		' building in the  ' + observation.styles[0] + 
		' style found at ' + observation.address +'.';
</script>

<div class={["gridcontainer", { big }]}>
	<div class="photo">
		<a data-sveltekit-preload-data="tap" href="/buildings/{observation.id}">
			<img {src} {alt} />
		</a>
		<!-- <p>{observation.imgcount}</p> -->
	</div>
	<div class="description">
		<a data-sveltekit-preload-data="tap" href="/buildings/{observation.id}">
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
		/* grid-template-rows: 79fr 21fr;
		aspect-ratio: 1 / 1; */
		max-width: 20vw;
		min-width: 20vw;
		height: 24vw;
		display: block;
		/* max-height: 100%; */
		background: #ffffff;
		border: 1px solid black;
		padding: 2%;
	}

	.big {
		max-width: 30vw;
		min-width: 30vw;
		height: 36vw;
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

	.big h1 {
		font-size: 2em;
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

	.big h2 {
		font-size: 1.4em;
	}
</style>
