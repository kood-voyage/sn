<script lang="ts">
	import * as Sheet from '$lib/components/ui/sheet/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Textarea } from '$lib/components/ui/textarea/index.js';
	let previewAvatar = '';
	let previewBanner = '';

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

<Sheet.Root>
	<Sheet.Trigger asChild let:builder>
		<Button builders={[builder]} variant="ghost">Settings</Button>
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

				<Input
					id="fileInputAvatar"
					type="file"
					class="col-span-3 text-red-500"
					on:change={PreviewAvatar}
				/>
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

				<Input
					id="fileInputBanner"
					type="file"
					class="col-span-3 text-red-500"
					on:change={PreviewBanner}
				/>
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
