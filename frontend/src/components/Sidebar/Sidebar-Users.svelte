<script>
    import { GetUsers, GetRooms } from '../../../wailsjs/go/main/App.js';
  
    let users = [];
    let rooms = [];
    let showUsers = true;
    let showRooms = true;
    let selection = { type: 0, value: "" };
  
    function toggleShowUsers() {
      showUsers = !showUsers;
    }
  
    function toggleShowRooms() {
      showRooms = !showRooms;
    }
  
    async function getUsers() {
      users = await GetUsers();
      if (users.length > 0) {
        setSelection({ type: 0, value: users[0] });
      }
    }
  
    async function getRooms() {
      rooms = await GetRooms();
    }
  
    function setSelection(obj) {
      selection = obj;
    }

    function handleKeyDown(event, type, user) {
      if (event.key === 'Enter' || event.key === ' ') {
          setSelection({ type:type, value: user});
      }
    }
      
    getUsers();
    getRooms();
  </script>

<div class="h-full">
    <!-- Friends TEXT -->
    <button 
      class="w-full flex items-center align-middle text-center mt-2 pt-0 pr-20 pb-1 pl-2 select-none cursor-pointer"
      on:click={toggleShowUsers}
    >
      <span class="inline-block transition-transform duration-150 ease-in-out" class:rotate-0={showUsers} class:-rotate-90={!showUsers}>
        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 16 16"><path fill="white" fill-rule="evenodd" d="M2.97 5.47a.75.75 0 0 1 1.06 0L8 9.44l3.97-3.97a.75.75 0 1 1 1.06 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0l-4.5-4.5a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/></svg>
      </span> 
      <p class="text-lg inline-block ml-1">
        Friends
      </p>
    </button>
      
    <ul class="cursor-pointer">
      {#if showUsers}
        {#each users as user, index }
          <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
          <li 
          class="text-[0.9rem] p-[0.3rem] m-0 rounded-md outline-none hover:bg-[rgba(0,0,0,0.125)] focus:bg-[rgba(0,0,0,0.125)]"
          class:!bg-[rgba(0,0,0,0.3)]={user === selection.value}
          class:!mb-3={index === users.length - 1}
          tabindex="0"
          on:click={() => setSelection({ type: 0, value: user })}
          on:keydown={(event) => handleKeyDown(event, 0, user)}
          >
            <div style="display: flex; height:24px; align-items:center;">
              <img class="h-full mr-2" alt="" src="data:image/webp;base64,UklGRqwCAABXRUJQVlA4WAoAAAAgAAAALwAALwAAVlA4II4CAABQDACdASowADAAPp08mEiloyIhMdmYALATiWUAwg+qCw7y5hqVSjuVyysub2NO2tXaUQAUAknl1FDLMYZbeoiHFSK0B9d0AxYLOmOKH+ZNstOIyM8e33IuMpsf2RgYx2kUUqtrEGfU8VhtgAD+/hV6l0JApk9Mc5vqA6UJSKhgbJFB9RPlLlDlnaYRCHX0Lf+LPZIU7OD94hoH4JS5vD6dxw+cT7cNcUSGbLKefrR7Er+H6qmgJcyiWWt3JslzDOtJfk9a7KYxmZs60KI5EFzgAfxB8AZrIgDJtJa+Qt2A+zE4eIeYd+Xj0Xivdi6WGWXzK/0WPoXPUMtR4d2u8iKsd1TqGrzS1Ihj0sd97t78WWB7ESmF1DSur0DiGvEUmBKWtZvAkIYIvoKjb3AmQ9g0m+baQpdWDMHdZw1h0nGSbdR+VjlwOsSzc973HrUV9xrpYJsKFO8480CPNJDpi9fSxerx6bHKYSnTHjPf98QrVYfSxddTjdcA7kC5xk4aFl57IPqLTY78hVy3QBUON8FmwNB/ANm3hlLwpKbYYUugM+HrUHSxIKCRrihS91UWS7o6yBbYwFRWeK96I2972upHLgZvnweatHDwZ9ogY7EEnnNYmVps4Gd6buXrCTIXqaWNLti/rp06uGNs3hljQFAOamJ2/F143nMdbf86maAg3lotDygFLGuv0IeZrWHLEmSl7Cx4TUJWjsnLpSLUjSv3wf88qjwzBOHr5qPKfoh+53JUFRBeSlPWIiwK0PRwP+uupoXHTx8m2d9fljAumKXdWyRebm23ifAD0xCvy6u6YPav9ePOyXf6kV38baVu7LyKECBN+KsCIB4pLK10mDLeFCJ3JFwdOm2weAnyfAA=" />
              {user}
            </div>
          </li>
        {/each}
        {:else if !showUsers && selection.type === 0}
        <li class="text-[0.9rem] p-[0.3rem] m-0 rounded-md outline-none !bg-[rgba(0,0,0,0.3)]">
          <div style="display: flex; height:24px; align-items:center;">
            <img class="h-full mr-2" alt="" src="data:image/webp;base64,UklGRqwCAABXRUJQVlA4WAoAAAAgAAAALwAALwAAVlA4II4CAABQDACdASowADAAPp08mEiloyIhMdmYALATiWUAwg+qCw7y5hqVSjuVyysub2NO2tXaUQAUAknl1FDLMYZbeoiHFSK0B9d0AxYLOmOKH+ZNstOIyM8e33IuMpsf2RgYx2kUUqtrEGfU8VhtgAD+/hV6l0JApk9Mc5vqA6UJSKhgbJFB9RPlLlDlnaYRCHX0Lf+LPZIU7OD94hoH4JS5vD6dxw+cT7cNcUSGbLKefrR7Er+H6qmgJcyiWWt3JslzDOtJfk9a7KYxmZs60KI5EFzgAfxB8AZrIgDJtJa+Qt2A+zE4eIeYd+Xj0Xivdi6WGWXzK/0WPoXPUMtR4d2u8iKsd1TqGrzS1Ihj0sd97t78WWB7ESmF1DSur0DiGvEUmBKWtZvAkIYIvoKjb3AmQ9g0m+baQpdWDMHdZw1h0nGSbdR+VjlwOsSzc973HrUV9xrpYJsKFO8480CPNJDpi9fSxerx6bHKYSnTHjPf98QrVYfSxddTjdcA7kC5xk4aFl57IPqLTY78hVy3QBUON8FmwNB/ANm3hlLwpKbYYUugM+HrUHSxIKCRrihS91UWS7o6yBbYwFRWeK96I2972upHLgZvnweatHDwZ9ogY7EEnnNYmVps4Gd6buXrCTIXqaWNLti/rp06uGNs3hljQFAOamJ2/F143nMdbf86maAg3lotDygFLGuv0IeZrWHLEmSl7Cx4TUJWjsnLpSLUjSv3wf88qjwzBOHr5qPKfoh+53JUFRBeSlPWIiwK0PRwP+uupoXHTx8m2d9fljAumKXdWyRebm23ifAD0xCvy6u6YPav9ePOyXf6kV38baVu7LyKECBN+KsCIB4pLK10mDLeFCJ3JFwdOm2weAnyfAA=" />
            {selection.value}</div>
        </li>
      {/if}
    </ul>
    

    <!-- ROOMS TEXT -->
    <button 
      class="w-full mt-1 flex items-center align-middle text-center pt-0 pr-20 pb-1 pl-2 select-none cursor-pointer"
      on:click={toggleShowRooms}
    >
      <span class="inline-block transition-transform duration-150 ease-in-out" class:rotate-0={showRooms} class:-rotate-90={!showRooms}>
        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 16 16"><path fill="white" fill-rule="evenodd" d="M2.97 5.47a.75.75 0 0 1 1.06 0L8 9.44l3.97-3.97a.75.75 0 1 1 1.06 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0l-4.5-4.5a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/></svg>
      </span> 
      <p class="text-lg inline-block ml-1">
        Rooms
      </p>
    </button>
    
    <ul class="cursor-pointer">
      {#if showRooms}
          {#each rooms as room, index }
            <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
            <li 
            class="text-[0.9rem] p-[0.3rem] m-0 rounded-md outline-none hover:bg-[rgba(0,0,0,0.125)] focus:bg-[rgba(0,0,0,0.125)]"
            class:!bg-[rgba(0,0,0,0.3)]={room === selection.value}
            class:!mb-3={index === room.length - 1}
            tabindex="0"
            on:click={() => setSelection({ type: 1, value: room })}
            on:keydown={(event) => handleKeyDown(event, 1, room)}
            >
              <div style="display: flex; height:24px; align-items:center;">
                <img class="h-full mr-2" alt="" src="data:image/webp;base64,UklGRqwCAABXRUJQVlA4WAoAAAAgAAAALwAALwAAVlA4II4CAABQDACdASowADAAPp08mEiloyIhMdmYALATiWUAwg+qCw7y5hqVSjuVyysub2NO2tXaUQAUAknl1FDLMYZbeoiHFSK0B9d0AxYLOmOKH+ZNstOIyM8e33IuMpsf2RgYx2kUUqtrEGfU8VhtgAD+/hV6l0JApk9Mc5vqA6UJSKhgbJFB9RPlLlDlnaYRCHX0Lf+LPZIU7OD94hoH4JS5vD6dxw+cT7cNcUSGbLKefrR7Er+H6qmgJcyiWWt3JslzDOtJfk9a7KYxmZs60KI5EFzgAfxB8AZrIgDJtJa+Qt2A+zE4eIeYd+Xj0Xivdi6WGWXzK/0WPoXPUMtR4d2u8iKsd1TqGrzS1Ihj0sd97t78WWB7ESmF1DSur0DiGvEUmBKWtZvAkIYIvoKjb3AmQ9g0m+baQpdWDMHdZw1h0nGSbdR+VjlwOsSzc973HrUV9xrpYJsKFO8480CPNJDpi9fSxerx6bHKYSnTHjPf98QrVYfSxddTjdcA7kC5xk4aFl57IPqLTY78hVy3QBUON8FmwNB/ANm3hlLwpKbYYUugM+HrUHSxIKCRrihS91UWS7o6yBbYwFRWeK96I2972upHLgZvnweatHDwZ9ogY7EEnnNYmVps4Gd6buXrCTIXqaWNLti/rp06uGNs3hljQFAOamJ2/F143nMdbf86maAg3lotDygFLGuv0IeZrWHLEmSl7Cx4TUJWjsnLpSLUjSv3wf88qjwzBOHr5qPKfoh+53JUFRBeSlPWIiwK0PRwP+uupoXHTx8m2d9fljAumKXdWyRebm23ifAD0xCvy6u6YPav9ePOyXf6kV38baVu7LyKECBN+KsCIB4pLK10mDLeFCJ3JFwdOm2weAnyfAA=" />
                {room}</div>
            </li>
          {/each}
      {:else if !showRooms && selection.type === 1}
           <li class="text-[0.9rem] p-[0.3rem] m-0 rounded-md outline-none !bg-[rgba(0,0,0,0.3)] ">
             <div style="display: flex; height:24px; align-items:center;">
              <img class="h-full mr-2" alt="" src="data:image/webp;base64,UklGRqwCAABXRUJQVlA4WAoAAAAgAAAALwAALwAAVlA4II4CAABQDACdASowADAAPp08mEiloyIhMdmYALATiWUAwg+qCw7y5hqVSjuVyysub2NO2tXaUQAUAknl1FDLMYZbeoiHFSK0B9d0AxYLOmOKH+ZNstOIyM8e33IuMpsf2RgYx2kUUqtrEGfU8VhtgAD+/hV6l0JApk9Mc5vqA6UJSKhgbJFB9RPlLlDlnaYRCHX0Lf+LPZIU7OD94hoH4JS5vD6dxw+cT7cNcUSGbLKefrR7Er+H6qmgJcyiWWt3JslzDOtJfk9a7KYxmZs60KI5EFzgAfxB8AZrIgDJtJa+Qt2A+zE4eIeYd+Xj0Xivdi6WGWXzK/0WPoXPUMtR4d2u8iKsd1TqGrzS1Ihj0sd97t78WWB7ESmF1DSur0DiGvEUmBKWtZvAkIYIvoKjb3AmQ9g0m+baQpdWDMHdZw1h0nGSbdR+VjlwOsSzc973HrUV9xrpYJsKFO8480CPNJDpi9fSxerx6bHKYSnTHjPf98QrVYfSxddTjdcA7kC5xk4aFl57IPqLTY78hVy3QBUON8FmwNB/ANm3hlLwpKbYYUugM+HrUHSxIKCRrihS91UWS7o6yBbYwFRWeK96I2972upHLgZvnweatHDwZ9ogY7EEnnNYmVps4Gd6buXrCTIXqaWNLti/rp06uGNs3hljQFAOamJ2/F143nMdbf86maAg3lotDygFLGuv0IeZrWHLEmSl7Cx4TUJWjsnLpSLUjSv3wf88qjwzBOHr5qPKfoh+53JUFRBeSlPWIiwK0PRwP+uupoXHTx8m2d9fljAumKXdWyRebm23ifAD0xCvy6u6YPav9ePOyXf6kV38baVu7LyKECBN+KsCIB4pLK10mDLeFCJ3JFwdOm2weAnyfAA=" />
              {selection.value}</div>
          </li>
      {/if}
    </ul>
</div>
