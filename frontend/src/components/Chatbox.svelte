<script lang="ts">
  import { onMount, afterUpdate } from "svelte";
  import { EventsOn } from "../../wailsjs/runtime/runtime";
  import { GetMessagesByChatID } from "../../wailsjs/go/controllers/MessageController";
  import { GetUser } from "../../wailsjs/go/controllers/UserController";
  import { models } from "../../wailsjs/go/models";
  import { message_store } from "../stores/messageStore.js";
  import { selection_store } from "../stores/selectionStore.js";
  import Message from "./Message.svelte";
  import Header from "./Header.svelte";

  let chatbox;
  let participants: Map<string, models.User> = new Map(); // Caching user details

  function getUserDetails(id: string) {
    if (participants.has(id)) {
      // Check if user details are already in cache
      return Promise.resolve(participants.get(id)); // Return cached user wrapped in a resolved promise
    }

    return GetUser(id)
      .then((user: models.User) => {
        participants.set(id, user); // Store user details in cache
        return user;
      })
      .catch((err) => {
        console.error("Error getting user: ", err);
      });
  }

  async function getMessages() {
    let chatId: string = $selection_store.id;
    GetMessagesByChatID(chatId)
      .then(async (messages: models.Message[]) => {
        const userIds = new Set(messages.map((message) => message.user_id));
        await Promise.all(
          Array.from(userIds).map((userId) => getUserDetails(userId)),
        );

        message_store.set(messages);
      })
      .catch((err) => {
        console.error("Failed to populate chatbox with messages:", err);
      });
  }

  onMount(() => {
    getMessages();
    EventsOn("message:new", () => {
      getMessages();
      console.log("message from wails event");
    });
  });

  afterUpdate(() => {
    if (!chatbox) {
      return;
    }
    // if new message is from sender, snap to bottom.
    // if new message is received, and already at bottom snap to new bottom.
    // if new message is received, and scrolled up, don't snap to bottom.
    chatbox.scrollTop = chatbox.scrollHeight;
  });

  $: $selection_store, getMessages(); // Trigger getMessages whenever selection_store changes
</script>

<div
  style="background-color: rgba(17, 24, 39, 0.5);"
  class="flex h-full flex-col rounded-md"
>
  <Header recipient={"Some Person"} />

  <div
    bind:this={chatbox}
    class="flex-grow h-20 overflow-auto scrollbar-chatbox flex flex-col-reverse shadow-xl"
  >
    <div class="p-2 pr-[0.1rem] pointer-events-none">
      {#each $message_store as message, i}
        <Message
          content={message.content}
          created_at={message.created_at}
          last_message_create_at={i > 0
            ? $message_store[i - 1].created_at
            : null}
          user_id={message.user_id}
          last_sender_user_id={i > 0 ? $message_store[i - 1].user_id : null}
          index={i}
          username={participants.get(message.user_id)?.username ?? "Unknown"}
          avatar={participants.get(message.user_id)?.avatar ??
            "default-avatar-url"}
        />
      {/each}
    </div>
  </div>
</div>
