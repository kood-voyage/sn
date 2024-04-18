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
	import Gear from 'svelte-radix/Gear.svelte';
	import Calendar from 'svelte-radix/Calendar.svelte';
	import ChatBubble from 'svelte-radix/ChatBubble.svelte';
	import Globe from 'svelte-radix/Globe.svelte';
	import Plus from 'svelte-radix/Plus.svelte';
	import Avatar from 'svelte-radix/Avatar.svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { setMode, resetMode } from 'mode-watcher';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import NavigationItem from './navigation-item.svelte';

	currentUserStore.set(data.data);
	currentUserFollowers.set(data.followers);
	currentUserFollowing.set(data.following);

	const { username, email, first_name, last_name, avatar } = $currentUserStore;
</script>

<ModeWatcher />

<div class="w-screen h-screen sm:h-auto sm:flex">
	<!-- nav Vertical-->
	<div class="hidden sm:block sm:w-[60px] relative shadow-md border-r">
		<div class="h-screen sm:w-[60px]">
			<div class="h-1/6"></div>

			<div class="flex flex-col h-4/6 items-center sm:w-[60px]">
				<NavigationItem href="/app" msg="Home" Icon={Home} />
				<NavigationItem href="/app/u" msg="Users" Icon={Person} />
				<NavigationItem href="/app/g" msg="Groups" Icon={Globe} />
				<NavigationItem href="/app/chat" msg="Chats" Icon={ChatBubble} />
				<NavigationItem href="/app/notification" msg="Notifications" Icon={Bell} />
				<NavigationItem href="/app/events" msg="Events" Icon={Calendar} />
			</div>

			<!-- // profile info -->

			<!-- profile info end  -->

			<div class="h-1/6">
				<DropdownMenu.Root>
					<DropdownMenu.Trigger asChild let:builder>
						<Button builders={[builder]} variant="ghost" class="w-[58px] h-[58px] p-0">
							<div class="flex flex-col items-center justify-center h-[32px] w-[32px] p-0">
								<img
									src={avatar}
									alt="avatar"
									class="rounded-full object-cover hover:rounded-[10px] h-[32px] w-[32px]"
								/>
							</div></Button
						>
					</DropdownMenu.Trigger>
					<DropdownMenu.Content class="w-56">
						<DropdownMenu.Label>My Account</DropdownMenu.Label>

						<DropdownMenu.Item >
							<a href="/app/create-post" class="flex w-full text-primary">
								<span class="mr-2">
									<Plus class="h-[1rem] w-[1rem]" />
								</span>

								<span >Create post</span>
							</a>
						</DropdownMenu.Item>

						<DropdownMenu.Separator />
						<DropdownMenu.Group>
							<DropdownMenu.Item>
								<a href="/app/u/{username}" class="flex w-full">
									<span class="mr-2">
										<Avatar class="h-[1rem] w-[1rem]" />
									</span>

									<span>Profile</span>
								</a>
							</DropdownMenu.Item>

							<DropdownMenu.Item>
								<a href="/app/settings" class="flex w-full">
									<span class="mr-2">
										<Gear class="h-[1rem] w-[1rem]" />
									</span>

									<span>Settings</span>
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
									<DropdownMenu.Item>{first_name}</DropdownMenu.Item>
									<DropdownMenu.Item>{last_name}</DropdownMenu.Item>
									<DropdownMenu.Item>{email}</DropdownMenu.Item>
								</DropdownMenu.SubContent>
							</DropdownMenu.Sub>
						</DropdownMenu.Group>

						<DropdownMenu.Item>
							<a href="/logout"> Log out </a>

							<!-- <DropdownMenu.Shortcut>⇧⌘Q</DropdownMenu.Shortcut> -->
						</DropdownMenu.Item>
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
		<div class="w-screen h-20 "></div>
	</div> -->

	<div class="sm:hidden">
		<!-- Mobile navigation button -->
		<div
			class="block px-2 py-1 text-gray-700 hover:text-gray-900 focus:outline-none focus:text-gray-900 h-20"
		></div>
	</div>
</div>
