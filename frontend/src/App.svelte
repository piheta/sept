<script>
    import Router from 'svelte-spa-router';
    import { wrap } from 'svelte-spa-router/wrap';
    import { GetAuthedUser } from "../wailsjs/go/main/App";
    import { auth_store } from "./stores/authStore";
    import { onMount } from 'svelte';

    // Import your route components
    import Auth from "./routes/Login.svelte";
    import Dashboard from "./routes/Dashboard.svelte";
    // Import other route components as needed

    let isAuthenticated = false;

    const routes = {
        '/': wrap({
            component: Dashboard,
            conditions: [
                () => isAuthenticated
            ]
        }),
        '/login': wrap({
            component: Auth,
            conditions: [
                () => !isAuthenticated
            ]
        }),
        // Add other routes as needed
    };

    async function checkAuth() {
        try {
            let user = await GetAuthedUser()
            $auth_store = user
            isAuthenticated = !!user.id
            console.log(user)
        } catch (error) {
            console.error("failed to get authed")
            isAuthenticated = false
        }
    }

    onMount(() => {
        checkAuth();
    });

    // Reactive statement to update routing when auth state changes
    $: if (isAuthenticated) {
        window.location.hash = '/';
    } else {
        window.location.hash = '/login';
    }
</script>

<main>
    <Router {routes}/>
</main>