<script lang="ts">
	import { JoinGroup, type GroupJson } from '$lib/client/api/group-requests';
	import * as Dialog from '$lib/components/ui/dialog';
	import { currentUserStore } from '$lib/store/user-store';
	import type { UserType } from '$lib/types/user';
	import type { PageData } from './$types';
	import Createeventform from './createeventform.svelte';
	import GroupPostForm from './groupPostForm.svelte';
	import Namelayout from './namelayout.svelte';
	import Reactform from './reactform.svelte';

	export let data: PageData;

	let id: string, name: string, description: string, image_path: string;
	const currentUser = $currentUserStore as UserType;

	const group = data.group;
	console.log('ALL EVENTS', data);
	let errorMessage = '';
	const groupPosts = data.posts;
	let isMember = false;

	data.allevents?.forEach((event) => {
		if (event.participants) {
			event.participants.forEach((participant) => {
				if (participant.id == currentUser.id) {
					event.is_participant = true;
					event.event_status = participant.event_status;
				}
			});
		}
	});

	if (currentUser && 'id' in currentUser && group?.ok) {
		if (data.group?.ok && data.group.group.creator_id == currentUser.id) {
			isMember = true;
		}
		if (data.group?.ok)
			data.group.group.members?.forEach((user) => {
				if (user && user.id == currentUser.id) {
					isMember = true;
				}
			});
	}

	let groupInf: GroupJson;
	if (group?.ok) {
		const data = group.group;
		groupInf = group.group;
		id = data.id;
		name = data.name;
		description = data.description;
		if (data.image_path) {
			image_path = data.image_path[0];
		} else {
			image_path =
				'https://static.vecteezy.com/system/resources/previews/005/337/799/original/icon-image-not-found-free-vector.jpg';
		}
	} else {
		name = '404 Not Found';
		description = '';
		image_path =
			'https://static.vecteezy.com/system/resources/previews/005/337/799/original/icon-image-not-found-free-vector.jpg';
	}

	async function joinGroup() {
		if (group?.ok) {
			const result = await JoinGroup(group.group.name, fetch);
		}
	}
</script>

<svelte:head>
	<title>g/{name}</title>
</svelte:head>

<!-- user profile page -->

<main class="flex">
	<div class=" m-auto h-full w-full max-w-[1096px]">
		<!-- profile info header -->
		<div class="profile-info relative">
			<!-- banner img  -->
			<div class="m-auto h-32 sm:h-60 max-w-[1096px] p-0 sm:px-2">
				<img
					class="h-full w-full object-cover object-center sm:rounded-xl"
					src={image_path}
					alt="banner"
				/>
			</div>

			<div class="max-w-[1096px] sm:px-2 h-16">
				{#if data.allevents}
					{#each data.allevents as event}
						<div
							class="w-full bg-slate-200/30 p-1 mt-1 h-full flex justify-between items-center sm:rounded-xl"
						>
							<p>TITLE {event.name}</p>
							<p>Description {event.description}</p>
							<p>Created at {event.created_at}</p>
							<p>User information {event.user_information}</p>
							{#if event.participants}
								<p>Participants: {event.participants.length}</p>
							{:else}
								<p>Participants: 0</p>
							{/if}
							{#if event.is_participant}
								<p class="text-sm rounded-md px-5 p-1 m-0.5 border bg-sky-500">
									I am {event.event_status}
								</p>
							{:else}
								<Dialog.Root>
									<Dialog.Trigger class="text-sm rounded-md px-5 p-1 m-0.5 border bg-sky-500">
										React
										<Dialog.Content>
											<Reactform />
										</Dialog.Content>
									</Dialog.Trigger>
								</Dialog.Root>
							{/if}
						</div>
					{/each}
				{/if}
				<div
					class="w-full bg-slate-200/30 p-1 mt-1 h-full flex justify-between items-center sm:rounded-xl"
				>
					<div class="h-full align-middle flex-col self-start">
						<p class="md:text-xl text-lg text-ellipsis w-full bold text-left font-bold mr-2">
							{name}
						</p>
						<p class="lines3 text-sm text-left text-slate-400">{description}</p>
					</div>
					<div class="flex flex-row">
						{#if isMember}
							<Dialog.Root>
								<Dialog.Trigger class="text-sm rounded-md px-5 p-1 m-0.5 border bg-sky-500"
									>Invite user</Dialog.Trigger
								>

								<Dialog.Content>
									{#if data.group?.ok}
										{#if data.group.group.members}
											<Namelayout userList={data.group.group.members} />
										{/if}
									{:else}
										<p class="m-2">Group Info Not found, try reloading the page!</p>
									{/if}
								</Dialog.Content>
							</Dialog.Root>

							<input type="text" hidden name="target_id" value={id} />

							<Dialog.Root>
								<Dialog.Trigger class="text-sm rounded-md px-5 p-1 m-0.5 border bg-sky-500"
									>Create Post</Dialog.Trigger
								>

								<Dialog.Content>
									<!-- I tried to put 2 PROPS to this component  -->
									<!-- data : Took from previous Form, and i dont remember why it's required  -->
									<!-- groupId - this PROP i put intentionally because i think we need PARENT_ID for post in groups -->
									{#if data.group?.ok}
										<GroupPostForm data={data.form} groupId={data.group?.group.id} />
									{:else}
										<p class="m-2">Group Info Not found, try reloading the page!</p>
									{/if}
								</Dialog.Content>
							</Dialog.Root>

							<Dialog.Root>
								<Dialog.Trigger class="text-sm rounded-md px-5 p-1 m-0.5 border bg-sky-500">
									Create event
									<Dialog.Content>
										<Createeventform data={data.form} currUser={currentUser} group={groupInf} />
									</Dialog.Content>
								</Dialog.Trigger>
							</Dialog.Root>
						{:else}
							<form on:submit={joinGroup} method="post" class=" text-center">
								<p>
									{#if errorMessage}
										{errorMessage}
									{/if}
								</p>
								<input type="text" hidden name="target_id" value={id} />
								<button class="text-sm rounded-md px-5 border bg-sky-500 p-1 m-0.5" type="submit">
									Join group
								</button>
							</form>
						{/if}
					</div>
				</div>
			</div>
		</div>

		<!-- group activity / posts -->

		<div class="h-full w-full sm:grid sm:grid-cols-2 md:grid-cols-3 gap-4 p-0 sm:p-4 md:mt-80">
			{#if groupPosts}
				{#each groupPosts as post}
					<div class="bg-white rounded-lg p-4 mb-4">
						<p class="text-xl font-bold">TITLE: {post.title}</p>
						<p class="text-gray-600">CONTENT: {post.content}</p>
						<div class="flex items-center mt-2">
							<p class="text-gray-700 mr-2">User ID: {post.user_id}</p>
							<p class="text-gray-700">Created At: {post.created_at}</p>
						</div>
						<p class="text-gray-700">IMAGE IF THERE IS {post.image_path}</p>
					</div>
				{/each}
			{/if}
			<!-- <div class="bg-pink-500 h-56 w-full sm:rounded-lg">group</div>
			<div class="bg-purple-500 h-56 w-full sm:rounded-lg">group</div>
			<div class="bg-red-500 h-56 w-full sm:rounded-lg">123</div>
			<div class="bg-yellow-500 h-56 w-full sm:rounded-lg">123</div>
			<div class="bg-orange-500 h-56 w-full sm:rounded-lg">123</div> -->
		</div>
	</div>
</main>

<style>
	/* absolute */
</style>
