<script lang="ts">
    import ListElement from "./ListElement.svelte";
    import DragDropList, { VerticalDropZone, reorder } from "svelte-dnd-list";
    import { selection_store } from '../../stores/selectionStore.js';

    export let draggable;
    export let showItems = true;
    export let items = [];

    function onDrop({ detail: { from, to } }) {
        if (!to || from === to) {
            return;
        }

        items = reorder(items, from.index, to.index);
    }
</script>

<ul class="cursor-pointer">
    {#if showItems}
        <!-- todo, make this godforsaken itemsize calculate automatically -->
        <DragDropList
            id="id"
            type={VerticalDropZone}
            itemSize={34}
            itemCount={items.length}
            allowDrag={draggable}
            on:drop={onDrop}
            let:index
        >
            <ListElement
                item={items[index]}
                selected={items[index].id === $selection_store.id}
                isLast={items[index].id === items[items.length - 1].id}
            />
        </DragDropList>
    {:else if !showItems && items.some(item => item.id === $selection_store.id)}
        <ListElement
        item={$selection_store}
        selected={true}
        isLast={true}
    />
    {/if}
</ul>
