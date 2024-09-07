<script>
    export let content;
    export let created_at;
    export let last_message_create_at
    export let index;
    export let username;
    export let avatar;
</script>

{#if index === 0 || (new Date(created_at).getTime() - new Date(last_message_create_at).getTime()) > 1 * 60 * 1000} <!-- Show image and timestamp if more than 1 minute have passed since the previous message -->
    <div class={`flex ${index === 0 ? 'mb-1' : 'mt-4 mb-1'}`}> <!-- first element should not have mt-4 -->
        <!-- svelte-ignore a11y-missing-attribute -->
        <img class="w-10 h-10 rounded-md pointer-events-none select-none" src={avatar} />
        <div class="flex flex-col justify-between ml-2">
            <h1 class="leading-none">
                <span class="font-semibold">{username ?? 'Unknown'}</span>
                <span class="text-gray-300 text-sm ml-2">
                    {created_at ? new Date(created_at).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) : 'Unknown Time'}
                </span>
            </h1>
            <p class={content?.includes('\n') ? '' : 'leading-none'}>
                {@html content?.replace(/\n/g, '<br>') ?? 'No content available'}
            </p>
        </div>
    </div>
    {:else}
    <!-- Show only the message content -->
    <div class="flex-col pl-12">
        <p>{@html content?.replace(/\n/g, '<br>') ?? 'No content available'}</p>
    </div>
{/if}