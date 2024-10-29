<script lang="ts">
    import SmallListElement from "./SmallListElement.svelte";
    import DragDropList, { VerticalDropZone, reorder } from "svelte-dnd-list";
    import { selection_store } from '../../stores/selectionStore.js';

    export let draggable;
    export let items = [];

    function onDrop({ detail: { from, to } }) {
        if (!to || from === to) {
            return;
        }

        items = reorder(items, from.index, to.index);
    }
</script>

<div class="w-[calc(100% - 1rem)] ml-2 mb-2 h-[1px] bg-gray-200 bg-opacity-5"></div>
<ul class="cursor-pointer">
        <!-- todo, make this godforsaken itemsize calculate automatically -->
        <DragDropList
            id={"small"}
            type={VerticalDropZone}
            itemSize={50}
            itemCount={items.length}
            allowDrag={draggable}
            on:drop={onDrop}
            let:index
        >
            <SmallListElement
                item={items[index]}
                selected={items[index].id === $selection_store.id}
                isLast={items[index].id === items[items.length - 1].id}
            />
        </DragDropList>
</ul>
