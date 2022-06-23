<script lang="ts">
    import "../shared/styles/app.css"
    import "../shared/styles/colors.css"
    import "../shared/styles/containers.css"
    import "../shared/styles/inputs.css"
    import "../shared/styles/text.css"
    import Navbar from "../shared/components/common/navigation/Navbar.svelte";
    import {onMount} from "svelte";

    onMount(() => {
        if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
            document.documentElement.classList.add('dark')
        } else {
            document.documentElement.classList.remove('dark')
        }
    })
</script>

<Navbar/>
<slot></slot>