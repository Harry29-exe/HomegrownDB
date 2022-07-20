<script lang="ts">
    import NavMenu from "../shared/components/common/navigation/NavMenu.svelte";
    import {Link} from "../shared/components/common/navigation/Link.js";
    import {devPaths} from "./dev/DevPaths";

    const paths= devPaths
</script>

<NavMenu links={paths}></NavMenu>