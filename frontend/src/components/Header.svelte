<script lang="ts">
    import { selection_store } from '../stores/selectionStore.js';
    
    // export let recipient;
</script>

<div class="flex items-center h-14 w-full">
    <img class=" w-10 min-w-10 h-10 ml-2 rounded-md object-cover object-top shadow-xl" alt="" src={$selection_store.avatar}/>

    <h1 class="pl-1 pr-1 mr-2 ml-2 rounded-md text-[1.25rem] flex-grow font-medium h-16 flex justify-start items-center select-text">{$selection_store.name}</h1>

    <button on:click={() => {}} class="relative w-10 h-10 mr-1 rounded-md flex items-center justify-center cursor-pointer transition-all duration-50 ease-in-out bg-opacity-0 hover:bg-opacity-5 hover:bg-gray-200">
        <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 24 24"><g fill="none" stroke="white" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 13l5.223 3.482a.5.5 0 0 0 .777-.416V7.87a.5.5 0 0 0-.752-.432L16 10.5"/><rect width="14" height="12" x="2" y="6" rx="2"/></g></svg>
    </button>

    <button on:click={() => {}} class="relative w-10 h-10 mr-2 rounded-md flex items-center justify-center cursor-pointer transition-all duration-50 ease-in-out bg-opacity-0 hover:bg-opacity-5 hover:bg-gray-200">
        <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 24 24"><path fill="none" stroke="white" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 14h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7a9 9 0 0 1 18 0v7a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3"/></svg>
    </button>
</div>
<div class="h-[1px] ml-2 mr-2 bg-gray-200 bg-opacity-5">
</div>