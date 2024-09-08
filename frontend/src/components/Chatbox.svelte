<script>
    import { onMount, afterUpdate } from 'svelte';
    import { GetChatMessages, GetUser } from "../../wailsjs/go/main/App.js";
    import { message_store } from '../stores/messageStore.js';
    import { selection_store } from '../stores/selectionStore.js';
    import Message from './Message.svelte';

    let chatbox;
    let participants = new Map(); // Caching user details

    async function getUserDetails(user_id) {
        if (participants.has(user_id)) {// Check if user details are already in cache
            return participants.get(user_id);
        }
        try {
            let user = await GetUser(user_id);
            participants.set(user_id, { username: user.username, avatar: user.avatar }); // Store user details in cache

            return user;
        } catch (error) {
            console.error("Error getting user: ", error);
            return null;
        }
    }

    async function getMessages() {
        try {
            let chatId = $selection_store.id; // Ensure reactivity
            let messages = await GetChatMessages(chatId);

            const userIds = new Set(messages.map(message => message.user_id)); // Collect unique user IDs
            await Promise.all(
                Array.from(userIds).map(userId => getUserDetails(userId)) // Fetch user details and update participants map
            );

            message_store.set(messages);
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

    $: $selection_store, getMessages(); // Trigger getMessages whenever selection_store changes
</script>

<div bind:this={chatbox} class="bg-gray-700 flex-grow h-20 rounded-md overflow-auto mt-2 scrollbar-chatbox flex flex-col-reverse">
    <div class="p-2">
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