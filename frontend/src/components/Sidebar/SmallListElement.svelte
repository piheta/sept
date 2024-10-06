<script>
    import { selection_store } from "../../stores/selectionStore.js"

    export let item;
    export let selected = false;
    export let isLast = false;
    export let img = true;

    function handleKeyDown(event) {
        if (event.key === "Enter" || event.key === " ") {
            selection_store.set({
                id: item.id,
                name: item.name,
                avatar: item.avatar
            });
        }
    }

    function clickItem() {
        selection_store.set({
            id: item.id,
            name: item.name,
            avatar: item.avatar
         });
    }

</script>

<li
    class="p-1 m-0 rounded-lg focus:list-none focus:outline-none outline-none hover:bg-[rgba(0,0,0,0.125)] focus:bg-[rgba(0,0,0,0.125)] ml-3.5 mr-1.5 flex justify-center items-center"
    style={selected ? "background-color: rgba(17, 24, 39, 0.5);" : ""}
    class:!mb-3={isLast}
    tabindex="0"
    on:click={clickItem}
    on:keydown={handleKeyDown}
>
        {#if img}
            <img
                class="h-10 w-10 object-cover rounded-md"
                alt=""
                src={item.avatar}
            />
        {/if}
</li>
