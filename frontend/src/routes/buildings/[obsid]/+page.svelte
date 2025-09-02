<script lang="ts">
	import type { Observation } from "$lib";
	import Menu from "$lib/Menu.svelte";
	import type { PageData } from "../[obsid]/$types";
	import Polaroid from "$lib/Polaroid.svelte";

    let { data }: { data: PageData } = $props();
    let observation = $state(data.observation)
</script>

<div class="pagecontainer">
        <Menu />
        <Polaroid {observation} big={true}/>
</div>

<style>
    .pagecontainer {
        display: flex;
        flex-direction: column;
        gap: 10px;
        overflow: hidden;
    }

</style>