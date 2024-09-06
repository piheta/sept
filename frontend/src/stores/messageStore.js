// messageStore.js
import { writable } from 'svelte/store';

export const message_store = writable([
    { "id": 1, "chat_id": 1, "user_id": 1, "content": "no messages", "created_at": "2024-09-05T10:17:14Z" },
]);


export const chatbox_scroll_store = writable(100)
