<script>
  import { GetServers } from '../../../wailsjs/go/main/App.js';
  let servers = [];
  let selection = { type: 0, value: "" };

  async function getServers() {
      servers = await GetServers();
      if (servers.length > 0) {
          setSelection({ type: 0, value: servers[0] });
      }
  }

  function setSelection(obj) {
      selection = obj;
  }

  getServers();

  function handleKeyDown(event, server) {
      if (event.key === 'Enter' || event.key === ' ') {
          setSelection({ type: 0, value: server });
      }
  }
</script>

<div class="h-full">
  <!-- SERVERS TEXT -->
  <div class="flex items-center align-middle text-center mt-2 pt-0 pr-20 pb-1 pl-2 select-none">
      <p class="text-lg inline-block ml-1">
          Servers
      </p>
  </div>

  <ul class="cursor-pointer">
      {#each servers as server, index}
      <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
      <li 
        class="text-[0.9rem] p-[0.3rem] m-0 rounded-md outline-none hover:bg-[rgba(0,0,0,0.125)] focus:bg-[rgba(0,0,0,0.125)]"
        class:!bg-[rgba(0,0,0,0.3)]={server === selection.value}
        class:!mb-3={index === servers.length - 1}
        tabindex="0"
        on:click={() => setSelection({ type: 0, value: server })}
        on:keydown={(event) => handleKeyDown(event, server)}
      >
        <div style="display: flex; height:24px; align-items:center;">
            <img class="h-full mr-2" alt="" src="data:image/webp;base64,UklGRqwCAABXRUJQVlA4WAoAAAAgAAAALwAALwAAVlA4II4CAABQDACdASowADAAPp08mEiloyIhMdmYALATiWUAwg+qCw7y5hqVSjuVyysub2NO2tXaUQAUAknl1FDLMYZbeoiHFSK0B9d0AxYLOmOKH+ZNstOIyM8e33IuMpsf2RgYx2kUUqtrEGfU8VhtgAD+/hV6l0JApk9Mc5vqA6UJSKhgbJFB9RPlLlDlnaYRCHX0Lf+LPZIU7OD94hoH4JS5vD6dxw+cT7cNcUSGbLKefrR7Er+H6qmgJcyiWWt3JslzDOtJfk9a7KYxmZs60KI5EFzgAfxB8AZrIgDJtJa+Qt2A+zE4eIeYd+Xj0Xivdi6WGWXzK/0WPoXPUMtR4d2u8iKsd1TqGrzS1Ihj0sd97t78WWB7ESmF1DSur0DiGvEUmBKWtZvAkIYIvoKjb3AmQ9g0m+baQpdWDMHdZw1h0nGSbdR+VjlwOsSzc973HrUV9xrpYJsKFO8480CPNJDpi9fSxerx6bHKYSnTHjPf98QrVYfSxddTjdcA7kC5xk4aFl57IPqLTY78hVy3QBUON8FmwNB/ANm3hlLwpKbYYUugM+HrUHSxIKCRrihS91UWS7o6yBbYwFRWeK96I2972upHLgZvnweatHDwZ9ogY7EEnnNYmVps4Gd6buXrCTIXqaWNLti/rp06uGNs3hljQFAOamJ2/F143nMdbf86maAg3lotDygFLGuv0IeZrWHLEmSl7Cx4TUJWjsnLpSLUjSv3wf88qjwzBOHr5qPKfoh+53JUFRBeSlPWIiwK0PRwP+uupoXHTx8m2d9fljAumKXdWyRebm23ifAD0xCvy6u6YPav9ePOyXf6kV38baVu7LyKECBN+KsCIB4pLK10mDLeFCJ3JFwdOm2weAnyfAA=" />
            {server}
        </div>
      </li>
      {/each}
  </ul>
</div>