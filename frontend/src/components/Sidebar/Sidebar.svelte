<script>
    import SidebarUsers from "./Sidebar-Users.svelte";
    import SidebarServers from "./Sidebar-Servers.svelte";
    import SidebarSettings from "./Sidebar-Settings.svelte";

    let sidebar_mode = 0;
</script>

<!-- 3 BUTTONS -->
<div class="flex flex-col h-full">
    <div class="flex-shrink-0">
        <div class="flex justify-center items-center gap-2 ml-2">
            <button
                on:click={() => {
                    sidebar_mode = 0;
                }}
                class="relative w-full h-16 bg-gray-900 rounded-md flex items-center justify-center cursor-pointer"
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="1.5em"
                    height="1.5em"
                    viewBox="0 0 24 24"
                    ><g
                        fill="none"
                        stroke="white"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        ><path
                            d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"
                        /><circle cx="9" cy="7" r="4" /><path
                            d="M22 21v-2a4 4 0 0 0-3-3.87m-3-12a4 4 0 0 1 0 7.75"
                        /></g
                    ></svg
                >
                {#if sidebar_mode === 0}<div
                        class="absolute z-50 w-4 h-1 bottom-2 rounded-sm bg-[#11FFEE]"
                    ></div>{/if}
            </button>

            <button
                on:click={() => {
                    sidebar_mode = 1;
                }}
                class="relative w-full h-16 bg-gray-900 rounded-md flex items-center justify-center cursor-pointer"
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="1.5em"
                    height="1.5em"
                    viewBox="0 0 24 24"
                    ><g
                        fill="none"
                        stroke="white"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        ><rect
                            width="20"
                            height="8"
                            x="2"
                            y="2"
                            rx="2"
                            ry="2"
                        /><rect
                            width="20"
                            height="8"
                            x="2"
                            y="14"
                            rx="2"
                            ry="2"
                        /><path d="M6 6h.01M6 18h.01" /></g
                    ></svg
                >
                {#if sidebar_mode === 1}<div
                        class="absolute z-50 w-4 h-1 bottom-2 rounded-sm bg-[#11FFEE]"
                    ></div>{/if}
            </button>

            <button
                on:click={() => {
                    sidebar_mode = 2;
                }}
                class="relative w-full h-16 bg-gray-900 rounded-md flex items-center justify-center cursor-pointer"
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="1.5em"
                    height="1.5em"
                    viewBox="0 0 24 24"
                    ><g
                        fill="none"
                        stroke="white"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        ><path d="M20 7h-9m3 10H5" /><circle
                            cx="17"
                            cy="17"
                            r="3"
                        /><circle cx="7" cy="7" r="3" /></g
                    ></svg
                >
                {#if sidebar_mode === 2}<div
                        class="absolute z-50 w-4 h-1 bottom-2 rounded-sm bg-[#11FFEE]"
                    ></div>{/if}
            </button>
        </div>

        <div class="flex relative mt-2 mb-2 ml-2 h-8">
            <svg
                class="absolute left-2 top-2 text-gray-500"
                xmlns="http://www.w3.org/2000/svg"
                width="1em"
                height="1em"
                viewBox="0 0 24 24"
                ><g
                    fill="none"
                    stroke="currentColor"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    ><circle cx="11" cy="11" r="8" /><path
                        d="m21 21l-4.3-4.3"
                    /></g
                ></svg
            >
            <input
                class="w-full bg-gray-900 placeholder-gray-500 rounded-md pl-8 focus:outline-none"
                placeholder="Search"
            />
        </div>
    </div>

    <div class="overflow-y-auto flex-grow no-scrollbar">
        {#if sidebar_mode === 0}
            <SidebarUsers />
        {:else if sidebar_mode === 1}
            <SidebarServers />
        {:else if sidebar_mode === 2}
            <SidebarSettings />
        {:else}
            <h1>error</h1>
        {/if}
    </div>
</div>
