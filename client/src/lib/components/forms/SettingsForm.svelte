<script lang="ts">
	import * as Sheet from '$lib/components/ui/sheet/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Gear } from 'svelte-radix';
</script>

<Sheet.Root>
	<Sheet.Trigger>
		<div class="flex items-center">
			<Gear class="h-[1rem] w-[1rem] mr-1" />
			<p>Edit</p>
		</div>
	</Sheet.Trigger>

	<Sheet.Content side="right">
		<Sheet.Header>
			<Sheet.Title>Settings</Sheet.Title>
			<Sheet.Description>
				Make changes to your profile here. Click save when you're done.
			</Sheet.Description>
		</Sheet.Header>

		<div class="grid gap-4 py-4">
			<div class="grid grid-cols-4 items-center gap-4">
				<Label for="name" class="text-right">Name</Label>
				<Input id="name" value="Pedro Duarte" class="col-span-3" />
			</div>
			<div class="grid grid-cols-4 items-center gap-4">
				<Label for="username" class="text-right">Username</Label>
				<Input id="username" value="@peduarte" class="col-span-3" />
			</div>
		</div>

		<Sheet.Footer>
			<Sheet.Close asChild let:builder>
				<Button builders={[builder]} type="submit">Save changes</Button>
			</Sheet.Close>
		</Sheet.Footer>
	</Sheet.Content>
</Sheet.Root>