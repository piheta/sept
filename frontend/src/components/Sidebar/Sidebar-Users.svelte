<script>
    import { GetUsers, GetRooms } from "../../../wailsjs/go/main/App.js";
    import List from "./List.svelte";
    import ListElement from "./ListElement.svelte";

    let users = [];
    let rooms = [];
    let showUsers = true;
    let showRooms = true;
    let selection = { type: 0, value: "" };

    async function getUsers() {
        users = await GetUsers();
        if (users.length > 0) {
            setSelection({ type: 0, value: users[0] });
        }
    }

    async function getRooms() {
        rooms = await GetRooms();
    }

    function setSelection(obj) {
        selection = obj;
    }

    function handleSelect(event) {
        setSelection({ type: event.detail.type, value: event.detail.item });
    }

    getUsers();
    getRooms();
</script>

<List title="Friends" collapsible={true} bind:showItems={showUsers}>
    {#each users as user, index}
        <ListElement
            item={user}
            selected={user === selection.value && selection.type === 0}
            isLast={index === users.length - 1}
            on:select={(event) =>
                handleSelect({ detail: { type: 0, item: event.detail } })}
        />
    {/each}
    <svelte:fragment slot="selected">
        {#if selection.type === 0}
            <ListElement item={selection.value} selected={true} />
        {/if}
    </svelte:fragment>
</List>

<List title="Rooms" collapsible={true} bind:showItems={showRooms}>
    {#each rooms as room, index}
        <ListElement
            item={room}
            selected={room === selection.value && selection.type === 1}
            isLast={index === rooms.length - 1}
            on:select={(event) =>
                handleSelect({ detail: { type: 1, item: event.detail } })}
        />
    {/each}
    <svelte:fragment slot="selected">
        {#if selection.type === 1}
            <ListElement item={selection.value} selected={true} />
        {/if}
    </svelte:fragment>
</List>
