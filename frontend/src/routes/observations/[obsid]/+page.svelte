<script lang="ts">
	import type { Observation } from "$lib";
	import Menu from "$lib/Menu.svelte";
	import Polaroid from "$lib/Polaroid.svelte";
	import type { PageData } from "../[obsid]/$types";

    let { data }: { data: PageData } = $props();
    let observation = $state(data.observation)
</script>

<div class="pagecontainer">
        <Menu />
        <Polaroid {observation} />
</div>

<style>
    .pagecontainer {
        display: flex;
        flex-direction: column;
        gap: 10px;
        overflow: hidden;
    }

</style>