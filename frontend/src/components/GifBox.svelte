<script>
    import { onMount } from 'svelte';

    export let mb;
    export let chat_id;
    export let sendMessage;
    export let toggleShow;

    const tenor = "AIzaSyDdux0EuzTX4pBG4H-il_X4xbe3p2BPOgg"; // Tenor said it can be used client side ¯\_(ツ)_/¯
    let searchTerm = "";

    let topGifsLeft = [];
    let topGifsRight = [];

    async function searchGif(searchTerm) {
        const clientkey = "Sept";
        const lmt = 8;

        const searchUrl = `https://tenor.googleapis.com/v2/search?q=${searchTerm}&key=${tenor}&client_key=${clientkey}&limit=${lmt}`;

        try {
            const response = await fetch(searchUrl);
            const data = await response.json();

            // Update the GIF URLs using the response data
            let topGifs = data.results;
            topGifsLeft = topGifs.filter((_, index) => index % 2 === 0); // Even indices
            topGifsRight = topGifs.filter((_, index) => index % 2 !== 0); // Odd indices
        } catch (error) {
            console.error("Error fetching data from Tenor API:", error);
        }
    }

    async function getFeaturedGifs() {
        const clientkey = "Sept";
        const lmt = 8;

        const featuredUrl = `https://tenor.googleapis.com/v2/featured?key=${tenor}&client_key=${clientkey}&limit=${lmt}`;

        try {
            const response = await fetch(featuredUrl);
            const data = await response.json();

            let topGifs = data.results;
            topGifsLeft = topGifs.filter((_, index) => index % 2 === 0); // Even indices
            topGifsRight = topGifs.filter((_, index) => index % 2 !== 0); // Odd indices
        } catch (error) {
            console.error("Error fetching featured GIFs:", error);
        }
    }

    function sendGif(giflink) {
        sendMessage(`<img class="mt-1.5 mb-1 rounded-md max-h-72 max-w-72" src="${giflink}" alt="gif" />`, chat_id);
        toggleShow();
    }

    function handleInput(event) {
        if (searchTerm.length > 2) {
            searchGif(searchTerm);
        }
    }

    onMount(() => {
        getFeaturedGifs();
    });
</script>

<div class="h-64 w-72 p-2 rounded-md shadow-lg absolute left-[40vw]" style="bottom:calc({mb}px + 1rem); background-color: rgba(17, 24, 39, 0.98);">
    <div class="flex relative mb-2 h-8">
        <svg class="absolute left-2 top-2 text-gray-400" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8" /><path d="m21 21l-4.3-4.3"/></g></svg>
        <input
            style="background-color: rgba(17, 24, 39, 0.6);"
            class="w-full placeholder-gray-400 rounded-md pl-8 focus:outline-none shadow-xl font-normal text-[0.9rem]"
            placeholder="Search Tenor"
            autocapitalize="off"
            autocomplete="off"
            bind:value={searchTerm}
            on:input={handleInput}
        />
    </div>

    <div class="overflow-y-auto w-full relative h-[calc(100%-2.5rem)] flex gap-2 no-scrollbar">
        <div class="w-1/2">
            {#each topGifsLeft as gif}
                <img
                    on:click={() => sendGif(gif.media_formats?.gif?.url)}
                    class="mb-2 w-full rounded-md shadow-lg hover:cursor-pointer hover:opacity-65"
                    src={gif.media_formats?.nanogif?.url}
                    alt="Preview GIF"
                />
            {/each}
        </div>

        <div class="w-1/2">
            {#each topGifsRight as gif}
            <img
                on:click={() => sendGif(gif.media_formats?.gif?.url)}
                class="mb-2 w-full rounded-md shadow-lg hover:cursor-pointer hover:opacity-65"
                src={gif.media_formats?.nanogif?.url}
                alt="Preview GIF"
            />
            {/each}
        </div>
    </div>
</div>
