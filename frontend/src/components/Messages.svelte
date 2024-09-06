<script>
    import { onMount, onDestroy, afterUpdate } from 'svelte';
    import { GetUserUserMessages } from "../../wailsjs/go/main/App.js";
    import { message_store, chatbox_scroll_store } from '../stores/messageStore.js';
    
    let chatbox;

    async function getMessages() {
        try {
            let messages = await GetUserUserMessages(1);
            let parsedData;
            if (typeof messages === 'string') {
                parsedData = JSON.parse(messages);
            } else {
                parsedData = messages;
            }

            message_store.set(parsedData);

        } catch (error) {
            console.error("Error getting message: ", error);
        }
    }

    onMount(() => {
        getMessages();
    });

    afterUpdate(() => {
        if (!chatbox) {
            return;
        }

        // if new message is from sender, snap to bottom.
        // if new message is recieved, and already at bottom snap to new bottom.
        // if new message is recieved, and scrolled up, dont snap to bottom.

        chatbox.scrollTop = chatbox.scrollHeight;
    })
</script>

<div 
    class="bg-gray-700 flex-grow rounded-md overflow-auto mt-2 p-2 scrollbar"
    bind:this={chatbox}
>
    {#each $message_store as message, i}
        <div class="flex mb-3 {i === $message_store.length - 1 ? 'mb-0' : ''}" >
            <img class="w-12 h-12 rounded-md pointer-events-none select-none" src="https://i.imgur.com/EXSmx6x.png" alt="cat" />

            <div class="flex-col ml-2">
                <h1>
                    <span class="font-semibold">User {message.user_id ?? 'Unknown'}</span>
                    <span class="text-gray-300 text-sm ml-2">
                        {message.created_at ? new Date(message.created_at).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) : 'Unknown Time'}
                    </span>
                </h1>
                <p class="">{message.content ?? 'No content available'}</p>
            </div>
        </div>
    {/each}
</div>


<style>
.scrollbar {
    overflow: hidden; /* Hide scrollbar by default */
    scrollbar-color: theme("colors.gray.900") rgba(0,0,0,0); /* Style scrollbar color */
    transition: opacity 0.3s ease-in; /* Add transition for opacity */
}

.scrollbar:hover {
    overflow: auto; /* Show scrollbar on hover */
    transition: opacity 0.3s ease-in;
}

/* For Webkit browsers (Chrome, Safari) */
.scrollbar::-webkit-scrollbar {
    width: 12px; /* Full width of the scrollbar */
}

.scrollbar::-webkit-scrollbar-thumb {
    background-color: theme("colors.gray.900");
    border-radius: 8px;
    border: 2px solid theme("colors.gray.700"); /* Creates space around the thumb */
    box-sizing: border-box; /* Ensure the border adds space inside the track */
}

.scrollbar::-webkit-scrollbar-track {
    background-color: rgba(0,0,0,0); /* Track background */
}
</style>