<script lang="ts">
  export let email: string;
</script>

<div class="flex relative mb-2 mt-2 h-8">
  <svg
    class="absolute left-2 top-2 text-gray-500"
    xmlns="http://www.w3.org/2000/svg"
    width="1em"
    height="1em"
    viewBox="0 0 24 24"
    ><g
      fill="none"
      stroke="currentColor"
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      ><rect width="20" height="16" x="2" y="4" rx="2" /><path
        d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"
      /></g
    ></svg
  >
  <input
    id="user-email"
    bind:value={email}
    type="email"
    class="w-full bg-gray-900 text-white placeholder-gray-500 rounded-md pl-8 focus:outline-none"
    placeholder="Email"
  />
</div>
