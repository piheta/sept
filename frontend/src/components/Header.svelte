<script>
    import { selection_store } from '../stores/selectionStore.js';
    
    export let recipient;
</script>

<div class="flex items-center h-full w-full">
    <img class="w-16 h-16 rounded-md" alt="" src={$selection_store.avatar}/>

    <h1 class="pl-2 mr-2 ml-2 rounded-md text-lg flex-grow bg-gray-700 h-16 flex items-center">{$selection_store.username}</h1>

    <button
    on:click={() => {
    }}
    class="relative w-16 h-16 bg-gray-900 rounded-md flex items-center justify-center cursor-pointer"
    >
        <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 24 24"><path fill="none" stroke="white" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 16.92v3a2 2 0 0 1-2.18 2a19.8 19.8 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.8 19.8 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.8 12.8 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.8 12.8 0 0 0 2.81.7A2 2 0 0 1 22 16.92M14.05 2a9 9 0 0 1 8 7.94m-8-3.94A5 5 0 0 1 18 10"/></svg>
    </button>
</div>
