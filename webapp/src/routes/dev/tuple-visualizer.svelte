<script lang="ts">
    import SerialDataVisualizer from "../../shared/components/bstucts/SerialDataVisualizer.svelte";

    const data = [
        "Created by 32",
        "Modified by 0",
        "Command executed before 32",
        "TID: (PageId: 0, TupleIndex: 0)",
        "NullBitmap: 00001111\n" +
        "| col1: -1 | col2: -1 | col3: 1 | col4: 1 |",
        "col1: 2",
        "col2: 4",
        "col3: 6",
        "col4: 9",
    ]

    let value: string;

    const debugValue = () => {
        console.log(value.length)
        for (let i = 0; i < value.length; i++) {
            console.log(value.charCodeAt(i))
        }
    }
</script>

<div class="base-component bg-accent-l2">
    <p class="font-semibold text-xl mx-auto">Tuple</p>

    <input class="text-black" bind:value on:change={debugValue}>

    <SerialDataVisualizer serialData={data}/>
</div>