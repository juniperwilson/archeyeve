<script lang="ts">
	import { FILE_PATH, type Observation, Style} from '.';
	import '@fontsource/roboto';

	let {
        data,
        tall = true,
        big = false,
        mini = false
	} : {
        data: Observation | Style,
        tall: boolean,
        big: boolean,
        mini: boolean
    }= $props();

    let src = FILE_PATH

    let height = tall ? "tall" : "short"
    let size = mini ? "mini": (big ? "big": "normal")

    let name = observation ? observation.name : (style ? style.name : "")
    let info = style ? style.era : ( observation ? observation.name : "")
</script>

<div class={["container", {size}, {height}]}>
	<div class="photo">
		<a data-sveltekit-preload-data="tap" href="/observations/{observation.id}">
			<img {src} />
		</a>
	</div>
	<div class="description">
		<a data-sveltekit-preload-data="tap" href="/observations/{observation.id}">
			{@render description()}
		</a>
	</div>
</div>

<style>
    .container {
        display: grid;
        border: 1px solid black;
		padding: 2%;
        width: 20vw;
        height: 24vw;
        max-width: 20vw;
    }

    /* tall normal */
    .tall.normal {
        width: 20vw;
        height: 24vw;
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

    img .tall {
		width: 100%;
		aspect-ratio: 1;
		object-fit: cover;
	}

    .photo {
		max-height: 79%;
		height: 100%;
	}

    .description {
		max-height: 100%;
		height: 100%;
		/* max-width: 100%; */
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
