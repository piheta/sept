<script lang="ts">
    let {
        content,
        created_at,
        last_message_create_at,
        user_id,
        last_sender_user_id,
        index,
        username,
        avatar
    } = $props();

    let isDifferentSender = last_sender_user_id !== user_id
    let isOlderThanOneMinute = (index > 0 && (new Date(created_at).getTime() - new Date(last_message_create_at).getTime()) > 1 * 60 * 1000)
</script>

{#if isDifferentSender || isOlderThanOneMinute}
    <!-- (Big Message) -->
    <div class={`flex ${index === 0 ? '' : 'mt-4'} `}> <!-- first element should not have mt-4 -->
        <img class="w-10 min-w-10 h-10 rounded-md pointer-events-none select-none object-cover object-top" alt="sender" src={avatar} />
        <div class="flex flex-col justify-between ml-2">
            <h1 class="leading-none">
                <span class="font-semibold hover:underline pointer-events-auto cursor-pointer">{username ?? 'Unknown'}</span>
                <span class="text-gray-400 text-[0.7rem] ml-2">
                    {created_at ? new Date(created_at).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) : 'Unknown Time'}
                </span>
            </h1>
            <p class="select-text pointer-events-auto text-gray-200 text-[0.95rem] mt-0.5 break-words overflow-auto max-w-full" class:break-all={!content.includes(" ")}>
                {@html content?.replace(/\n/g, '<br>') ?? 'No content available'}
            </p>
        </div>
    </div>
    
    {:else}
    <!-- Show only the message content (Small Message) -->
    <div class="flex-col pl-12 pointer-events-auto relative group">
        <span class="text-gray-400 text-[0.7rem] absolute left-1 mt-[0.25rem] opacity-0 group-hover:opacity-100 transition-opacity pointer-events-none">
            {created_at ? new Date(created_at).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) : ''}
        </span>
        <p class="select-text text-gray-200 text-[0.95rem]">{@html content?.replace(/\n/g, '<br>') ?? 'No content available'}</p>
    </div>
    
{/if}