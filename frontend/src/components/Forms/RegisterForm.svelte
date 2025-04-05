<script lang="ts">
  export let username = "";
</script>

<div class="flex relative mb-2 h-8">
  <svg
    class="absolute left-2 top-2 text-gray-500"
    xmlns="http://www.w3.org/2000/svg"
    width="1.1em"
    height="1.1em"
    viewBox="0 0 24 24"
    ><g
      fill="none"
      stroke="currentColor"
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      ><path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" /><circle
        cx="12"
        cy="7"
        r="4"
      /></g
    ></svg
  >
  <input
    id="username"
    bind:value={username}
    type="text"
    class="w-full bg-gray-900 text-white placeholder-gray-500 rounded-md pl-8 focus:outline-none"
    placeholder="Username"
  />
</div>
