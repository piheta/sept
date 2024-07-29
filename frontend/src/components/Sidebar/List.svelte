<script>
    import { createEventDispatcher } from "svelte";
    import ListElement from "./ListElement.svelte";
    import { flip } from "svelte/animate";
    import { dndzone } from "svelte-dnd-action";

    export let title = "";
    export let collapsible = false;
    export let showItems = true;
    export let items = [];
    export let selectedItem = null;
    export let flipDurationMs = 300;
    let dropTargetStyle = { outline: "black solid 0px" };

    const dispatch = createEventDispatcher();

    function toggleShow() {
        if (collapsible) {
            showItems = !showItems;
        }
    }

    function handleSelect(item) {
        dispatch("select", item);
    }

    function handleDndConsider(e) {
        // dispatch('consider', e.detail);
        items = e.detail.items;
    }

    function handleDndFinalize(e) {
        // dispatch('finalize', e.detail);
        items = e.detail.items;
    }
</script>

{#if title}
    {#if collapsible}
        <button
            class="w-full flex items-center align-middle text-center mt-0 pt-0 pr-20 pb-1 pl-2 select-none cursor-pointer"
            on:click={toggleShow}
        >
            <span
                class="inline-block transition-transform duration-150 ease-in-out"
                class:rotate-0={showItems}
                class:-rotate-90={!showItems}
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="1em"
                    height="1em"
                    viewBox="0 0 16 16"
                    ><path
                        fill="white"
                        fill-rule="evenodd"
                        d="M2.97 5.47a.75.75 0 0 1 1.06 0L8 9.44l3.97-3.97a.75.75 0 1 1 1.06 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0l-4.5-4.5a.75.75 0 0 1 0-1.06"
                        clip-rule="evenodd"
                    /></svg
                >
            </span>
            <p class="text-lg inline-block ml-1">
                {title}
            </p>
        </button>
    {:else}
        <div
            class="w-full flex items-center align-middle text-center mt-0 pt-0 pr-20 pb-1 pl-2 select-none"
        >
            <p class="text-lg inline-block ml-1">
                {title}
            </p>
        </div>
    {/if}
{/if}

<ul
    class="cursor-pointer"
    use:dndzone={{ items, flipDurationMs, dropTargetStyle }}
    on:consider={handleDndConsider}
    on:finalize={handleDndFinalize}
>
    {#if showItems}
        {#each items as item (item.id)}
            <div animate:flip={{ duration: flipDurationMs }}>
                <ListElement
                    {item}
                    selected={item === selectedItem}
                    isLast={item.id === items[items.length - 1].id}
                    on:select={() => handleSelect(item)}
                />
            </div>
        {/each}
    {:else if !showItems && $$slots.selected}
        <slot name="selected"></slot>
    {/if}
</ul>
