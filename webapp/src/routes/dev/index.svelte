<script lang="ts">
    import NavMenu from "../../shared/components/common/navigation/NavMenu.svelte";
    import {Link} from "../../shared/components/common/navigation/Link.js";

    const pageList: Link[] = [
        new Link("/dev/tuple-visualizer", "Tuple visualizer"),
        new Link("/dev/tuple-visualizer", "Tuple visualizer"),
        new Link("/dev/tuple-visualizer", "Tuple visualizer"),
        new Link("/dev/tuple-visualizer", "Tuple visualizer"),
        new Link("/dev/tuple-visualizer", "Tuple visualizer"),
        new Link("/dev/tuple-visualizer", "Tuple visualizer"),
    ]
</script>

<NavMenu links={pageList}></NavMenu>