<script lang="ts">

    export let serialData: string[] = [];
    export let rowSize: number = 0;
    export let rowSizes: number[] = [];

    let data: string[][] = [];
    let singleRow = false;
    if (!serialData) {
        throw new Error("serial data is required");
    } else if (rowSize != 0) {
        let j = -1;
        for (let i = 0; i < serialData.length; i++) {
            if (i % rowSize == 0) {
                j++
                data.push([]);
            }
            data[j].push(serialData[i])
        }
    } else if (rowSizes.length > 0) {
        let currentElement = 0;
        for (let row = 0; row < rowSizes.length; row++) {
            data.push([])
            for (let rowElement = 0; rowElement < rowSizes[row]; rowElement++) {
                if (serialData.length < currentElement) {
                    break;
                }
                data[row].push(serialData[currentElement])
                currentElement++;
            }
        }
    } else {
        singleRow = true;
        data[0] = serialData;
    }
</script>

{#if !singleRow}
<div class="wrapper v-stack">
    {#each data as row}
        <div class="row h-stack">
            {#each row as str}
                <div class="center">{str}</div>
            {/each}
        </div>
    {/each}
</div>
{:else}
<div class="single-row-wrapper">
    {#each data[0] as str}
        <div class="center">{str}</div>
    {/each}
</div>
{/if}

<style>
    .wrapper {
        @apply outline outline-2 outline-white rounded-lg overflow-hidden;
    }

    .wrapper .row:not(:last-child) {
        @apply border-b-2 border-white w-full;
    }

    .row {
        @apply w-full;
    }

    .row * {
        @apply flex-grow p-2 h-full relative;
    }

    .row *:not(:last-child) {
        @apply border-r-2 border-white;
    }

    .single-row-wrapper {
        @apply outline outline-2 outline-white rounded-lg overflow-hidden;
        @apply flex flex-row flex-wrap;
    }

    .single-row-wrapper div {
        @apply relative mr-[-2px] mb-[-2px] p-[8px] pr-[9px] pb-[9px];
        @apply border-b-2 border-r-2 border-white flex-grow text-center;
    }
</style>