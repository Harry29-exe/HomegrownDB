<script lang="ts">
    import Logo from "../others/Logo.svelte";
    import {Link} from "./Link";
    import {onMount} from "svelte";
    import {goto} from "$app/navigation"

    export let links: Link[] = [
        new Link("/", "Home"),
        new Link("/docs", "Documentation"),
        new Link("/dev", "Development")
    ]

    const toggleDarkMode = () => {
        if (document.documentElement.classList.contains("dark")) {
            document.documentElement.classList.remove('dark');
            themeIcon = '🌣';
        } else {
            document.documentElement.classList.add('dark');
            themeIcon = '☽'
            localStorage.theme = 'dark'
        }
    }

    let themeIcon = '🌣';
    onMount(() => {
        themeIcon = darkModeOn()? '☽': '🌣';
    })

    const darkModeOn = (): boolean => {
        return document.documentElement.classList.contains("dark");
    }

    const go = (path: string): () => void => {
        return () => goto(path);
    }
</script>

<div class="navbar">
    <p on:click={go("/")} class="hover:cursor-pointer select-none hover:underline decoration-4">
        <Logo/>
    </p>

    <div class="flex-1"></div>

    {#each links as link}
        <p class="link px-4" on:click={go(link.path)}>{link.name}</p>
        <p class="text-2xl font-bold">/</p>
    {/each}
    <p class="clickable px-2" on:click={toggleDarkMode}>{themeIcon}</p>
</div>