<script lang="ts">
	import { ModeWatcher } from 'mode-watcher';
	import Sheet from '$lib/components/Sheet.svelte';

	export let data;

	const { username, email, first_name, last_name } = data;

	///
	import Sun from 'svelte-radix/Sun.svelte';
	import Moon from 'svelte-radix/Moon.svelte';
	import Home from 'svelte-radix/Home.svelte';
	import Person from 'svelte-radix/Person.svelte';
	import Bell from 'svelte-radix/Bell.svelte';
	import Gear from 'svelte-radix/Gear.svelte';
	import Calendar from 'svelte-radix/Calendar.svelte';
	import ChatBubble from 'svelte-radix/ChatBubble.svelte';
	import Globe from 'svelte-radix/Globe.svelte';

	import { Button } from '$lib/components/ui/button/index.js';

	import { setMode, resetMode } from 'mode-watcher';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';

	///
</script>

<ModeWatcher />

<div class=" h-screen sm:h-auto sm:flex">
	<!-- nav Vertical-->
	<div class="hidden sm:block sm:w-[60px] relative shadow-md border-r">
		<div class="h-screen sm:w-[60px]">
			<div class="h-1/6"></div>

			<div class="flex flex-col h-4/6 items-center sm:w-[60px]">
				<a href="/app">
					<div class="button">
						<Home class="h-[1.2rem] w-[1.2rem] m-auto self-center" />
					</div>
				</a>

				<a href="/app/u">
					<div class="button">
						<Person class="h-[1.2rem] w-[1.2rem] m-auto self-center" />
					</div>
				</a>

				<a href="/app/g">
					<div class="button">
						<Globe class="h-[1.2rem] w-[1.2rem] m-auto self-center" />
					</div>
				</a>

				<a href="/app/chat">
					<div class="button">
						<ChatBubble class="h-[1.2rem] w-[1.2rem] m-auto self-center" />
					</div>
				</a>

				<a href="/app">
					<div class="button">
						<Bell class="h-[1.2rem] w-[1.2rem] m-auto self-center" />
					</div>
				</a>

				<a href="/app">
					<div class="button">
						<Calendar class="h-[1.2rem] w-[1.2rem] m-auto self-center" />
					</div>
				</a>

				<!-- <a href="/app">
					<div class="button">
						<Gear class="h-[1.2rem] w-[1.2rem] m-auto self-center" />
					</div>
				</a> -->
			</div>

			<!-- // profile info -->

			<!-- profile info end  -->

			<div class="h-1/6">
				<DropdownMenu.Root>
					<DropdownMenu.Trigger asChild let:builder>
						<Button builders={[builder]} variant="ghost" class="w-[58px] h-[58px] p-0">
							<div class="flex flex-col items-center justify-center h-[32px] w-[32px] p-0">
								<img
									src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed={username}"
									alt="avatar"
									class="rounded-full object-cover hover:rounded-[10px]"
								/>
							</div></Button
						>
					</DropdownMenu.Trigger>
					<DropdownMenu.Content class="w-56">
						<DropdownMenu.Label>My Account</DropdownMenu.Label>
						<DropdownMenu.Item class="text-blue-500">{username}</DropdownMenu.Item>

						<DropdownMenu.Separator />
						<DropdownMenu.Group>
							<DropdownMenu.Item>
								<a href="/app/u/{username}" class="w-full"> Profile</a>
							</DropdownMenu.Item>

							<DropdownMenu.Item>
								<a href="/app/settings" class="flex">
									<span class="mr-1 m-auto">
										<Gear class="h-[1rem] w-[1rem]" />
									</span>

									<span>Settings</span>
								</a>
							</DropdownMenu.Item>
						</DropdownMenu.Group>

						<DropdownMenu.Sub>
							<DropdownMenu.SubTrigger>
								<div class="flex">
									<span class="flex mr-1 m-auto">
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

						<DropdownMenu.Separator />
						<DropdownMenu.Item>
							Log out
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

<style>
	.button {
		@apply flex h-[58px] w-[58px] cursor-pointer rounded transition-all duration-300  hover:bg-slate-100 hover:dark:bg-slate-900;
	}
</style>
