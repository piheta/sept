<script>
    import { onMount, afterUpdate } from 'svelte';
    import { GetUserUserMessages, GetUser } from "../../wailsjs/go/main/App.js";
    import { message_store } from '../stores/messageStore.js';
    import { selection_store } from '../stores/selectionStore.js';
    import Message from './Message.svelte';

    let chatbox;
    let participants = new Map(); // Caching user details

    async function getUserDetails(user_id) {
        // Check if user details are already in cache
        if (participants.has(user_id)) {
            return participants.get(user_id);
        }
        try {
            let user = await GetUser(user_id);
            let parsedData;
            if (typeof user === 'string') {
                parsedData = JSON.parse(user);
            } else {
                parsedData = user;
            }

            // Store user details in cache
            participants.set(user_id, { username: parsedData.username, avatar: parsedData.avatar });

            return parsedData;
        } catch (error) {
            console.error("Error getting user: ", error);
            return null; // Handle errors gracefully
        }
    }

    async function getMessages() {
        try {
            let chatId = $selection_store.id; // Ensure reactivity
            let messages = await GetUserUserMessages(chatId);
            let parsedData;

            if (typeof messages === 'string') {
                parsedData = JSON.parse(messages);
            } else {
                parsedData = messages;
            }

            // Collect unique user IDs
            const userIds = new Set(parsedData.map(message => message.user_id));
            // Fetch user details and update participants map
            await Promise.all(
                Array.from(userIds).map(userId => getUserDetails(userId))
            );

            // Update message store
            message_store.set(parsedData);

        } catch (error) {
            console.error("Error getting messages: ", error);
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
        // if new message is received, and already at bottom snap to new bottom.
        // if new message is received, and scrolled up, don't snap to bottom.
        chatbox.scrollTop = chatbox.scrollHeight;
    });

    // Trigger getMessages whenever selection_store changes
    $: $selection_store, getMessages();

</script>

<div bind:this={chatbox} class="bg-gray-700 flex-grow h-20 rounded-md overflow-auto mt-2 p-2 scrollbar flex flex-col-reverse">
    <div>
      {#each $message_store as message, i}
        <Message
          content={message.content}
          created_at={message.created_at}
          last_message_create_at={i > 0 ? $message_store[i - 1].created_at : null}
          user_id={message.user_id}
          last_sender_user_id={i > 0 ? $message_store[i - 1].user_id : null}
          index={i}
          username={participants.get(message.user_id)?.username ?? 'Unknown'}
          avatar={participants.get(message.user_id)?.avatar ?? 'default-avatar-url'}
        />
      {/each}
    </div>
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