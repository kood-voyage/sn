<script lang="ts">
	import PostForm from '$lib/components/post-form.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import type { GroupJson } from '$lib/server/api/group-requests';
	import { currentUserStore } from '$lib/store/user-store';
	import type { PageData } from './$types';
	import GroupPostForm from './groupPostForm.svelte';

	export let data: PageData;
	let id: string, name: string, description: string, image_path: string;
	const groupResp = data.group;
	const currentUser = $currentUserStore;
	let isMember = false;

	console.log('--------');
	console.log(groupResp.data);
	console.log('--------');

	if (currentUser && 'id' in currentUser) {
		if (groupResp.data.creator_id == currentUser.id) {
			isMember = true;
		}
		groupResp.data?.members?.forEach((user) => {
			if (user.id == currentUser.id) {
				isMember = true;
			}
		});
	}
	try {
		const data = groupResp.data as GroupJson;
		id = data.id;
		name = data.name;
		description = data.description;
		image_path = data.image_path[0];
	} catch (err) {
		console.error(err);
		name = '404 Not Found';
		description = '';
		image_path =
			'https://static.vecteezy.com/system/resources/previews/005/337/799/original/icon-image-not-found-free-vector.jpg';
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
							<form action="?/invite" method="post" class=" text-center">
								<input type="text" hidden name="target_id" value={id} />
								<button class="text-sm rounded-md px-5 border bg-sky-500 p-1 m-0.5" type="submit">
									Invite User
								</button>
							</form>
							<form>
								<input type="text" hidden name="target_id" value={id} />

								<Dialog.Root>
									<Dialog.Trigger class="text-sm rounded-md px-5 p-1 m-0.5 border bg-sky-500"
										>Create Post</Dialog.Trigger
									>

									<Dialog.Content>
										<GroupPostForm data={data.form} groupId={groupResp.data.id} />

									</Dialog.Content>
								</Dialog.Root>
							</form>
						{:else}
							<form action="?/join" method="post" class=" text-center">
								<input type="text" hidden name="target_id" value={id} />
								<button class="text-sm rounded-md px-5 border bg-sky-500 p-1 m-0.5" type="submit">
									Join group
								</button>
							</form>
						{/if}
						<!-- 
						-->
					</div>
				</div>
			</div>
		</div>

		<!-- group activity / posts -->

		<div class="h-full w-full sm:grid sm:grid-cols-2 md:grid-cols-3 gap-4 p-0 sm:p-4 mt-5 md:mt-0">
			<div class="bg-pink-500 h-56 w-full sm:rounded-lg">group</div>
			<div class="bg-purple-500 h-56 w-full sm:rounded-lg">group</div>
			<div class="bg-red-500 h-56 w-full sm:rounded-lg">123</div>
			<div class="bg-yellow-500 h-56 w-full sm:rounded-lg">123</div>
			<div class="bg-orange-500 h-56 w-full sm:rounded-lg">123</div>
		</div>
	</div>
</main>

<style>
	/* absolute */
</style>
