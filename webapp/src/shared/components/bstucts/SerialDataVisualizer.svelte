<script lang="ts">
    export let serialData: string[];
</script>

<div class=" border-2 flex flex-row flex-wrap rounded-lg overflow-hidden">
    {#each serialData as data}
        <p class="border-2 p-2 text-lg">
            {data}
        </p>
    {/each}
</div>