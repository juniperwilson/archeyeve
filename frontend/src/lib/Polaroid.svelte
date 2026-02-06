<script lang="ts">
	import type { Snippet } from 'svelte';
	import { BASE_PATH } from '.';
	import '@fontsource/roboto';

	let {
		photo,
		description,
		tall = true,
		big = false,
		mini = false
	}: {
		photo: Snippet;
		description: Snippet;
		tall?: boolean;
		big?: boolean;
		mini?: boolean;
	} = $props();

	let height = tall ? 'tall' : 'short';
	let size = mini ? 'mini' : big ? 'big' : 'normal';
</script>

<div class="polaroid-wrapper">
<div class="polaroid {size} {height}">
	<div class="photo">
		{@render photo()}
	</div>
	<div class="description">
		{@render description()}
	</div>
</div>
</div>

<style>
	.polaroid {
		display: grid;
		border: 1px solid black;
		padding: 6.5%;
		width: 20vw;
		height: 25vw;
		max-width: 20vw;
	}

	/* tall normal */
	.tall.normal {
		width: 20vw;
		height: 25vw;
		max-width: 20vw;
	}

	/* short normal */
	.short.normal {
		width: 15vw;
		height: 15vw;
	}

	/* tall big */
	.tall.big {
		width: 30vw;
		height: 36vw;
	}

	.short.big {
		width: 36vw;
		height: 36vw;
	}

	.short.mini {
		max-width: 12vw;
		max-height: 12vw;
	}

	.tall.mini {
		max-width: 10vw;
		max-height: 12vw;
	}

	:global(.polaroid img.tall) {
		width: 100%;
		aspect-ratio: 1;
		object-fit: cover;
	}

	.photo {
		max-height: 100%;
	}

    :global(.description a) {
        color: black;
        text-decoration: none;
    }

    :global(.description h2) {
        white-space: wrap;

    }
    

	:global(.polaroid.short .photo img) {
        width: 100%;
        object-fit: cover;
        aspect-ratio: 4/3;
	}

    :global(.polaroid.tall .photo img) {
        width: 100%;
        object-fit: cover;
        aspect-ratio: 1;
	}

	.description {
		max-height: 100%;
		height: 100%;
	}

	:global(.polaroid.tall h1) {
		font-size: 1.4em;
	}

	:global(.polaroid.short.normal h1) {
		font-size: 1.1em;
	}

	:global(.polaroid.big h1) {
		font-size: 2em;
	}

	:global(.polaroid.normal h2) {
		/* margin: 0%;
		margin-top: 2%; */
		font-size: 1em;
		line-height: 1;
		hyphens: auto;
		overflow: hidden;
		font-family: 'Roboto';
	}

	:global(.polaroid.big h2) {
		font-size: 1.4em;
	}
</style>
