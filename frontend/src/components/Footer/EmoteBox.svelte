<script lang="ts">
    import { onMount } from 'svelte';

    let {
        mb,
        chat_id,
        sendMessage,
        toggleShow
    } = $props();

    let searchTerm = $state('');
    let emotes = $state([]);

    const emoteImports = import.meta.glob('../../assets/emotes/*.webp', { eager: true });

    onMount(() => {
        emotes = Object.entries(emoteImports).map(([path, module]) => {
            const name = path.split('/').pop().replace('.webp', '');
            return { url: module.default, name };
        });
    });

    const sendEmote = (url) => {

        sendMessage(`<img class="mt-1.5 mb-1 rounded-md max-h-10 max-w-15" src="${url}" alt="gif" />`, chat_id);
        toggleShow(false);
    };

    let filteredEmotes = $derived(searchTerm
        ? emotes.filter((emote) => emote.name.toLowerCase().includes(searchTerm.toLowerCase()))
        : emotes);
</script>

<div class="h-64 w-72 p-2 rounded-md shadow-lg absolute left-[40vw]" style="bottom:calc({mb}px + 1rem); background-color: rgba(17, 24, 39, 0.98);">
    <div class="flex relative mb-2 h-8">
        <svg class="absolute left-2 top-2 text-gray-400" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8" /><path d="m21 21l-4.3-4.3"/></g></svg>
        <input
            style="background-color: rgba(17, 24, 39, 0.6);"
            class="w-full placeholder-gray-400 rounded-md pl-8 focus:outline-none shadow-xl font-normal text-[0.9rem]"
            placeholder="Search Emotes"
            autocapitalize="off"
            autocomplete="off"
            bind:value={searchTerm}
        />
    </div>

    <div class="overflow-y-auto w-full grid grid-cols-[repeat(auto-fill,minmax(1.5rem,1fr))] gap-1 no-scrollbar">
        {#each filteredEmotes as emote}
            <img
                onclick={() => sendEmote(emote.url)}
                class="w-6 h-6 hover:cursor-pointer hover:opacity-65"
                src={emote.url}
                alt={emote.name}
            />
        {/each}
    </div>
</div>