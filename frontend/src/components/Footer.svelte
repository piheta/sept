<script>
    import { SendMessage } from "../../wailsjs/go/main/App.js";
    import { message_store } from '../stores/messageStore.js';

    export let recipient;
    let input_txt = "";

    async function test() {
        try {
            let x = await SendMessage(input_txt);
            message_store.set(x); // Use the `set` method to update the store
            console.log("Message sent and response received: ", x);
        } catch (error) {
            console.error("Error sending message: ", error);
        }
    }
</script>

<div class="h-full w-full rounded-md bg-gray-900 p-1">
    <div>o o o o</div>
    <textarea
        bind:value={input_txt}
        class="w-full bg-gray-900 resize-none m-0 focus:outline-none active:outline-none"
        placeholder="Message {recipient}"
    />
    <div class="m-0 flex justify-end">
        <button on:click={test}>
            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24">
                <path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 2l-7 20l-4-9l-9-4Zm0 0L11 13"/>
            </svg>
        </button>
    </div>
</div>