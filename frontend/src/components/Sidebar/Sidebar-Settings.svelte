<script lang="ts">
    import List from "./List.svelte";
    import { auth_store } from "../../stores/authStore";
    import { Exit, LogOut } from "../../../wailsjs/go/controllers/AuthController";
    import { SearchDht } from "../../../wailsjs/go/controllers/SignalingController";
    import {replace} from 'svelte-spa-router'

    let selection = $state("");

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

    function logOut() {
        LogOut().then(() => {
            auth_store.set({
                id: null,
                username: null,
                ip: null,
                avatar: null
            });

            replace('/login');
        }).catch((err) => {
            console.error("Failed to log out:", err);
        })
    }

    function searchDht(username) {
        console.log(username)
        SearchDht(username).then((user) => {
            console.log(user)
        })
    }
    
</script>

<List
    title="Settings"
    items={settings}
    selectedItem={selection}
    draggable={0}
    img={false}
/>
<button onclick={logOut}>Log Out</button><br/>
<button onclick={() => {Exit()}}>Exit</button><br/>
<button onclick={() => searchDht("Picheta")}>Search DHT</button>