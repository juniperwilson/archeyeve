<script lang="ts">
	import Icon from './Icon.svelte';
</script>

<div class="container">
	<a data-sveltekit-preload-data="tap" href="/observations/add">
	<div class="logo">
		<Icon name="castle" color="#ffffff" width="40" />
		<h1>ArchEyeve</h1>
	</div>
</a>
	<a data-sveltekit-preload-data="tap" href="/buildings/view">
		<div class="icon">
			<Icon name="castle" color="#ffffff" width="50" />
			<h2>Buildings</h2>
		</div>
	</a>
	<a data-sveltekit-preload-data="tap" href="/">
		<div class="icon">
			<Icon name="eye2" color="#ffffff" width="50" />
			<h2>Observations</h2>
		</div>
	</a>
	<a data-sveltekit-preload-data="tap" href="/styles/view">
		<div class="icon">
			<Icon name="compass1" color="#ffffff" width="50" />
			<h2>Styles</h2>
		</div>
	</a>
</div>

<style>
	.container {
		max-width: 100%;
		max-height: 100%;
		display: flex;
		flex-direction: row;
		gap: 100px;

		background: #5e548e;
		background: #a266c7;
		background: #6d5fb3;
		padding: 10px;
		border: 1px solid black;
		color: white;
		align-items: center;
	}

	.logo {
		display: inline-flex;
		justify-content: center;
		gap: 10px;
	}

	.icon {
		display: inline-flex;
		justify-content: center;
		align-items: center;
		gap: 10px;
	}

	a {
		color: white;
		text-decoration: none;
	}
</style>
