<script>
    import List from "./List.svelte";
    import { auth_store } from "../../stores/authStore";
    import { Exit, LogOut, GetIps } from "../../../wailsjs/go/main/App";
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

    function getIps() {
        GetIps().then((ips) => {
            console.log("ips:", ips)
        }).catch((err) => {
            console.error("failed to get ips: ", err);
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
<button on:click={logOut}>Log Out</button><br/>
<button on:click={() => {Exit()}}>Exit</button><br/>
<button on:click={getIps}>Get Ips</button>