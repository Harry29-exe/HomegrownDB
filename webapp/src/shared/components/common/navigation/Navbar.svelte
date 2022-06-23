<script lang="ts">
    import Logo from "../others/Logo.svelte";
    import {Link} from "./Link";
    import {onMount} from "svelte";

    export let links: Link[] = [
        new Link("/", "Home"),
        new Link("/dev", "Development")
    ]

    const toggleDarkMode = () => {
        if (document.documentElement.classList.contains("dark")) {
            document.documentElement.classList.remove('dark');
            themeIcon = '🌣';
        } else {
            document.documentElement.classList.add('dark');
            themeIcon = '☽'
        }
    }

    let themeIcon = '🌣';
    onMount(() => {
        themeIcon = darkModeOn()? '☽': '🌣';
    })

    const darkModeOn = (): boolean => {
        return document.documentElement.classList.contains("dark");
    }
</script>

<div class="w-full h-16 bg-accent-l3 h-stack p-4 pl-0 text-2xl font-bold text-white relative">
    <Logo/>

    <div class="flex-1"></div>

    {#each links as link}
        <p class="link">{link.name}</p>
    {/each}
    <p class="clickable" on:click={toggleDarkMode}>{themeIcon}</p>
</div>