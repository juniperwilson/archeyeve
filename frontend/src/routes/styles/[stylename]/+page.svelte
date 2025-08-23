<script lang="ts">
	import Menu from "$lib/Menu.svelte";
	import type { PageData } from "../[styleid]/$types";

    let { data }: { data: PageData } = $props();
</script>

<div class="pagecontainer">
        <Menu />
</div>

<style>
    .pagecontainer {
        display: flex;
        flex-direction: column;
        gap: 10px;
        overflow: hidden;
    }

</style>