<script>
    import { GetUsers, GetRooms } from "../../../wailsjs/go/main/App.js";
    import List from "./List.svelte";

    let users = [];
    let rooms = [];
    let showUsers = true;
    let showRooms = true;
    let selection = { type: 0, value: "" };

    async function getUsers() {
        const result = await GetUsers();
        console.log("Received users:", result);
        users = result;
        if (users.length > 0) {
            setSelection({ type: 0, value: users[0] });
        }
    }

    async function getRooms() {
        const result = await GetRooms();
        console.log("Received rooms:", result);
        rooms = result;
    }

    function setSelection(obj) {
        selection = obj;
    }

    function handleSelect(event) {
        setSelection({ type: event.detail.type, value: event.detail.item });
    }

    function handleDndConsider(e) {
        users = e.detail.items;
    }

    function handleDndFinalize(e) {
        users = e.detail.items;
    }

    getUsers();
    getRooms();
</script>

<List 
    title="Friends" 
    items={users}
    selectedItem={selection.value}
    collapsible={true}
    on:select={event => handleSelect({ detail: { type: 0, item: event.detail } })}
    bind:showItems={showUsers}
/>

<List 
    title="Rooms" 
    items={rooms}
    selectedItem={selection.value}
    collapsible={true}
    on:select={event => handleSelect({ detail: { type: 1, item: event.detail } })}
    bind:showItems={showRooms}
/>
