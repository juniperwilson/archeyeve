<script lang="ts">
    import GridView from "$lib/GridView.svelte";
	import Menu from "$lib/Menu.svelte";
	import SearchTab from "$lib/SearchTab.svelte";
    import { doSearch } from "$lib/api";
	import StylesBar from "$lib/StylesBar.svelte";
	import type { PageData } from "./$types";

    const styles = ["manueline", "mudejar", "romanesque", "brutalism", "neo-mudejar", "neo-romanesque", "castro", "português suave"]
    // const stylesOLD = ['manueline', 'mudejar', 'gothic', 'baroque', 'romanesque', 'mannerist', 'military', 'neoclassical', 'roman', 'industrial', 'castro', 'renaissance', 'rococo', 'pombaline', 'art deco', 'modern', 'megalithic'];
    let { data }: { data: PageData } = $props();
    let observations = $state(data.observations)

    async function styleSearch(style: string, focus: boolean) {
        if (focus) {
            observations = await doSearch({style: style})
            console.log("started style search")
        } else {
            observations = await doSearch({})
            // console.log("unselected")
        }
    }
</script>

<div class="pagecontainer">
        <Menu />
        <StylesBar {styles} {styleSearch}/>
    <div class="content"> 
        <GridView {observations} />
        <SearchTab />
    </div>
</div>

<style>
    .pagecontainer {
        display: flex;
        flex-direction: column;
        gap: 10px;
        overflow: hidden;
    }

    .content {
        display: grid;
        grid-template-columns: 8fr 2fr ;
        grid-template-rows: 1fr;
        gap: 10px;
        height: 100%;
    }
</style>