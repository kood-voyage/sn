<script lang="ts">
	import { ModeWatcher } from 'mode-watcher';
	export let data;

	import {
		currentUserFollowers,
		currentUserFollowing,
		currentUserStore
	} from '$lib/store/user-store.js';
	import Sun from 'svelte-radix/Sun.svelte';
	import Moon from 'svelte-radix/Moon.svelte';
	import Home from 'svelte-radix/Home.svelte';
	import Person from 'svelte-radix/Person.svelte';
	import Bell from 'svelte-radix/Bell.svelte';
	// import Gear from 'svelte-radix/Gear.svelte';
	// import Calendar from 'svelte-radix/Calendar.svelte';
	import ChatBubble from 'svelte-radix/ChatBubble.svelte';
	import Globe from 'svelte-radix/Globe.svelte';
	import Plus from 'svelte-radix/Plus.svelte';
	import Avatar from 'svelte-radix/Avatar.svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { setMode, resetMode } from 'mode-watcher';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import NavigationItem from '../../../lib/components/navigation/navigation-item.svelte';
	// import { PUBLIC_LOCAL_PATH } from '$env/static/public';
	import { onDestroy, onMount } from 'svelte';
	import { closeWebSocket, connectWebSocket } from '$lib/client/websocket';
	import { logOut } from '$lib/client/api/user-requests';
	import SettingsForm from '$lib/components/forms/SettingsForm.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	// import { invalidateAll } from '$app/navigation';
	// import { webSocketStore } from '$lib/store/websocket-store.js';
	// console.log(data);
	import { page, navigating } from '$app/stores';

	currentUserStore.set(data.data);
	currentUserFollowers.set(data.followers);
	currentUserFollowing.set(data.following);

	onMount(async () => {
		connectWebSocket();
	});

	onDestroy(closeWebSocket);

	const handleLogout = () => {
		logOut();
		closeWebSocket();
	};

	$page.url.pathname;
</script>

<ModeWatcher />
<Toaster />

<div class="w-screen h-screen flex">
	<!-- nav Vertical-->
	<div class="w-[60px] relative shadow-md border-r bg-neutral-100 dark:bg-neutral-900">
		<div class="h-screen sm:w-[60px]">
			<div class="h-1/6"></div>

			<div class="flex flex-col h-4/6 items-center sm:w-[60px]">
				<NavigationItem href="/app" msg="Home" Icon={Home} />
				<NavigationItem href="/app/u" msg="Users" Icon={Person} />
				<NavigationItem href="/app/g" msg="Groups" Icon={Globe} />
				<NavigationItem href="/app/chat" msg="Chats" Icon={ChatBubble} />
				<NavigationItem href="/app/notification" msg="Notifications" Icon={Bell} />
				<!-- <NavigationItem href="/app/events" msg="Events" Icon={Calendar} /> -->
			</div>

			<!-- profile info end  -->

			<div class="h-1/6">
				<div>
					<SettingsForm />
				</div>
				<DropdownMenu.Root>
					<DropdownMenu.Trigger asChild let:builder>
						<Button
							builders={[builder]}
							variant="default"
							class="w-[58px] h-[58px] p-0 rounded bg-transparent shadow-none"
						>
							<div class="flex flex-col items-center justify-center h-[32px] w-[32px] p-0">
								<img
									src={$currentUserStore
										? $currentUserStore.avatar
										: 'https://thumbs.dreamstime.com/b/default-avatar-profile-vector-user-profile-default-avatar-profile-vector-user-profile-profile-179376714.jpg'}
									alt="avatar"
									class="rounded-full object-cover h-[32px] w-[32px]"
								/>
							</div></Button
						>
					</DropdownMenu.Trigger>
					<DropdownMenu.Content class="w-56">
						<DropdownMenu.Label>My Account</DropdownMenu.Label>

						<DropdownMenu.Item>
							<a href="/app/create-post" class="flex w-full text-primary">
								<span class="mr-2">
									<Plus class="h-[1rem] w-[1rem]" />
								</span>

								<span>Create post</span>
							</a>
						</DropdownMenu.Item>

						<DropdownMenu.Separator />
						<DropdownMenu.Group>
							<DropdownMenu.Item>
								<a href="/app/u/{$currentUserStore.username}" class="flex w-full">
									<span class="mr-2">
										<Avatar class="h-[1rem] w-[1rem]" />
									</span>

									<span>Profile</span>
								</a>
							</DropdownMenu.Item>
						</DropdownMenu.Group>

						<DropdownMenu.Sub>
							<DropdownMenu.SubTrigger>
								<div class="flex">
									<span class="flex mr-2 m-auto">
										<Sun class="h-[1rem] w-[1rem] block dark:hidden " />

										<Moon class="h-[1rem] w-[1rem] hidden dark:block" />
									</span>

									<span>Theme</span>
								</div>
							</DropdownMenu.SubTrigger>
							<DropdownMenu.SubContent>
								<DropdownMenu.Item on:click={() => setMode('light')}>Light</DropdownMenu.Item>
								<DropdownMenu.Item on:click={() => setMode('dark')}>Dark</DropdownMenu.Item>
								<DropdownMenu.Item on:click={() => resetMode()}>System</DropdownMenu.Item>
							</DropdownMenu.SubContent>
						</DropdownMenu.Sub>

						<DropdownMenu.Separator />

						<DropdownMenu.Group>
							<DropdownMenu.Sub>
								<DropdownMenu.SubTrigger>About</DropdownMenu.SubTrigger>
								<DropdownMenu.SubContent>
									<DropdownMenu.Item>{$currentUserStore.first_name}</DropdownMenu.Item>
									<DropdownMenu.Item>{$currentUserStore.last_name}</DropdownMenu.Item>
									<DropdownMenu.Item>{$currentUserStore.email}</DropdownMenu.Item>
								</DropdownMenu.SubContent>
							</DropdownMenu.Sub>
						</DropdownMenu.Group>

						<DropdownMenu.Item on:click={handleLogout}>Log out</DropdownMenu.Item>
					</DropdownMenu.Content>
				</DropdownMenu.Root>
			</div>
		</div>
	</div>

	<!--nav Mobile-->

	<div class="w-full">
		<slot />
	</div>

	<!-- <div class="sm:hidden">
		<div class="w-screen h-20"></div>
	</div> -->

	<div class="sm:hidden">
		<!-- Mobile navigation button -->
		<div
			class="block px-2 py-1 text-gray-700 hover:text-gray-900 focus:outline-none focus:text-gray-900 h-20"
		></div>
	</div>
</div>
