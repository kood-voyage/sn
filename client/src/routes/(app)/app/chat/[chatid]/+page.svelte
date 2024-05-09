<script lang="ts">
	import Message from './message.svelte';
	import Header from './header.svelte';

	import type { PageData } from './$types';
	import Editor from './editor.svelte';
	import { afterUpdate, onMount } from 'svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { PaperPlane } from 'svelte-radix';
	import { sendMessageTo } from './send-message';
	import type { DisplayData } from './+page';
	import { currentUserStore } from '$lib/store/user-store';
	import type { ChatLine } from '$lib/client/api/chat-requests';

	export let data: PageData;

	let user: DisplayData;
	let chat_lines: ChatLine[];
	$: if (data.ok) {
		console.log('USERs', user);
		console.log('USERs', chat_lines);
		user = data.chatData.display_data;
		chat_lines = data.chatData.lines_data;
		if (chat_lines && chat_lines.length != 0) chat_lines = chat_lines.sort(() => -1);
	}

	type userAllocate = { [key: string]: DisplayData };

	const user_allocation: userAllocate = {};

	$: if (user) {
		user_allocation[$currentUserStore.id] = {
			chat_id: user.chat_id,
			user_id: $currentUserStore.id,
			display_name: $currentUserStore.username,
			cover:
				$currentUserStore.cover || 'https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Error',
			avatar:
				$currentUserStore.avatar || 'https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Error'
		};
		user_allocation[user.user_id] = user;
	}

	let editorContent: string;
	// $: console.log(editorContent);

	const user1 = {
		username: 'Surinam',
		avatar:
			'https://pyxis.nymag.com/v1/imgs/630/6e0/eb215ad90cd826b9e57ff505f54c5c7228-07-avatar.1x.rsquare.w1400.jpg'
	};

	const user2 = {
		username: 'Nikita',
		avatar:
			'https://static.vecteezy.com/system/resources/thumbnails/002/002/403/small/man-with-beard-avatar-character-isolated-icon-free-vector.jpg'
	};

	const time = '11/12/04 44:30';

	const msg =
		'What do you think about our current project crisis? dasdasda sd asd asd asd askjdkjasndjasn dask jdnaskjndjkas ndkanskdjnaskj nakjsnd kandjkanskj dnasjkd nkasnd kas';

	let scrollContainer: HTMLDivElement;
	function scrollToBottom() {
		if (scrollContainer) scrollContainer.scrollTop = scrollContainer.scrollHeight;
	}

	onMount(scrollToBottom);
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
			<!-- <Message user={user1} {time} {msg} />
			<Message user={user2} {time} {msg} />
			<Message user={user1} {time} {msg} />
			<Message user={user2} {time} {msg} />
			<Message user={user1} {time} {msg} />
			<Message user={user2} {time} {msg} />
			<Message user={user1} {time} {msg} />
			<Message user={user2} {time} {msg} />
			<Message user={user1} {time} {msg} />
			<Message user={user2} {time} {msg} />
			<Message user={user1} {time} {msg} />
			<Message user={user2} {time} {msg} />
			<Message user={user1} {time} {msg} />
			<Message user={user2} {time} {msg} />
			<Message user={user1} {time} {msg} />
			<Message user={user2} {time} {msg} />
			<Message user={user1} {time} {msg} />
			<Message user={user2} {time} {msg} />
			<Message user={user1} {time} {msg} />
			<Message user={user2} {time} {msg} /> -->
			<!-- More messages -->
		</div>
		<div style="align-items: center;" class="flex flex-row w-full">
			<Editor bind:editorContent />
			<Button
				on:click={async () => {
					if (!user.chat_id || !user.user_id) {
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

					console.log(editorContent);
					if (editorContent)
						console.log(await sendMessageTo(editorContent, user.chat_id, user.user_id));
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
