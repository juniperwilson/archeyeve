<script lang="ts">
	import type { Style } from "$lib";
	import Menu from "$lib/Menu.svelte";
	import ObservationPolaroid from "$lib/ObservationPolaroid.svelte";
	import BuildingPolaroid from "$lib/BuildingPolaroid.svelte";
    import GridView from "$lib/GridView.svelte";
	import type { PageData } from "../[stylename]/$types";

    let { data }: { data: PageData } = $props();
    let style = $state(data.style)
    let obs = $state(data.observations)
    let observations = $state(obs.slice(0,5))
</script>

<div class="pagecontainer">
        <Menu />
        <div class="content">
            <div class="infocontainer">
                <div class="info">
                    <div class="mapBox">

                    </div>
                    <div class="descriptioncontainer">
                        <div class="description">
                            <h1 class="underlined">Style Name {style.name}</h1>
                            <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's
                                 standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to
                                  make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, 
                                  remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem 
                                  Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker 
                                  including versions of Lorem Ipsum.</p>
                        </div>
                        <div class="linneage">
                            <h2>Linneage</h2>
                            <div class="linneageBox"></div>
                        </div>
                    </div>
                </div>
            </div>
                <div class="polaroidslide">
                    <BuildingPolaroid building={observations[1]} big={true}></BuildingPolaroid>
                </div>
        </div>
</div>

<style>
    .pagecontainer {
        display: flex;
        flex-direction: column;
        gap: 10px;
        overflow: hidden;
    }

    .content {
        display: flex;
        flex-direction: row;
        gap: 30px;
        justify-content: center;
        /* align-items: center; */
        /* width: 100%; */
        /* height: 100%; */
    }
    
    .infocontainer {
        display: flex;
        flex-direction: column;
        /* height: 100%; */
        gap: 30px;
    }

    .info {
        display: flex;
        flex-direction: row;
        gap: 30px;
        height: 100%;
    }

    .descriptioncontainer {
        display: flex;
        flex-direction: column;
        width: 30vw;
    }

    .map {
        width: 35vw;
        /* height: 100%; */
    }

    .mapBox {
        border: 1px solid black;
        height: 55vh;
        width: 35vh;
    }

    .linneageBox {
        border: 1px solid black;
        height: 25vh;
    }

    .underlined {
        border-bottom: solid 2px black;
    }

</style>