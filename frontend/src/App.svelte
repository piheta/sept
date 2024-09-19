<script>
    import { StartLoggedIn } from "../wailsjs/go/main/App";
    import Auth from "./pages/Auth.svelte";
    import Dashboard from "./pages/Dashboard.svelte";
    import { auth_store } from "./stores/authStore";

    if ($auth_store.id) {
        StartLoggedIn($auth_store.id)
    }
</script>


{#if $auth_store.id }
    <Dashboard />
{:else}
    <Auth />
{/if}
