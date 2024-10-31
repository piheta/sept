<script lang="ts">
    import List from "./List.svelte";

    let servers = [];
    let selection = $state({ type: 0, value: "" });

    function setSelection(obj) {
        selection = obj;
    }

    // getServers();

    function handleSelect(event) {
        setSelection({ type: 0, value: event.detail });
    }

</script>

<List 
    title="Servers"
    items={servers}
    selectedItem={selection.value}
    on:select={handleSelect}
/>