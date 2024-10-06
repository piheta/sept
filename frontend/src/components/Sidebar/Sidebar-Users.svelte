<script>
    import { GetChats } from "../../../wailsjs/go/main/App.js";
    import List from "./List.svelte";
    import { selection_store } from '../../stores/selectionStore.js';
    import SmallList from "./SmallList.svelte";

    export let small;
    let chats = [];
    let showChats = true; // Updated to show chats instead of users

    async function getChats() {
        const result = await GetChats();
        console.log("Received chats:", result);
        chats = result;

        if (chats.length > 0 && $selection_store.name) { // Automatically selects the first chat if there are chats available and no chat has been selected yet
            console.log(chats[0]);
            $selection_store = chats[0];
        }
    }

    getChats();
</script>

{#if small}
    <SmallList 
        items={chats}
    />
{:else}
    <List
        title="Friends"
        items={chats}
        collapsible={true}
        bind:showItems={showChats}
    />
{/if}