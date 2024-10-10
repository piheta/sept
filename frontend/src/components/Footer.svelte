<script>
    import { SendMessage } from "../../wailsjs/go/main/App.js";
    import { message_store } from '../stores/messageStore.js';
    import { selection_store } from '../stores/selectionStore.js';
    import EmoteBox from "./EmoteBox.svelte";
    import GifBox from "./GifBox.svelte";

    export let recipient;
    export let height;
    let input_txt = "";
    const lineHeight = 20; // Height to add for each new line after the first
    let originalHeight = height; // Store the original height passed from the parent
    const maxFooterHeight = () => window.innerHeight / 2; // Maximum footer height

    let showGifBox = false;
    let showEmoteBox = false;


    function toggleShowGifBox(){showGifBox = !showGifBox}
    function toggleShowEmoteBox(){showEmoteBox = !showEmoteBox}

    function sendMessage(message, chat_id) {
        if (message.length < 1) return;

        SendMessage(message, chat_id).then((allMessagesInChat) => {
            message_store.set(allMessagesInChat);
            input_txt = "";
            adjustHeight();
            
        }).catch((err) => {
            console.error("failed to send message, ", err)
        });
    }
        

    function handleKeyDown(event) {
        if (event.key === 'Enter' && !event.shiftKey) {
            event.preventDefault();
            sendMessage(input_txt.trim(), $selection_store.id);
            height = originalHeight; // Reset height to the original after sending
        }
    }

    // Interpolates towards the target height
    function interpolateHeight(currentHeight, targetHeight, duration = 50) {
        const start = performance.now();
        function step(timestamp) {
            const progress = Math.min((timestamp - start) / duration, 1); // Ensure progress is capped at 1
            const interpolatedHeight = currentHeight + (targetHeight - currentHeight) * progress;
            height = Math.min(interpolatedHeight, maxFooterHeight()); // Make sure we don't exceed the max height
            if (progress < 1) {
                requestAnimationFrame(step); // Continue animating until progress reaches 1
            }
        }
        requestAnimationFrame(step);
    }

    function adjustHeight() {
        const newLines = (input_txt.match(/\n/g) || []).length;
        const targetHeight = originalHeight + Math.max(0, newLines - 1) * lineHeight;
        const currentHeight = height;
        interpolateHeight(currentHeight, targetHeight); // Smoothly interpolate between current and target height
    }

    // Watch for input changes and recalculate height dynamically
    $: adjustHeight(); // Recalculate height whenever input_txt or height changes
</script>

