<script>
    import Auth from "./pages/Auth.svelte";
    import Dashboard from "./pages/Dashboard.svelte";
    import { auth_store } from "./stores/authStore";
</script>


<!-- {#if $auth_store.user_id } -->
    <!-- <Dashboard /> -->
<!-- {:else} -->
    <Auth />
<!-- {/if} -->
