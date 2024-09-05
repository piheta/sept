<!-- Messages.svelte -->
<script>
    import { message_store } from '../stores/messageStore.js';
</script>

<div class="bg-gray-700 flex-grow rounded-md mt-2 overflow-auto">
    {$message_store}
</div>