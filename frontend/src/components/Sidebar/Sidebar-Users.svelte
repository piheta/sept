<script>
    import { GetUsers } from "../../../wailsjs/go/main/App.js";
    import List from "./List.svelte";
    import { selection_store } from '../../stores/selectionStore.js';

    let users = [];
    let showUsers = true;

    async function getUsers() {
        const result = await GetUsers();
        console.log("Received users:", result);
        users = result;
        if (users.length > 0 && $selection_store.username == "none") { //Automatically selects the first user: If there are users available and no user has been selected yet
            $selection_store = users[0]
        }
    }


    getUsers();
</script>

<List
    title="Friends"
    items={users}
    collapsible={true}
    bind:showItems={showUsers}
/>
