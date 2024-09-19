import { writable } from 'svelte/store';

// Initialize the store with data from localStorage if it exists, otherwise with a default object
const storedUser = JSON.parse(localStorage.getItem('user')) || {
    id: null,
    username: null,
    ip: null,
    avatar: null
};

export const auth_store = writable(storedUser);

// Subscribe to the store and update localStorage whenever the store changes
auth_store.subscribe((value) => {
    localStorage.setItem('user', JSON.stringify(value));
});