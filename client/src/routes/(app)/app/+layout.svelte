<script lang="ts">
	import * as Resizable from '$lib/components/ui/resizable';
	import { ModeWatcher } from 'mode-watcher';
	import Sun from 'svelte-radix/Sun.svelte';
	import Moon from 'svelte-radix/Moon.svelte';
	import Home from 'svelte-radix/Home.svelte';

	import Person from 'svelte-radix/Person.svelte';

	import Gear from 'svelte-radix/Gear.svelte';
	import Bell from 'svelte-radix/Bell.svelte';

	import Calendar from 'svelte-radix/Calendar.svelte';

	import ChatBubble from 'svelte-radix/ChatBubble.svelte';

	import { setMode, resetMode } from 'mode-watcher';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';

	export let defaultCollapsed = false;

	// import Icon from '@iconify/svelte';

	let isCollapsed = defaultCollapsed;

	function onLayoutChange(sizes: number[]) {
		document.cookie = `PaneForge:layout=${JSON.stringify(sizes)}`;
	}

	function onCollapse() {
		isCollapsed = true;
		document.cookie = `PaneForge:collapsed=${true}`;
	}

	function onExpand() {
		isCollapsed = false;
		document.cookie = `PaneForge:collapsed=${false}`;
	}
</script>

<div class="h-screen">
	<ModeWatcher />

	<Resizable.PaneGroup
		direction="horizontal"
		{onLayoutChange}
		class="w-screen rounded-lg border items-stretch"
		defaultSize={8}
		minSize={4}
		maxSize={8}
	>
		<Resizable.Pane class="bg-slate-100 dark:bg-slate-900">
			<div class="flex flex-col h-full justify-between">
				<div class="h-1/3"></div>

				<div class="flex flex-col h-4/6 items-center">
					<a href="/app">
						<div class="button">
							<Home class="h-[1.8rem] w-[1.8rem] m-auto self-center" />
						</div>
					</a>

					<a href="/app/u/person">
						<div class="button">
							<Person class="h-[1.8rem] w-[1.8rem] m-auto self-center" />
						</div>
					</a>
					
					<a href="/app/chat">
						<div class="button">
							<ChatBubble class="h-[1.8rem] w-[1.8rem] m-auto self-center" />
						</div>
					</a>

					<div class="button">
						<Bell class="h-[1.8rem] w-[1.8rem] m-auto self-center" />
					</div>

					<div class="button">
						<Calendar class="h-[1.8rem] w-[1.8rem] m-auto self-center" />
					</div>

					<div class="button">
						<Gear class="h-[1.8rem] w-[1.8rem] m-auto self-center" />
					</div>
					<DropdownMenu.Root>
						<DropdownMenu.Trigger asChild let:builder>
							<Button
								builders={[builder]}
								variant="outline"
								size="icon"
								class="w-[60px] h-[60px] rounded shadow-none bg-slate-100 dark:bg-slate-900  transition-all duration-300  hover:bg-slate-300 hover:dark:bg-slate-800  border-none"
							>
								<Sun
									class="h-[1.8rem] w-[1.8rem]  rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
								/>
								<!-- 
								<Icon
									icon="meteocons:clear-day-fill"
									width="32"
									height="32"
									class=" rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
								/> -->
								<Moon
									class="absolute h-[1.8rem] w-[1.8rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
								/>
								<!-- <Icon
									icon="meteocons:clear-night"
									width="32"
									height="32"
									class="absolute  rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
								/> -->
								<span class="sr-only">Toggle theme</span>
							</Button>
						</DropdownMenu.Trigger>
						<DropdownMenu.Content align="end">
							<DropdownMenu.Item on:click={() => setMode('light')}>Light</DropdownMenu.Item>
							<DropdownMenu.Item on:click={() => setMode('dark')}>Dark</DropdownMenu.Item>
							<DropdownMenu.Item on:click={() => resetMode()}>System</DropdownMenu.Item>
						</DropdownMenu.Content>
					</DropdownMenu.Root>
				</div>

				<div class="flex flex-col h-1/3 items-center">
					<img
						src="https://cc-prod.scene7.com/is/image/CCProdAuthor/portrait-photography_P6a_379x392?$pjpeg$&jpegSize=100&wid=379"
						alt="avatar"
						class="w-[48px] h-[48px] rounded-full object-cover transition-all duration-300 hover:rounded-[10px]"
					/>
				</div>
			</div>
		</Resizable.Pane>

		<Resizable.Handle class="hover:bg-gradient-to-b from-cyan-500 to-blue-500 hover:w-[2px] " />

		<Resizable.Pane defaultSize={92} minSize={90} maxSize={96}>
			<div class="flex w-full h-full">
				<slot />
			</div>
		</Resizable.Pane>
	</Resizable.PaneGroup>
</div>

<style>
	.button {
		@apply flex h-[60px] w-[60px] cursor-pointer rounded transition-all duration-300  hover:bg-slate-300 hover:dark:bg-slate-800;
	}

	.icon {
		@apply h-[2rem] w-[2rem];
	}
</style>
