<script>
    // import { GetServers } from "../../../wailsjs/go/controllers/App.js";
    import List from "./List.svelte";

    let servers = [];
    let selection = { type: 0, value: "" };

    // async function getServers() {
    //     const result = await GetServers();
    //     console.log("Received servers:", result);
    //     servers = result;
    //     if (servers.length > 0) {
    //         setSelection({ type: 0, value: servers[0] });
    //     }
    // }

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