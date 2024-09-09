<script>
    import { slide } from 'svelte/transition';
    import Checkmark from './assets/icons/Checkmark.svelte';

    let username = '';
    let email = '';
    let password = '';
    let showPassword = false;

    function togglePassword() {
        showPassword = !showPassword;
    }

    function handleSubmit(event) {
        event.preventDefault();
        // Handle form submission
        console.log({ username, email, password });
    }
</script>

<style>
    .reveal-password {
        cursor: pointer;
    }
</style>

<div class="h-[100vh] w-full flex justify-center items-center">
    <form on:submit|preventDefault={handleSubmit}>
        <fieldset>
            <!-- <legend class="text-center text-white mb-4 text-lg">Register</legend> -->
            <div class="flex flex-col">
                <label for="username">Username:</label>
                <div class="flex relative mt-2 mb-2 h-8">
                    <svg class="absolute left-2 top-2 text-gray-500" xmlns="http://www.w3.org/2000/svg" width="1.1em" height="1.1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></g></svg>
                    <input
                        id="username"
                        bind:value={username}
                        type="text"
                        class="w-full bg-gray-900 text-white placeholder-gray-500 rounded-md pl-8 focus:outline-none"
                        placeholder="Username"
                    />
                </div>

                <label for="user-email">Email:</label>
                <div class="flex relative mt-2 mb-2 h-8">
                    <svg class="absolute left-2 top-2 text-gray-500" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/></g></svg>
                    <input
                        id="user-email"
                        bind:value={email}
                        type="email"
                        class="w-full bg-gray-900 text-white placeholder-gray-500 rounded-md pl-8 focus:outline-none"
                        placeholder="Email"
                    />
                </div>

                <label for="password">Password:</label>
                <div class="flex relative mt-2 mb-2 h-8">
                    <svg class="absolute left-2 top-2 text-gray-500" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2.586 17.414A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814a6.5 6.5 0 1 0-4-4z"/><circle cx="16.5" cy="7.5" r=".5" fill="currentColor"/></svg>
                    {#if showPassword}
                        <input
                            id="password"
                            bind:value={password}
                            type="text"
                            class="w-full bg-gray-900 text-white placeholder-gray-500 rounded-md pl-8 focus:outline-none"
                            placeholder="********"
                        />
                    {:else}
                        <input
                            id="password"
                            bind:value={password}
                            type="password"
                            class="w-full bg-gray-900 text-white placeholder-gray-500 rounded-md pl-8 focus:outline-none"
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

                {#if password.length > 0}
                    <div class="text-[0.9rem]" transition:slide>
                        <p class="flex"><Checkmark hide={password.length < 8} /> At least 8 characters long</p>
                        <p class="flex"><Checkmark hide={!/[a-z]/.test(password)} /> At least 1 lowercase letter</p>
                        <p class="flex"><Checkmark hide={!/[A-Z]/.test(password)} /> At least 1 uppercase letter</p>
                        <p class="flex"><Checkmark hide={!/[!@#$%^&*(),.?":{}|<>]/.test(password)} /> At least 1 special symbol</p>
                        <p class="flex"><Checkmark hide={!/\d/.test(password)} /> At least 1 number</p>
                    </div>
                {/if}
                <input class="bg-gray-700 text-white mt-4 h-8 rounded-md cursor-pointer" type="submit" value="Register">
            </div>
        </fieldset>
    </form>
</div>
