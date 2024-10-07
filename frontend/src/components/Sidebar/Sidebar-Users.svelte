<script>
    import { GetChats } from "../../../wailsjs/go/main/App.js";
    import List from "./List.svelte";
    import { selection_store } from '../../stores/selectionStore.js';
    import SmallList from "./SmallList.svelte";

    export let small;
    let chats = [];
    let showChats = true; // Updated to show chats instead of users

    function getChats() {
        GetChats().then((result) => {
            chats = result;
            // Automatically selects the first chat if there are chats available and no chat has been selected yet
            if (chats.length > 0 && $selection_store.name) {
                $selection_store = chats[0];
            }
        }).catch((err) => {
            console.error("failed to get chats, ", err)
        })
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