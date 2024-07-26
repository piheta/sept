<script>
  import List from "./List.svelte";
  import ListElement from "./ListElement.svelte";

    let selection = "";

    function setSelection(obj) {
      selection = obj;
    }

    let settings = ["General", "User", "Audio", "Proxy", "Bots", "Theme"]
    setSelection(settings[0])

</script>


<List title="Settings">
  {#each settings as setting, index }
    <ListElement 
      item={setting}
      img={false}
      selected={setting === selection}
      isLast={index === settings.length - 1}
      on:select={() => {setSelection(setting)}}
    />
  {/each}
</List>
