<script>
    import { SendMessage } from "../../wailsjs/go/main/App.js";
    import { message_store } from '../stores/messageStore.js';
    import { selection_store } from '../stores/selectionStore.js';

    export let recipient;
    let input_txt = "";

    async function sendMessage(message, chat_id) {
        input_txt = input_txt.trim()
        if (input_txt.length < 1) return;
        try {
            input_txt
            let allMessagesInChat = await SendMessage(message, chat_id);
            message_store.set(allMessagesInChat);
            input_txt = ""
        } catch (error) {
            console.error("Error sending message: ", error);
        }
    }

    function handleKeyDown(event) {
        if (event.key === 'Enter' && !event.shiftKey) {
            event.preventDefault();
            sendMessage(input_txt, $selection_store.id);
        }
    }
</script>

<div class="h-full w-full rounded-md bg-gray-900 p-1 flex flex-col shadow-xl">
    <div>o o o o</div>
    <textarea
        bind:value={input_txt}
        class="w-full bg-gray-900 resize-none m-0 focus:outline-none active:outline-none flex-grow"
        placeholder="Message {$selection_store.name}"
        on:keydown={handleKeyDown}
    />
    <div class="m-1 flex justify-end">
        <button on:click={() => {sendMessage(input_txt, $selection_store.id)}}>
            <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 24 24"><path fill="none" stroke="white" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.714 3.048a.498.498 0 0 0-.683.627l2.843 7.627a2 2 0 0 1 0 1.396l-2.842 7.627a.498.498 0 0 0 .682.627l18-8.5a.5.5 0 0 0 0-.904zM6 12h16"/></svg>
        </button>
    </div>
</div>