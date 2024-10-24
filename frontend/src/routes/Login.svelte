<script lang="ts">
    import { slide } from 'svelte/transition';
    import Checkmark from '../assets/icons/Checkmark.svelte';
    import { auth_store } from "../stores/authStore.js"
    import { Register, Login } from '../../wailsjs/go/controllers/AuthController';
    import { replace } from 'svelte-spa-router';

    let username = '';
    let email = '';
    let password = '';
    let showPassword = false;
    let loginForm = true;

    function toggleLoginForm() {
        loginForm = !loginForm;
    }

    function togglePassword() {
        showPassword = !showPassword;
    }


    async function handleSubmit(event) {
    event.preventDefault(); // Prevent default form submission behavior

    console.log("Submitting form with:", { username, email, password });

    try {
        if (!loginForm) {
            // Register the user if it's not the login form
            await Register(username, email, password);
            toggleLoginForm(); // Switch to the login form after registration
            return;
        }

        // Login the user
        let user = await Login(email, password);
        if (!user) {
            throw new Error("Invalid login credentials");
        }

        console.log("Logged in user:", user);

        // Set the auth store with user details
        auth_store.set({
            id: user.id,
            username: user.username,
            ip: user.ip,
            avatar: user.avatar
        });

        // Redirect to homepage after successful login
        replace("/")
        
        } catch (error) {
            console.error("Error during submission:", error.message);
            // You can also display the error to the user, e.g., set an error state
        }
    }
</script>

<style>
    .reveal-password {
        cursor: pointer;
    }
</style>

<div class="absolute top-0 w-full h-20 wails-drag">

</div>
<div class="h-[100vh] w-full flex justify-center items-center">
    <div class="pr-6">
        <pre class="text-2xl text-gray-900 font-bold translate-y-[0.35rem] leading-6 select-none cursor-default">
╔═╗╔═╗╔═╗╔╦╗
╚═╗╠═ ╠═╝ ║
╠═╝╚═╝╩   ╩
        </pre>

    </div>    
    <form on:submit|preventDefault={handleSubmit}>
        <fieldset>
            <!-- <legend class="text-center text-white mb-4 text-lg">Register</legend> -->
            <div class="flex flex-col w-56">
                <label for="username">
                    {loginForm ? 'Login' : 'Register'} 
                    <span on:click={toggleLoginForm} class="text-gray-400 pl-2 hover:underline cursor-pointer">
                      {loginForm ? 'Register' : 'Login'}
                    </span>
                  </label>
                  
                <div class="flex relative mb-2 mt-2 h-8">
                    <svg class="absolute left-2 top-2 text-gray-500" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/></g></svg>
                    <input
                        id="user-email"
                        bind:value={email}
                        type="email"
                        class="w-full bg-gray-900 text-white placeholder-gray-500 rounded-md pl-8 focus:outline-none"
                        placeholder="Email"
                    />
                </div>

                <!-- <label for="user-email">Email:</label> -->
                 {#if !loginForm}
                    <div class="flex relative mb-2 h-8">
                        <svg class="absolute left-2 top-2 text-gray-500" xmlns="http://www.w3.org/2000/svg" width="1.1em" height="1.1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></g></svg>
                        <input
                            id="username"
                            bind:value={username}
                            type="text"
                            class="w-full bg-gray-900 text-white placeholder-gray-500 rounded-md pl-8 focus:outline-none"
                            placeholder="Username"
                            autofocus
                        />
                    </div>
                {/if}

                <div class="flex relative mb-2 h-8">
                    <svg class="absolute left-2 top-2 text-gray-500" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2.586 17.414A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814a6.5 6.5 0 1 0-4-4z"/><circle cx="16.5" cy="7.5" r=".5" fill="currentColor"/></svg>
                    {#if showPassword}
                        <input
                            id="password"
                            bind:value={password}
                            type="text"
                            class="w-full bg-gray-900 text-white placeholder-gray-500 rounded-md pl-8 pr-7 focus:outline-none"
                            placeholder="********"
                        />
                    {:else}
                        <input
                            id="password"
                            bind:value={password}
                            type="password"
                            class="w-full bg-gray-900 text-white placeholder-gray-500 rounded-md pl-8 pr-7 focus:outline-none"
                            placeholder="********"
                        />
                    {/if}
                    <button type="button" on:click={togglePassword} class="absolute right-2 top-2 text-gray-500 reveal-password">
                        {#if showPassword && password.length > 0}
                            <!-- Eye icon to indicate password is visible -->
                            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2.062 12.348a1 1 0 0 1 0-.696a10.75 10.75 0 0 1 19.876 0a1 1 0 0 1 0 .696a10.75 10.75 0 0 1-19.876 0"/><circle cx="12" cy="12" r="3"/></g></svg>
                            <!-- Eye slash icon to indicate password is hidden -->
                        {:else if !showPassword && password.length > 0}
                            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.733 5.076a10.744 10.744 0 0 1 11.205 6.575a1 1 0 0 1 0 .696a10.8 10.8 0 0 1-1.444 2.49m-6.41-.679a3 3 0 0 1-4.242-4.242"/><path d="M17.479 17.499a10.75 10.75 0 0 1-15.417-5.151a1 1 0 0 1 0-.696a10.75 10.75 0 0 1 4.446-5.143M2 2l20 20"/></g></svg>
                        {/if}
                    </button>



                </div>

                {#if !loginForm && password.length > 0}
                    <div class="text-[0.9rem]" transition:slide>
                        <p class="flex"><Checkmark hide={password.length < 8} /> At least 8 characters long</p>
                        <p class="flex"><Checkmark hide={!/[a-z]/.test(password)} /> At least 1 lowercase letter</p>
                        <p class="flex"><Checkmark hide={!/[A-Z]/.test(password)} /> At least 1 uppercase letter</p>
                        <p class="flex"><Checkmark hide={!/[!@#$%^&*(),.?":{}|<>]/.test(password)} /> At least 1 special symbol</p>
                        <p class="flex"><Checkmark hide={!/\d/.test(password)} /> At least 1 number</p>
                    </div>
                {/if}
                <input class="bg-gray-700 text-white mt-2 h-8 rounded-md cursor-pointer" type="submit" value={loginForm ? 'Login' : 'Register'}>
            </div>
        </fieldset>
    </form>
</div>
