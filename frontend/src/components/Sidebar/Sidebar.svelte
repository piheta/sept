<script lang="ts">
  import SidebarUsers from "./Sidebar-Users.svelte";
  import SidebarServers from "./Sidebar-Servers.svelte";
  import SidebarSettings from "./Sidebar-Settings.svelte";
  import { Search } from "../../../wailsjs/go/controllers/AuthController";
  import {
    SearchDht,
    SendOffer,
  } from "../../../wailsjs/go/controllers/SignalingController";

  export let small;
  let inputRef;

  let sidebar_mode = 0;
  let searchQuery = "";
  let searchResult = null;

  let foundUser = null;

  async function search(query) {
    try {
      let result = await Search(query);
      searchResult = result;
    } catch (error) {
      console.error(error);
    }
  }

  function handleInput(event) {
    if (searchQuery.length > 2) {
      search(searchQuery);
    } else {
      searchResult = null;
    }

    if (searchQuery.startsWith("add:")) {
      foundUser = null;
      sidebar_mode = 3;
    } else {
      sidebar_mode = 0;
    }
  }

  function handleEnter(event) {
    if (event.key === "Enter" && sidebar_mode === 3) {
      SearchDht(searchQuery.split(" ")[1])
        .then((user) => {
          console.log(user);
          foundUser = user;
        })
        .catch((err) => {
          console.log(err);
        });
    }
  }

  function fillSearchBoxAdd() {
    foundUser = null;
    sidebar_mode = 3;
    searchQuery = "add: ";
    inputRef.focus();
  }

  $: if (small && (sidebar_mode === 3 || sidebar_mode == 2)) {
    sidebar_mode = 0;
    searchQuery = "";
  }
</script>

<div class="flex flex-col h-full">
  <div class="flex-shrink-0">
    <!--* 3 BUTTONS -->
    <div
      class={`flex justify-center items-center ml-2 ${small ? "flex-col !gap-0 !mt-0" : "gap-4 mt-2"}`}
    >
      <button
        on:click={() => {
          sidebar_mode = 0;
          searchQuery = "";
        }}
        class:!mt-2={small}
        class="relative w-full h-12 rounded-md flex items-center justify-center cursor-pointer"
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
            ><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" /><circle
              cx="9"
              cy="7"
              r="4"
            /><path d="M22 21v-2a4 4 0 0 0-3-3.87m-3-12a4 4 0 0 1 0 7.75" /></g
          ></svg
        >
        <!-- {#if sidebar_mode === 0}<div class="absolute z-50 w-6 h-1 bottom-2.5 rounded-sm bg-[#11FFEE]"></div>{/if} -->
      </button>

      <button
        on:click={() => {
          sidebar_mode = 1;
          searchQuery = "";
        }}
        class="relative w-full h-12 rounded-md flex items-center justify-center cursor-pointer"
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
            ><rect width="20" height="8" x="2" y="2" rx="2" ry="2" /><rect
              width="20"
              height="8"
              x="2"
              y="14"
              rx="2"
              ry="2"
            /><path d="M6 6h.01M6 18h.01" /></g
          ></svg
        >
        <!-- {#if sidebar_mode === 1}<div class="absolute z-50 w-4 h-1 bottom-2 rounded-sm bg-[#11FFEE]"></div>{/if} -->
      </button>

      <button
        on:click={() => {
          sidebar_mode = 2;
          searchQuery = "";
        }}
        class:hidden={small}
        class="relative w-full h-12 rounded-md flex items-center justify-center cursor-pointer"
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
            ><path d="M20 7h-9m3 10H5" /><circle cx="17" cy="17" r="3" /><circle
              cx="7"
              cy="7"
              r="3"
            /></g
          ></svg
        >
        <!-- {#if sidebar_mode === 2}<div class="absolute z-50 w-4 h-1 bottom-2 rounded-sm bg-[#11FFEE]"></div>{/if} -->
      </button>

      <button
        on:click={() => {
          sidebar_mode = 0;
        }}
        class:hidden={!small}
        class="relative w-full h-12 rounded-md flex items-center justify-center cursor-pointer"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="1.5em"
          height="1.5em"
          viewBox="0 0 24 24"
          ><g
            fill="none"
            stroke="currentColor"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            ><circle cx="11" cy="11" r="8" /><path d="m21 21l-4.3-4.3" /></g
          ></svg
        >
      </button>
    </div>

    <!--* SEARCH -->
    <div class="flex relative mt-1.5 mb-2 ml-2 h-8" class:hidden={small}>
      <svg
        class="absolute left-2 top-2 text-gray-400"
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
          ><circle cx="11" cy="11" r="8" /><path d="m21 21l-4.3-4.3" /></g
        ></svg
      >
      <input
        style="background-color: rgba(17, 24, 39, 0.6);"
        class="w-full placeholder-gray-400 rounded-md pl-8 focus:outline-none shadow-xl font-normal text-[0.9rem]"
        placeholder="Search"
        autocapitalize="off"
        autocomplete="off"
        bind:this={inputRef}
        bind:value={searchQuery}
        on:input={handleInput}
        on:keydown={handleEnter}
      />
    </div>
  </div>

  <div class="overflow-y-auto flex-grow no-scrollbar">
    {#if sidebar_mode === 0}
      <SidebarUsers {small} {fillSearchBoxAdd} />
    {:else if sidebar_mode === 1}
      <SidebarServers />
    {:else if sidebar_mode === 2}
      <SidebarSettings />
    {:else if sidebar_mode === 3}
      <p class="text-md inline-block ml-3">Add Friend</p>
      {#if foundUser}
        <div
          class="mt-1 h-[34px] items-center text-[0.9rem] p-[0.3rem] m-0 rounded-md focus:list-none focus:outline-none outline-none hover:bg-[rgba(0,0,0,0.125)] focus:bg-[rgba(0,0,0,0.125)] ml-2 flex"
        >
          <svg
            class="mr-2"
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
              ><path d="M18 21a6 6 0 0 0-12 0" /><circle
                cx="12"
                cy="11"
                r="4"
              /><rect width="18" height="18" x="3" y="3" rx="2" /></g
            ></svg
          >
          {foundUser.username}
          <button
            on:click={() => {
              SendOffer(foundUser.ip);
            }}
            class="ml-auto hover:cursor-pointer"
            ><svg
              xmlns="http://www.w3.org/2000/svg"
              width="1.25em"
              height="1.25em"
              viewBox="0 0 24 24"
              fill="currentColor"
              ><path
                d="M11 11V7H13V11H17V13H13V17H11V13H7V11H11ZM12 22C6.47715 22 2 17.5228 2 12C2 6.47715 6.47715 2 12 2C17.5228 2 22 6.47715 22 12C22 17.5228 17.5228 22 12 22ZM12 20C16.4183 20 20 16.4183 20 12C20 7.58172 16.4183 4 12 4C7.58172 4 4 7.58172 4 12C4 16.4183 7.58172 20 12 20Z"
              ></path></svg
            ></button
          >
        </div>
      {:else}
        <p class="ml-4 mt-1 text-gray-600 text-sm">Enter to search</p>
      {/if}
    {:else}
      <h1>error</h1>
    {/if}
    <!-- <h1>{searchResult}</h1> -->
  </div>
</div>
