import { writable } from 'svelte/store';

// Initialize the store with data from localStorage if it exists, otherwise with a default object
const storedSelection = JSON.parse(localStorage.getItem('selection')) || {
    id: 1,
    user_id: "1",
    username: "none",
    ip: "127.0.0.1",
    avatar: ""
};

export const selection_store = writable(storedSelection);

// Subscribe to the store and update localStorage whenever the store changes
selection_store.subscribe((value) => {
    localStorage.setItem('selection', JSON.stringify(value));
});