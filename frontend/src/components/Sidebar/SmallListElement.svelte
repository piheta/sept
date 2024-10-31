<script lang="ts">
    import { selection_store } from "../../stores/selectionStore.js"

    interface Props {
        item: any;
        selected?: boolean;
        isLast?: boolean;
        img?: boolean;
    }

    let {
        item,
        selected = false,
        isLast = false,
        img = true
    }: Props = $props();

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
    onclick={clickItem}
    onkeydown={handleKeyDown}
>
        {#if img}
            <img
                class="h-10 w-10 object-cover rounded-md select-none"
                style="-webkit-user-drag: none;"
                alt=""
                src={item.avatar}
            />
        {/if}
</li>