{#if showGifBox}
    <GifBox mb={height} chat_id={$selection_store.id} sendMessage={sendMessage} toggleShow={toggleShowGifBox} />
{/if}
{#if showEmoteBox}
    <EmoteBox mb={height} chat_id={$selection_store.id} sendMessage={sendMessage} toggleShow={toggleShowEmoteBox} />
{/if}

<div class="w-full rounded-md p-2 flex flex-col shadow-xl scrollbar-chatbox" style="height:{height}px; background-color: rgba(17, 24, 39, 0.5);">
    <div class="flex -ml-0.5 *:mr-1 mb-0.5">
        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" fill="currentColor"><path d="M8 11H12.5C13.8807 11 15 9.88071 15 8.5C15 7.11929 13.8807 6 12.5 6H8V11ZM18 15.5C18 17.9853 15.9853 20 13.5 20H6V4H12.5C14.9853 4 17 6.01472 17 8.5C17 9.70431 16.5269 10.7981 15.7564 11.6058C17.0979 12.3847 18 13.837 18 15.5ZM8 13V18H13.5C14.8807 18 16 16.8807 16 15.5C16 14.1193 14.8807 13 13.5 13H8Z"></path></svg>
        <svg xmlns="http://www.w3.org/2000/svg" class="-ml-0.5 hover:bg-gray-200 hover:bg-opacity-5" width="1em" height="1em" viewBox="0 0 24 24" fill="currentColor"><path d="M15 20H7V18H9.92661L12.0425 6H9V4H17V6H14.0734L11.9575 18H15V20Z"></path></svg>
        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" fill="currentColor"><path d="M8 3V12C8 14.2091 9.79086 16 12 16C14.2091 16 16 14.2091 16 12V3H18V12C18 15.3137 15.3137 18 12 18C8.68629 18 6 15.3137 6 12V3H8ZM4 20H20V22H4V20Z"></path></svg>
        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 16 16"><path fill="white" d="M9.905 2.815a.75.75 0 0 1 .38.99l-4 9a.75.75 0 1 1-1.37-.61l4-9a.75.75 0 0 1 .99-.38M4.498 5.19a.75.75 0 0 1 .063 1.058L3.003 8l1.558 1.752a.75.75 0 1 1-1.122.996l-2-2.25a.75.75 0 0 1 0-.996l2-2.25A.75.75 0 0 1 4.5 5.19m7.004 0a.75.75 0 0 1 1.059.062l2 2.25a.75.75 0 0 1 0 .996l-2 2.25a.75.75 0 0 1-1.122-.996L12.996 8L11.44 6.248a.75.75 0 0 1 .063-1.058"/></svg>
        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" fill="currentColor"><path d="M8 4H21V6H8V4ZM4.5 6.5C3.67157 6.5 3 5.82843 3 5C3 4.17157 3.67157 3.5 4.5 3.5C5.32843 3.5 6 4.17157 6 5C6 5.82843 5.32843 6.5 4.5 6.5ZM4.5 13.5C3.67157 13.5 3 12.8284 3 12C3 11.1716 3.67157 10.5 4.5 10.5C5.32843 10.5 6 11.1716 6 12C6 12.8284 5.32843 13.5 4.5 13.5ZM4.5 20.4C3.67157 20.4 3 19.7284 3 18.9C3 18.0716 3.67157 17.4 4.5 17.4C5.32843 17.4 6 18.0716 6 18.9C6 19.7284 5.32843 20.4 4.5 20.4ZM8 11H21V13H8V11ZM8 18H21V20H8V18Z"></path></svg>
        <div class="bg-gray-200 bg-opacity-10 w-[1px]"></div>
        <button on:click={toggleShowEmoteBox}><svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="white" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 14s1.5 2 4 2s4-2 4-2M9 9h.01M15 9h.01"/></g></svg></button>
        <button on:click={toggleShowGifBox}><svg xmlns="http://www.w3.org/2000/svg" style="margin-top: -2px;" width="1.2em" height="1.2em" viewBox="0 0 16 16"><path fill="currentColor" d="M5.052 6.706c.481-.05.853.037.986.103a.5.5 0 1 0 .447-.894c-.351-.176-.928-.267-1.537-.203c-.96.1-1.948.934-1.948 2.297c0 1.385 1.054 2.3 2.3 2.3c.58 0 1.1-.272 1.397-.553c.262-.248.303-.577.303-.783v-.964a.5.5 0 0 0-.5-.5h-.807a.5.5 0 1 0 0 1H6v.464a.4.4 0 0 1-.006.071a1.13 1.13 0 0 1-.694.265c-.731 0-1.3-.505-1.3-1.3c0-.818.567-1.252 1.052-1.303M9 6.21a.5.5 0 0 0-1 0v3.6a.5.5 0 1 0 1 0zm1.5-.5a.5.5 0 0 0-.5.5v3.6a.5.5 0 0 0 1 0V8.506l1.003-.006a.5.5 0 0 0-.006-1L11 7.506v-.797h1.5a.5.5 0 0 0 0-1zM3.5 2A2.5 2.5 0 0 0 1 4.5v7A2.5 2.5 0 0 0 3.5 14h9a2.5 2.5 0 0 0 2.5-2.5v-7A2.5 2.5 0 0 0 12.5 2zM2 4.5A1.5 1.5 0 0 1 3.5 3h9A1.5 1.5 0 0 1 14 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11.5z"/></svg></button>

        <svg xmlns="http://www.w3.org/2000/svg" class="ml-auto" width="1em" height="1em" viewBox="0 0 24 24"><path fill="none" stroke="white" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18m-2 0v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6m3 0V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2m-6 5v6m4-6v6"/></svg>
        <svg xmlns="http://www.w3.org/2000/svg" class="!mr-0" width="1em" height="1em" viewBox="0 0 24 24" fill="currentColor"><path d="M21.7267 2.95694L16.2734 22.0432C16.1225 22.5716 15.7979 22.5956 15.5563 22.1126L11 13L1.9229 9.36919C1.41322 9.16532 1.41953 8.86022 1.95695 8.68108L21.0432 2.31901C21.5716 2.14285 21.8747 2.43866 21.7267 2.95694ZM19.0353 5.09647L6.81221 9.17085L12.4488 11.4255L15.4895 17.5068L19.0353 5.09647Z"></path></svg>
    </div>
    <textarea
        bind:value={input_txt}
        style="background-color: rgba(0, 0, 0, 0);"
        class="w-full resize-none m-0 mt-1 focus:outline-none active:outline-none flex-grow text-[0.95rem] scrollbar-chatbox"
        placeholder="Message {$selection_store.name}"
        on:keydown={handleKeyDown}
        on:input={adjustHeight}
    />
</div>