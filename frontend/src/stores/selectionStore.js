import { writable } from 'svelte/store';

export const selection_store = writable({
    id: 1,
    user_id: "1",
    username: "none",
    ip: "127.0.0.1",
    avatar: ""
});