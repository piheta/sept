<script lang="ts">
    import Router from 'svelte-spa-router';
    import { wrap } from 'svelte-spa-router/wrap';
    import { GetAuthedUser } from "../wailsjs/go/controllers/AuthController";
    import { auth_store } from "./stores/authStore";
    import { onMount } from 'svelte';
    import {replace} from 'svelte-spa-router'

    import Auth from "./routes/Login.svelte";
    import Dashboard from "./routes/Dashboard.svelte";

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

    function checkAuth() {
        GetAuthedUser().then((user) => {
            auth_store.set(user);
            isAuthenticated = !!user.id;
        }).catch(() => {
            isAuthenticated = false;
        })
    }

    onMount(() => {
        checkAuth();
    });

    // Watch auth_store for changes
    // $: isAuthenticated = !!$auth_store.id;

    $: if (isAuthenticated) {
        replace("/");
    } else {
        replace("/login");
    }

</script>

<main>
    <Router {routes}/>
</main>