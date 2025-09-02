<script lang="ts">
	import Polaroid from './Polaroid.svelte';
	import type { Observation } from '$lib';

	let { observations, tall}  = $props();
</script>

<div class="container">
	{#each observations as o}
			<Polaroid data={o} {tall} />
	{/each}
</div>

<style>
	.container {
		display: flex;
		flex-direction: row;
		justify-content: center;
		flex-wrap: wrap;
		gap: 20px;
		max-height: 100%;
		padding-right: 10px;
	}
</style>
