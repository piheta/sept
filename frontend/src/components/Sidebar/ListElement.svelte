<script>
    export let item;
    export let selected = false;
    export let isLast = false;
    export let img = true;

    function handleKeyDown(event) {
        if (event.key === "Enter" || event.key === " ") {
            dispatch("select", item);
        }
    }

    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher();
</script>

<li
    class="text-[0.9rem] p-[0.3rem] m-0 rounded-md outline-none hover:bg-[rgba(0,0,0,0.125)] focus:bg-[rgba(0,0,0,0.125)]"
    class:!bg-gray-900={selected}
    class:!mb-3={isLast}
    tabindex="0"
    on:click={() => dispatch("select", item)}
    on:keydown={handleKeyDown}
>
    <div class="flex h-6 items-center">
        {#if img}<img
                class="h-full mr-2"
                alt=""
                src="data:image/webp;base64,UklGRqwCAABXRUJQVlA4WAoAAAAgAAAALwAALwAAVlA4II4CAABQDACdASowADAAPp08mEiloyIhMdmYALATiWUAwg+qCw7y5hqVSjuVyysub2NO2tXaUQAUAknl1FDLMYZbeoiHFSK0B9d0AxYLOmOKH+ZNstOIyM8e33IuMpsf2RgYx2kUUqtrEGfU8VhtgAD+/hV6l0JApk9Mc5vqA6UJSKhgbJFB9RPlLlDlnaYRCHX0Lf+LPZIU7OD94hoH4JS5vD6dxw+cT7cNcUSGbLKefrR7Er+H6qmgJcyiWWt3JslzDOtJfk9a7KYxmZs60KI5EFzgAfxB8AZrIgDJtJa+Qt2A+zE4eIeYd+Xj0Xivdi6WGWXzK/0WPoXPUMtR4d2u8iKsd1TqGrzS1Ihj0sd97t78WWB7ESmF1DSur0DiGvEUmBKWtZvAkIYIvoKjb3AmQ9g0m+baQpdWDMHdZw1h0nGSbdR+VjlwOsSzc973HrUV9xrpYJsKFO8480CPNJDpi9fSxerx6bHKYSnTHjPf98QrVYfSxddTjdcA7kC5xk4aFl57IPqLTY78hVy3QBUON8FmwNB/ANm3hlLwpKbYYUugM+HrUHSxIKCRrihS91UWS7o6yBbYwFRWeK96I2972upHLgZvnweatHDwZ9ogY7EEnnNYmVps4Gd6buXrCTIXqaWNLti/rp06uGNs3hljQFAOamJ2/F143nMdbf86maAg3lotDygFLGuv0IeZrWHLEmSl7Cx4TUJWjsnLpSLUjSv3wf88qjwzBOHr5qPKfoh+53JUFRBeSlPWIiwK0PRwP+uupoXHTx8m2d9fljAumKXdWyRebm23ifAD0xCvy6u6YPav9ePOyXf6kV38baVu7LyKECBN+KsCIB4pLK10mDLeFCJ3JFwdOm2weAnyfAA="
            />
        {/if}
        {item.name}
    </div>
</li>