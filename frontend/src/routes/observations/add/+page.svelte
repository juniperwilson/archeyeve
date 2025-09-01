<script lang="ts">
	import type { Observation } from "$lib";
	import ButtonGrid from "$lib/ButtonGrid.svelte";
	import Menu from "$lib/Menu.svelte";
	import Polaroid from "$lib/Polaroid.svelte";
	import type { PageData } from "../[obsid]/$types";

    let typeSelection = ""

    let { data }: { data: PageData } = $props();
    let buildingTypes = ["Religious", "Residential", "Commercial", "Industrial", "Infrastructure", "Military", "Agricultural", "Civic", "Medical", "Educational", "Transport", "Hotel", "Palace", "Sports", "Office", "Complex", "Unknown", "Other"]
</script>

<div class="pagecontainer">
        <Menu />
        <div class="content">
        <ButtonGrid titles={buildingTypes} onclick={() => {}} selected={typeSelection}></ButtonGrid>
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
        justify-content: center;
    }

</style>