<script lang="ts">
    import {Link} from "./Link.js";
    import {goto} from "$app/navigation";

    export let links: Link[] = [];

    const goToFunc = (path: string): () => void => {
        return () => goto(path);
    }
</script>

<div class="lg-box bg-accent-l2 wrapper flex-grid">
    {#each links as link}
        <div class="md-box card link"
             on:click={goToFunc(link.path)}>
            {link.name}
        </div>
    {/each}
</div>

<style>
    .wrapper {
        @apply w-11/12 md:w-2/3 px-2 py-5 mx-auto my-10 md:my-16;
    }

    .card {
        @apply md:w-[45%] xl:w-[30%] mx-2 my-4
        border-2 hover:cursor-pointer text-xl font-semibold hover:scale-105;
        transition: transform 0.2s;
    }
</style>