<script lang="ts">
    import NavMenu from "../../shared/components/common/navigation/NavMenu.svelte";
    import {Link} from "../../shared/components/common/navigation/Link.js";

    const pageList: Link[] = [
        new Link("/tuple-deserializer", "Tuple visualizer"),
        new Link("/tuple-deserializer", "Tuple visualizer"),
        new Link("/tuple-deserializer", "Tuple visualizer"),
        new Link("/tuple-deserializer", "Tuple visualizer"),
        new Link("/tuple-deserializer", "Tuple visualizer"),
        new Link("/tuple-deserializer", "Tuple visualizer"),
    ]
</script>

<NavMenu links={pageList}></NavMenu>