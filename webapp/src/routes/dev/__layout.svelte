<script>
    import NavSidePanel from "../../shared/components/common/navigation/NavSidePanel.svelte";
</script>

<div class="w-full absolute top-16 top-16 flex flex-row" style="height: calc(100vh - 64px)">
    <div class="m-auto mt-0 relative w-full">
        <slot></slot>
    </div>

    <NavSidePanel></NavSidePanel>
</div>