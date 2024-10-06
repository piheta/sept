<script>
    import { selection_store } from '../stores/selectionStore.js';
    
    export let recipient;
</script>

<div class="flex items-center h-full w-full">
    <img class="w-16 h-16 rounded-md object-cover object-top shadow-xl" alt="" src={$selection_store.avatar}/>

    <h1 class="pl-1 pr-1 mr-2 ml-2 rounded-md text-[1.25rem] flex-grow  h-16 flex justify-start items-center select-text">{$selection_store.name}</h1>

    <button on:click={() => {}} class="relative w-16 h-16 bg-gray-100 rounded-md flex items-center justify-center cursor-pointer transition-all duration-50 ease-in-out bg-opacity-0 hover:bg-opacity-5 hover:bg-gray-200">
       <svg xmlns="http://www.w3.org/2000/svg" width="1.6em" height="1.6em" viewBox="0 0 24 24" fill="currentColor"><path d="M22 17.0022C21.999 19.8731 19.9816 22.2726 17.2872 22.8616L16.6492 20.9476C17.8532 20.7511 18.8765 20.0171 19.4649 19H17C15.8954 19 15 18.1046 15 17V13C15 11.8954 15.8954 11 17 11H19.9381C19.446 7.05369 16.0796 4 12 4C7.92038 4 4.55399 7.05369 4.06189 11H7C8.10457 11 9 11.8954 9 13V17C9 18.1046 8.10457 19 7 19H4C2.89543 19 2 18.1046 2 17V12C2 6.47715 6.47715 2 12 2C17.5228 2 22 6.47715 22 12V12.9987V13V17V17.0013V17.0022Z"></path></svg>
    </button>
</div>
