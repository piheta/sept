<script>
    let selection = "";

    function setSelection(obj) {
      selection = obj;
    }

    let settings = ["General", "User", "Audio", "Proxy", "Bots", "Theme"]
    setSelection(settings[0])

    function handleKeyDown(event, setting) {
      if (event.key === 'Enter' || event.key === ' ') {
          setSelection(setting);
      }
  }
</script>

<div class="h-full">
    <!-- SETTINGS TEXT -->
    <div class="flex items-center align-middle text-center mt-2 pt-0 pr-20 pb-1 pl-2 select-none">
      <p class="text-lg inline-block ml-1">
        Settings
      </p>
    </div>

    <ul class="cursor-pointer">
        {#each settings as setting, index }
        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
        <li 
        class="text-[0.9rem] p-[0.3rem] m-0 rounded-md outline-none hover:bg-[rgba(0,0,0,0.125)] focus:bg-[rgba(0,0,0,0.125)]"
        class:!bg-gray-900={setting === selection}
        class:!mb-3={index === settings.length - 1}
        tabindex="0"
        on:click={() => setSelection(setting)}
        on:keydown={(event) => handleKeyDown(event, setting)}
      >
          <div class="flex h-6 items-center pl-2">
            {setting}
          </div>
        </li>
        {/each}
    </ul>
</div>
