<script>
    import List from "./List.svelte";

    let selection = "";

    function setSelection(obj) {
        selection = obj;
    }

    let settings = [
        { id: 1, name: "General" },
        { id: 2, name: "User" },
        { id: 3, name: "Audio" },
        { id: 4, name: "Proxy" },
        { id: 5, name: "Bots" },
        { id: 6, name: "Theme" }
    ];

    setSelection(settings[0].id);
</script>

<List
    title="Settings"
    items={settings}
    selectedItem={selection}
    draggable={0}
    img={false}
/>