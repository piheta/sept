<script>
    import { GetChats } from "../../../wailsjs/go/controllers/App.js";
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

    function toggleShow() {
            showChats = !showChats;
    }

    getChats();
</script>

{#if small}
    <SmallList 
        items={chats}
    />
{:else}
    <button on:click={toggleShow} class="w-full flex items-center align-middle text-center mt-0 pt-0 pb-1 pl-3 cursor-pointer group">
        <span class="inline-block transition-transform duration-150 ease-in-out" class:rotate-0={showChats} class:-rotate-90={!showChats}>
            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 16 16" ><path fill="white" fill-rule="evenodd" d="M2.97 5.47a.75.75 0 0 1 1.06 0L8 9.44l3.97-3.97a.75.75 0 1 1 1.06 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0l-4.5-4.5a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/></svg>
        </span>
        <p class="text-lg inline-block ml-1">
            Friends
        </p>
        <span class="ml-auto mr-1 w-5 transition-opacity duration-150 ease-in-out opacity-0 group-hover:opacity-100">
            <svg xmlns="http://www.w3.org/2000/svg" width="1.2em" height="1.2em" viewBox="0 0 24 24"><g fill="none" stroke="white" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></g></svg>
        </span>
    </button>

    <List
        items={chats}
        bind:showItems={showChats}
    />
{/if}