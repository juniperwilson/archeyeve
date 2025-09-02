<script lang="ts">
	import { FILE_PATH, type Observation } from '.';
	import '@fontsource/roboto';

	let {
		observation,
		big
	}: {
		observation: Observation;
		big: Boolean
	} = $props();

	$inspect(observation)

	let src = "/static/0/1.jpeg"
	let alt = 'An observation of a ' + observation.condition + ' ' + observation.type + 
		' building in the  ' + observation.styles[0] + 
		' style found at ' + observation.address +'.';
</script>

<div class={["gridcontainer", { big }]}>
	<div class="photo">
		<a data-sveltekit-preload-data="tap" href="/observations/{observation.id}">
			<img {src} {alt} />
		</a>
		<!-- <p>{observation.imgcount}</p> -->
	</div>
	<div class="description">
		<a data-sveltekit-preload-data="tap" href="/observations/{observation.id}">
			<h1>{observation.name}</h1>
		</a>
	</div>
</div>

<style>
	.gridcontainer {
		display: grid;
		/* grid-template-rows: 79fr 21fr;
		aspect-ratio: 610 / 510; */
		max-width: 15vw;
		min-width: 15vw;
		height: 15vw;
		display: block;
		max-height: 100%;
		background: #ffffff;
		border: 1px solid black;
		padding: 2%;
	}

	.big {
		max-width: 36vw;
		width: 36vw;
		aspect-ratio: 610 / 510;
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
		font-size: 1.1em;
		line-height: 1;
		/* word-break: break-all; */
		hyphens: auto;
		overflow: hidden;
	}

	.big h1 {
		font-size: 2em;
	}
</style>
