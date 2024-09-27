<script>
    import Router from 'svelte-spa-router';
    import { wrap } from 'svelte-spa-router/wrap';
    import { GetAuthedUser } from "../wailsjs/go/main/App";
    import { auth_store } from "./stores/authStore";
    import { onMount } from 'svelte';
    import {push, pop, replace} from 'svelte-spa-router'


    // Import your route components
    import Auth from "./routes/Login.svelte";
    import Dashboard from "./routes/Dashboard.svelte";
    // Import other route components as needed

    let isAuthenticated = false;

    const routes = {
        '/': wrap({
            component: Dashboard,
            conditions: [
                () => isAuthenticated // Route to Dashboard only if authenticated
            ]
        }),
        '/login': wrap({
            component: Auth,
            conditions: [
                () => !isAuthenticated // Route to Login only if not authenticated
            ]
        })
    };

    async function checkAuth() {
        try {
            let user = await GetAuthedUser();
            $auth_store = user;
            isAuthenticated = !!user.id;
        } catch (error) {
            isAuthenticated = false;
        }
    }

    onMount(() => {
        checkAuth();
    });

    // Watch auth_store for changes
    $: isAuthenticated = !!$auth_store.id;

    $: if (isAuthenticated) {
        replace("/");
    } else {
        replace("/login");
    }

</script>

<main>
    <Router {routes}/>
</main>