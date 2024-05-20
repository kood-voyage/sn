<script lang="ts">
	import Message from './message.svelte';
	import Header from './header.svelte';
	import type { PageData } from './$types';
	import Editor from './editor.svelte';
	import { afterUpdate, onMount } from 'svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { PaperPlane } from 'svelte-radix';
	import { sendMessageByWebsocket, sendMessageTo } from './send-message';
	import type { DisplayData } from './+page';
	import { currentUserStore } from '$lib/store/user-store';
	import type { ChatLine } from '$lib/client/api/chat-requests';
	import { messageStore } from '$lib/store/websocket-store';

	export let data: PageData;

	let user: DisplayData;
	let chat_lines: ChatLine[] = [];
	$: if (data.ok) {
		user = data.chatData.display_data;
		chat_lines = data.chatData.lines_data;
		if (chat_lines && chat_lines.length != 0)
			chat_lines = chat_lines.sort((linex, liney) => {
				if (linex.created_at > liney.created_at) {
					return 1;
				} else if (linex.created_at == liney.created_at) return 0;
				return -1;
			});
		messageStore.set(chat_lines);
	}

	type userAllocate = { [key: string]: DisplayData };

	const user_allocation: userAllocate = {};

	$: if (user) {
		user_allocation[$currentUserStore.id] = {
			chat_id: user.chat_id,
			id: $currentUserStore.id,
			username: $currentUserStore.username,
			cover:
				$currentUserStore.cover || 'https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Error',
			avatar:
				$currentUserStore.avatar || 'https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Error'
		};
		user_allocation[user.id] = user;
	}

	let editorContent: string;
	let scrollContainer: HTMLDivElement;
	function scrollToBottom() {
		if (scrollContainer) scrollContainer.scrollTop = scrollContainer.scrollHeight;
	}

	onMount(() => {
		messageStore.subscribe((value) => {
			chat_lines = value;
		});
		scrollToBottom();
	});
	afterUpdate(scrollToBottom);
</script>

<div class="container w-full p-0 m-0">
	{#if user != undefined}
		<Header {user} />
	{:else}
		<div class="h-12 bg-neutral-50 dark:bg-neutral-800 px-4 w-full">
			<div class="flex p-2 h-full">
				<img
					src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Error"
					alt="avatar"
					class="w-8 rounded-full mr-2"
				/>
				<p class="h-full text-center">User Not Found</p>
			</div>
		</div>
	{/if}

	{#if user}
		<div bind:this={scrollContainer} class="messages flex flex-col m-2 overflow-y-scroll h-full">
			{#if chat_lines}
				{#each chat_lines as line (line)}
					<Message user={user_allocation[line.user_id]} time={line.created_at} msg={line.message} />
				{/each}
			{/if}
		</div>
		<div style="align-items: center;" class="flex flex-row w-full p-4">
			<Editor bind:editorContent />
			<Button
				on:click={async () => {
					if (!user.chat_id || !user.id) {
						alert(
							'Something Went Wrong - Contact Customer Support and Ask For A Refund! \n Error: user_id or chat_id Not Found'
						);
						return;
					}
					if (editorContent == '' || typeof editorContent != 'string') {
						alert(
							'Nah no empty messages please, not enough money to send all of these requests! 🥺🙇'
						);
						return;
					}

					if (editorContent) {
						const messageResp = await sendMessageTo(
							editorContent,
							user.chat_id,
							$currentUserStore.id
						);

						if (!messageResp.ok) {
							alert('Something went wrong sending the message! >>> ' + messageResp.message);
							return;
						}
						console.log('messageResp >>> ', messageResp);
						messageStore.update((old) => {
							old.push(messageResp.chatLine);
							return old;
						});
						// console.log(messageResp.chatLine instanceof ChatLine);
						sendMessageByWebsocket(user.id, $currentUserStore.id, messageResp.chatLine);
					}
					editorContent = '';
				}}
			>
				<PaperPlane />
			</Button>
		</div>
	{:else}
		<p>COULD NOT FIND THE CHAT</p>
		<p>Try refreshing the page!</p>
	{/if}
</div>

<style>
	/* This is the outer container */
	div.container {
		display: flex;
		flex-direction: column;
		height: 100vh; /* Set the height of the container to the full viewport height */
		min-width: -webkit-fill-available;
	}

	/* Messages fill the remaining space */
	.messages {
		flex: 1; /* This will make the messages section expand to fill the space */
		overflow-y: auto; /* Enables scrolling */
	}
</style>
