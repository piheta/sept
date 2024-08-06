<script>
    export let recipient;
    let input_txt = "";
</script>

<div class="h-full w-full rounded-md bg-gray-900 p-1">
    <div class="">hi</div>
    <textarea
        bind:value={input_txt}
        class="w-full bg-gray-900 resize-none m-0 focus:outline-none active:outline-none"
        placeholder="Message {recipient}"
    />
    <div class="m-0 flex justify-end">
        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 2l-7 20l-4-9l-9-4Zm0 0L11 13"/></svg>
    </div>
</div>

<style>
</style>
