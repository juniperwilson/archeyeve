<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import '@fontsource/metamorphous'

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children?.()}

<style>
	:global(*, :after, :before) {
		box-sizing: border-box;
		font-family: "Metamorphous";
	}

	

</style>