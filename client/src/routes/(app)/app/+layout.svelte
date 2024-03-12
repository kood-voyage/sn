<script lang="ts">
	import { ModeWatcher } from 'mode-watcher';

	///
	import Sun from 'svelte-radix/Sun.svelte';
	import Moon from 'svelte-radix/Moon.svelte';
	import Home from 'svelte-radix/Home.svelte';
	import Person from 'svelte-radix/Person.svelte';
	import Gear from 'svelte-radix/Gear.svelte';
	import Bell from 'svelte-radix/Bell.svelte';
	import Calendar from 'svelte-radix/Calendar.svelte';
	import ChatBubble from 'svelte-radix/ChatBubble.svelte';
	import Globe from 'svelte-radix/Globe.svelte';

	import * as Sheet from '$lib/components/ui/sheet/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Textarea } from '$lib/components/ui/textarea/index.js';

	///

	let previewAvatar = '';
	let previewBanner = '';

	console.log(previewAvatar);

	import { setMode, resetMode } from 'mode-watcher';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';

	function PreviewAvatar() {
		var preview = document.getElementById('previewAvatar');
		var fileInput = document.getElementById('fileInputAvatar');
		console.log();

		var file = fileInput.files[0];
		var reader = new FileReader();

		console.log(preview);

		reader.onloadend = function () {
			previewAvatar = reader.result;
		};

		if (file) {
			reader.readAsDataURL(file);
		} else {
			previewAvatar = '';
		}
	}

	function PreviewBanner() {
		var preview = document.getElementById('previewBanner');
		var fileInput = document.getElementById('fileInputBanner');
		console.log();

		var file = fileInput.files[0];
		var reader = new FileReader();

		console.log(preview);

		reader.onloadend = function () {
			previewBanner = reader.result;
		};

		if (file) {
			reader.readAsDataURL(file);
		} else {
			previewBanner = '';
		}
	}
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

				<DropdownMenu.Root>
					<DropdownMenu.Trigger asChild let:builder>
						<Button
							builders={[builder]}
							variant="ghost"
							size="icon"
							class="w-[58px] h-[58px] rounded shadow-none  transition-all duration-300  hover:bg-slate-100 hover:dark:bg-slate-900  border-none"
						>
							<Sun
								class="h-[1.2rem] w-[1.2rem]  rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
							/>

							<Moon
								class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
							/>

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

			<!-- // profile info -->
			<Sheet.Root>
				<Sheet.Trigger asChild let:builder>
					<Button builders={[builder]} variant="ghost" class="w-[60px] h-[60px]">
						<div class="flex flex-col h-1/6 items-center justify-center">
							<img
								src="https://cc-prod.scene7.com/is/image/CCProdAuthor/portrait-photography_P6a_379x392?$pjpeg$&jpegSize=100&wid=379"
								alt="avatar"
								class="w-8 h-8 rounded-full object-cover transition-all duration-300 hover:rounded-[10px]"
							/>
						</div></Button
					>
				</Sheet.Trigger>

				<Sheet.Content side="left">
					<Sheet.Header>
						<Sheet.Title>Edit profile</Sheet.Title>
						<Sheet.Description>
							Make changes to your profile here. Click save when you're done.
						</Sheet.Description>
					</Sheet.Header>
					<div class="grid gap-4 py-4">
						<!-- AVATAR -->
						<div class="m-auto">
							{#if previewAvatar}
								<img
									id="previewAvatar"
									src={previewAvatar}
									alt="previewAvatar"
									class="rounded-full m-auto w-24 h-24"
								/>
							{/if}

							<Label for="fileInputAvatar" class="text-right">Avatar Upload</Label>

							{#if !previewAvatar}
								<Input
									id="fileInputAvatar"
									type="file"
									class="col-span-3 text-red-500"
									on:change={PreviewAvatar}
								/>
							{/if}
						</div>

						<!-- BANNER -->
						<div class="m-auto w-full">
							{#if previewBanner}
								<img
									id="previewBanner"
									src={previewBanner}
									alt="previewBanner"
									class="w-full h-16 object-cover"
								/>
							{/if}

							<Label for="fileInputBanner" class="text-right">Banner Upload</Label>

							{#if !previewBanner}
								<Input
									id="fileInputBanner"
									type="file"
									class="col-span-3 text-red-500"
									on:change={PreviewBanner}
								/>
							{/if}
						</div>

						<div class="grid w-full gap-1.5">
							<Label for="description">Your bio</Label>
							<Textarea placeholder="Type your message here." id="description" class="max-h-48" />
						</div>
					</div>
					<Sheet.Footer>
						<Sheet.Close asChild let:builder>
							<Button builders={[builder]} type="submit">Save changes</Button>
						</Sheet.Close>
					</Sheet.Footer>
				</Sheet.Content>
			</Sheet.Root>
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
