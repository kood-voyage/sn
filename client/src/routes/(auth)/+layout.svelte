<script lang="ts">
	import { ModeWatcher } from 'mode-watcher';
	import Sun from 'svelte-radix/Sun.svelte';
	import Moon from 'svelte-radix/Moon.svelte';

	import Home from 'svelte-radix/Home.svelte';

	import { setMode, resetMode } from 'mode-watcher';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import Icon from '@iconify/svelte';
</script>



<div class="w-screen h-screen flex flex-col">
	<ModeWatcher />
	<nav class="h-12 w-full">
		<div class="flex h-full items-center justify-end mr-2">
			<span class="font-semibold"></span>

			<div class="flex flex-col">


				<DropdownMenu.Root>
					<DropdownMenu.Trigger asChild let:builder>
						<Button builders={[builder]} variant="outline" size="icon">
							<Sun
								class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
							/>

							<!-- <Icon
								icon="line-md:sunny-outline-twotone-loop"
								width="32"
								height="32"
								class=" rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
							/> -->
							<Moon
								class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
							/>
							<!-- <Icon
								icon="line-md:moon-twotone-alt-loop"
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
		</div>
	</nav>
	<div class="w-3/12 m-auto items-center border px-10 py-14 rounded-xl">
		<slot />
	</div>
</div>
