<script lang="ts">
    import {Link} from "./Link";
    import {goto} from "$app/navigation"

    export let paths: Link[]

    let open = true;

    function toggle() {
        open = !open;
    }
</script>

<div on:click={toggle}
        class="expander bg-accent-l2 center hover:bg-accent-l1">
    {#if open}
        <p>&gt;</p>
    {:else}
        <p>&lt;</p>
    {/if}
</div>

{#if open}
<div class="relative bg-accent-l3 z-0 max-w-[200px] min-w-[200px] h-full flex-grow-0 v-stack px-2">
    {#each paths as link}
    <div class="link-element" on:click={goto(link.path)}>
        <div class="text-xl">{link.name}</div>
        <div class="text-md">{link.description}</div>
    </div>
    {/each}
</div>
{/if}

<style>
    .expander {
        @apply relative z-0 max-w-[16px] min-h-[16px] w-[16px] h-full;
        @apply flex-grow-0 hover:cursor-pointer;
    }

    .link-element {
        @apply w-full px-1 py-2;
        @apply border-b-2 border-white;
        @apply hover:underline hover:cursor-pointer;
    }
</style>