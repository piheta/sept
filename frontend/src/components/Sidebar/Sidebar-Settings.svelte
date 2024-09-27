<script>
    import List from "./List.svelte";
    import { auth_store } from "../../stores/authStore";
    import { LogOut } from "../../../wailsjs/go/main/App";
    import {replace} from 'svelte-spa-router'

    let selection = "";

    function setSelection(obj) {
        selection = obj;
    }

    let settings = [
        { id: 1, name: "General" },
        { id: 2, name: "User" },
        { id: 3, name: "Audio" },
        { id: 4, name: "Proxy" },
        { id: 5, name: "Bots" },
        { id: 6, name: "Theme" }
    ];

    setSelection(settings[0].id);


    async function logOut() {
        try {
            let logout_err = await LogOut();
            if (logout_err != null) {
                console.error("Failed to log out:", logout_err);
                return;
            }

            auth_store.set({
                id: null,
                username: null,
                ip: null,
                avatar: null
            });

            // Clear any other relevant stores or local storage
            replace('/login');
        } catch (error) {
            console.error("An unexpected error occurred during logout:", error);
        }
    }
    
</script>

<List
    title="Settings"
    items={settings}
    selectedItem={selection}
    draggable={0}
    img={false}
/>
<button on:click={logOut}>Log Out</button>