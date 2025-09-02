<script lang="ts">
	import type { Snippet } from 'svelte';

	let { polaroids } : { polaroids: Snippet }= $props();
</script>

<div class="container">
	{@render polaroids()}
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
