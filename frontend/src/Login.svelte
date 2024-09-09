
<div class="h-[100vh] w-full flex justify-center items-center">
    <form id="app-login">
        <fieldset>
            <!-- <legend class="text-center mb-4">Login</legend> -->
                <div class="flex flex-col *:rounded-md">
                <label for="user-name">Username:</label>
                <div class="flex relative mt-2 mb-2 h-8">
                    <svg class="absolute left-2 top-2 text-gray-500" xmlns="http://www.w3.org/2000/svg" width="1.1em" height="1.1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></g></svg>
                    <input
                        name="user-name"
                        type="email"
                        class="w-full bg-gray-900 placeholder-gray-500 rounded-md pl-8 focus:outline-none"
                        placeholder="username"
                    />
                </div>
                <label for="password">Password:</label>
                <div class="flex relative mt-2 mb-2 h-8">
                    <svg class="absolute left-2 top-2 text-gray-500" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2.586 17.414A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814a6.5 6.5 0 1 0-4-4z"/><circle cx="16.5" cy="7.5" r=".5" fill="currentColor"/></svg>
                    <input
                        name="password"
                        type="password"
                        class="w-full bg-gray-900 placeholder-gray-500 rounded-md pl-8 focus:outline-none"
                        placeholder="********"
                    />
                </div>
                <input class="bg-gray-700 mt-2 h-8" name="login" type="submit" value="Login">
            </div>
        </fieldset>
    </form>
</div>