<script lang="ts">
    import RangeSlider from 'svelte-range-slider-pips';
    import '@fontsource/roboto'
	import Button from './Button.svelte';

    let style = ""
    let values = $state([-100, 1992]);
    let displayVals = $state([0, 100])

    let min = $state(-100)
    let max = $state(410)

    function formatter(v:number) {
       if (v == -100) {
        return "10.000 BC"
       } else if (v == 0) {
        return "0" 
       } else if (v == 40) {
        return "1000"
       } else if (v == 90) {
        return "1500"
       } else if (v == 140) {
        return "1755"
       } else if (v == 410) {
        return "2025"
       } else {
        return ""
       }
    }

    function handleFormatter(v:number) {
        if (v <= 0) {
            //inc 100 prehistory
            return v * 100
        } else if (0 < v && v <= 40 ){
            //inc 25; 1-1000
            return v * 25
        } else if (40 < v && v <= 90) {
            //inc 10; 1001-1500
            return v * 10 + 600
        } else if (90 < v && v <= 140) {
            //inc 5; 1501-1755 start 91
            return v*5+1055
        } else {
            //inc 1; 1756-2025 start 141
            return v+1615
        }
    }

    function search() {
        
    }

</script>
<div class="container">
    <h1>Filters</h1>
        
    <form>
        <label for="style" >Style:</label>
        <input id="style" value={style} />
        <RangeSlider pips all="label" formatter={formatter} {handleFormatter} float range {min} {max} bind:values />
        <p>Structures built between {formatter(values[0])} and {formatter(values[1])}.</p>
        <Button class="primary lg" on:click={search}>Search</Button>


    </form>
</div>

<style>
    
    .container {
        background: #9f86c0;
        border-width: 10%;
        border-style: solid;
        height: 100%;
        max-width: 100%;
        border: 1px solid black;
        text-align: center;
        color: white;
        font-family: "Roboto";
    }

</style>