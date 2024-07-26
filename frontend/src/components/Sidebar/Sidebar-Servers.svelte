<script>
    import { GetServers } from '../../../wailsjs/go/main/App.js';
    import List from './List.svelte';
    import ListElement from './ListElement.svelte';

    let servers = [];
    let selection = { type: 0, value: "" };

    async function getServers() {
    servers = await GetServers();
    if (servers.length > 0) {
        setSelection({ type: 0, value: servers[0] });
    }
    }

    function setSelection(obj) {
    selection = obj;
    }

    getServers();

    function handleSelect(event) {
    setSelection({ type: 0, value: event.detail });
    }
</script>


<List title="Servers">
    {#each servers as server, index}
        <ListElement 
        item={server}
        selected={server === selection.value}
        isLast={index === servers.length - 1}
        on:select={handleSelect}
        />
    {/each}
</List>
