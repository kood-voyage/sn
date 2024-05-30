<script lang="ts">
	import { ModeWatcher } from 'mode-watcher';
	import Sun from 'svelte-radix/Sun.svelte';
	import Moon from 'svelte-radix/Moon.svelte';
	import { setMode, resetMode } from 'mode-watcher';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
</script>

<div
	class="w-screen h-screen flex flex-col bg-gradient-to-r"
>
	<ModeWatcher />
	<nav class="absolute bottom-8 right-8 rounded-full">
		<div class="flex h-full items-center justify-end rounded-full">
			<span class="font-semibold"></span>

			<div class="flex flex-col rounded-full">
				<DropdownMenu.Root>
					<DropdownMenu.Trigger asChild let:builder>
						<Button
							builders={[builder]}
							variant="outline"
							class="rounded-full w-12 h-12"
							size="icon"
						>
							<Sun
								class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
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
		</div>
	</nav>

	<div class="w-full sm:w-2/3 md:w-2/4 lg:w-1/3 m-auto items-center px-10 py-14 rounded-xl ">
		<slot />
	</div>
</div>
