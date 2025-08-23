<script lang="ts">
    import type { Observation } from '.'
    import { FILE_PATH } from '.';
    import '@fontsource/roboto'

    let {
        observation,
    }:{
        observation: Observation
    } = $props()

    let src = `${FILE_PATH}${observation.id}/1.jpeg`
    let alt = `${observation.style}`
</script>

<div class="gridcontainer">
    <div class="photo">
        <img {src} {alt}/>
        <!-- <p>{observation.imgcount}</p> -->
    </div>
    <div class="description">

        <h1>{observation.name}</h1>

        {#if !observation.year}
            <h2>{observation.style}</h2>
        {:else}
            <h2>{observation.style}, {observation.year}</h2>
        {/if}
    </div>
</div>

<style>

    .gridcontainer {
        display: grid;
        grid-template-rows: 79fr 21fr;
        aspect-ratio: 610 / 510;
        max-width: 20vw;
        max-height: 100%;
        background: #ffffff;
        border: 1px solid black;
        padding: 2%;
    }

    .description {
        max-height: 100%;
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
        font-family: "Roboto";
    }

    .photo {
        max-height: 100%;
    }

    img {
        width: 100%;
        aspect-ratio: 1;
        object-fit: cover;
    }
</style>