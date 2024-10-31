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
    class="text-[0.9rem] p-[0.3rem] m-0 rounded-md focus:list-none focus:outline-none outline-none hover:bg-[rgba(0,0,0,0.125)] focus:bg-[rgba(0,0,0,0.125)] ml-2"
    style={selected ? "background-color: rgba(17, 24, 39, 0.5);" : ""}
    class:!mb-3={isLast}
    tabindex="0"
    onclick={clickItem}
    onkeydown={handleKeyDown}
>
    <div class="flex h-6 items-center">
        {#if img}<img
                class="h-full mr-2 select-none rounded-sm"
                style="-webkit-user-drag: none;"
                alt=""
                src={item.avatar}
            />
        {/if}
        {item.name}
    </div>
</li>
