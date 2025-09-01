<script lang="ts">
	import RangeSlider from 'svelte-range-slider-pips';
	import '@fontsource/roboto';
	import Button from './Button.svelte';
	import type { Search } from './api';
	import SelectableButtonsBar from './SelectableButtonsBar.svelte';
	import portugalMap from '$lib/assets/portugalMap.png'

	let { search } = $props();

	let date = new Date();

	let areaSearch = $state({
		name: 'Square'
	});

	let values = $state([-100, 1992]);
	let min = $state(-100);
	let max = $state(410);
	let minLat = 36;
	let maxLat = 43;
	let minLng = -10;
	let maxLng = -6;
	let lats = $state([minLat, maxLat]);
	let longs = $state([minLng, maxLng]);

	let years = $state([-10000, date.getFullYear()]);
	let style = $state('');
	let radius = $state(0);

	function formatter(v: number) {
		if (v == -100) {
			return '10.000 BC';
		} else if (v == 0) {
			return '0';
		} else if (v == 40) {
			return '1000';
		} else if (v == 90) {
			return '1500';
		} else if (v == 140) {
			return '1755';
		} else if (v == 410) {
			return '2025';
		} else {
			return '';
		}
	}

	function find() {
		if (areaSearch.name == 'Circular') {
			lats[1] = 0;
			longs[1] = 0;
		} else {
			radius = 0;
		}
		let s: Search = {
			style: style,
			aftyear: years[0],
			befyear: years[1],
			lat1: lats[0],
			lng1: longs[0],
			lat2: lats[1],
			lng2: lats[1],
			rad: radius
		};

		search(s);
	}

	function handleFormatter(v: number) {
		if (v <= 0) {
			//inc 100 prehistory
			return v * 100;
		} else if (0 < v && v <= 40) {
			//inc 25; 1-1000
			return v * 25;
		} else if (40 < v && v <= 90) {
			//inc 10; 1001-1500
			return v * 10 + 600;
		} else if (90 < v && v <= 140) {
			//inc 5; 1501-1755 start 91
			return v * 5 + 1055;
		} else {
			//inc 1; 1756-2025 start 141
			return v + 1615;
		}
	}

	function changeSearch(searchType: string) {
		areaSearch.name = searchType;
	}
</script>

<div class="container">
	<h1>Filters</h1>

	<form>
		<div class="inputLabel">
			<label for="style">Style:</label>
			<input id="style" bind:value={style} />
		</div>
		<!-- year -->
		<div class="areaSearch">
			<div class="sliders">
				<RangeSlider
					class="slider"
					pips
					all="label"
					{formatter}
					{handleFormatter}
					float
					range
					{min}
					{max}
					bind:values
				/>
			</div>
			<p>Structures built between {handleFormatter(values[0])} and {handleFormatter(values[1])}.</p>
		</div>
		<div class="areaSearch">
			<SelectableButtonsBar
				selected={areaSearch}
				titles={['Square', 'Circular']}
				onclick={changeSearch}
			/>
			{#if areaSearch.name == 'Square'}
				<!-- lat - lat -->
				<div class="sliderMap">
					<div class="sliders">
						<RangeSlider
							vertical={true}
							step={0.01}
							float
							range
							min={minLat}
							max={maxLat}
							bind:values={lats}
						/>
					</div>
					<img src={portugalMap} alt="map of Portugal"/>
				</div>
				<p>Latitudes between {lats[0].toFixed(2)} and {lats[1].toFixed(2)}.</p>
				<!-- lng - lng -->
				<div class="sliders">
					<RangeSlider step={0.01} float range min={minLng} max={maxLng} bind:values={longs} />
				</div>
				<p>Longitudes between {longs[0].toFixed(2)} and {longs[1].toFixed(2)}.</p>
			{:else}
				<div class="sliderMap">
					<div class="sliders">
						<RangeSlider
							vertical={true}
							step={0.01}
							float
							min={minLat}
							max={maxLat}
							bind:value={lats[0]}
						/>
					</div>
					<img src={portugalMap} alt="map of Portugal"/>
				</div>
				<p>Latitude {lats[0].toFixed(2)}</p>
				<div class="sliders">
					<RangeSlider step={0.01} float min={minLng} max={maxLng} bind:value={longs[0]} />
				</div>
				<p>Longitude {longs[0].toFixed(2)}.</p>
				<div class="inputLabel">
					<label for="radius">Radius: </label>
					<input id="radius" type="number" bind:value={radius} min="0" max="500" />
				</div>
			{/if}
		</div>
		<div class="searchButton">
			<Button onclick={find} text="SEARCH" isSelected="false" />
		</div>
	</form>
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		justify-content: start;
		background: #9f86c0;
		border-width: 10%;
		border-style: solid;
		height: 100%;
		max-width: 30vw;
		border: 1px solid black;
		text-align: center;
		color: white;
		font-family: 'Roboto';
		padding: 20px;
	}

	form {
		height: 100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 15px;
		width: 100%;
		background-color: white;
		color: black;
		padding: 5%;
	}

	.areaSearch {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		width: 100%;
		background-color: white;
		color: black;
		padding: 5%;
		border: 1px solid black;
	}

	.searchButton {
		height: fit-content;
		margin: 5%;
	}

	.sliders {
		width: 90%;
	}

	.inputLabel {
		text-wrap: none;
	}

	.sliderMap {
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
		padding: 5%;
	}

	img {
		width: 70%;
		height: auto;
		max-width: 100%;
	}


</style>
