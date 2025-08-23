<script>
    import '@fontsource/roboto'
	
	let buttonProps = {
	class:[$$restProps.class]
	}
</script>
	<button on:click
			on:mouseover
			on:mouseenter
			on:mouseleave
		{...buttonProps}>
			<slot/>
	</button>

<style>
	button {
		font-size:1em;
        height: 100%;
        border-radius: 0px;
        border: 1px solid black;
        background-color: #f2eee2;
        font-family: "Roboto";
        font-weight: bold;
        text-wrap: nowrap;
	}

    .selected {
        background-color: #231942;
        background-color: #6dae47;
        color: white;
    }

    button:hover {
        /* text-shadow: 1px 1px 1px grey; */
        background-color: #8bb174;
    }

    /* button:focus {
        background-color: #231942;
        background-color: #6dae47;
        color: white;
    } */
    
</style>