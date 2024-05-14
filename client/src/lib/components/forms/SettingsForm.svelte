<script lang="ts">
	import * as Sheet from '$lib/components/ui/sheet/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Gear } from 'svelte-radix';
	import * as Tooltip from '../ui/tooltip';

	import { handleImageCopression } from '$lib/client/image-compression.js';
	import { currentUserStore } from '$lib/store/user-store';

	$: previewAvatar = $currentUserStore.avatar;
	$: previewCover = $currentUserStore.cover;
	$: description = $currentUserStore.description;

	let fileInputAvatar: HTMLInputElement;
	async function PreviewAvatar() {
		let file = fileInputAvatar.files[0];

		const imageResp = await handleImageCopression(file);
		if (imageResp.ok && imageResp.file) {
			file = imageResp.file;
		}
		let reader = new FileReader();
		reader.onloadend = function () {
			previewAvatar = reader.result;
		};
		if (file) {
			reader.readAsDataURL(file);
		} else {
			previewAvatar = '';
		}
	}

	let fileInputCover: HTMLInputElement;
	async function PreviewCover() {
		let file = fileInputCover.files[0];
		const imageResp = await handleImageCopression(file);
		if (imageResp.ok && imageResp.file) {
			file = imageResp.file;
		}
		let reader = new FileReader();
		reader.onloadend = function () {
			previewCover = reader.result;
		};
		if (file) {
			reader.readAsDataURL(file);
		} else {
			previewCover = '';
		}
	}



</script>

<Sheet.Root>
	<Sheet.Trigger>
		<Tooltip.Root>
			<Tooltip.Trigger>
				<div class="button">
					<Gear class="h-[1.4rem] w-[1.4rem] m-auto" />
				</div>
			</Tooltip.Trigger>
			<Tooltip.Content side="right">
				<p>Edit Profile</p>
			</Tooltip.Content>
		</Tooltip.Root>
	</Sheet.Trigger>

	<Sheet.Content side="right">
		<Sheet.Header>
			<Sheet.Title>Settings</Sheet.Title>
			<Sheet.Description class="text-sky-500 text-sm">
				Make changes to your profile here. Click save when you're done.
			</Sheet.Description>
		</Sheet.Header>

		<div class=" m-auto">
			<form id="settingForm" enctype="multipart/form-data" >
				<!-- AVATAR -->
				<div class="m-auto">
					{#if previewAvatar}
						<img
							id="previewAvatar"
							src={previewAvatar}
							alt="previewAvatar"
							class="rounded-full m-auto w-24 h-24 object-cover"
						/>
					{/if}

					<label for="fileInputAvatar" class="text-right">Avatar</label>

					<input
						src={previewAvatar}
						id="fileInputAvatar"
						name="fileInputAvatar"
						type="file"
						class="col-span-3 text-red-500"
						bind:this={fileInputAvatar}
						on:change={PreviewAvatar}
						accept="image/png, image/jpeg"
					/>
				</div>

				<!-- BANNER -->
				<div class="m-auto w-full">
					{#if previewCover}
						<img
							id="previewCover"
							src={previewCover}
							alt="previewCover"
							class="w-full h-16 object-cover"
						/>
					{/if}

					<label for="fileInputCover" class="text-right">Banner</label>

					<input
						src={previewCover}
						id="fileInputCover"
						name="fileInputCover"
						type="file"
						class="col-span-3 text-red-500"
						bind:this={fileInputCover}
						on:change={PreviewCover}
						accept="image/png, image/jpeg"
					/>
				</div>

				<div class="grid w-full gap-1.5">
					<label for="description">Your bio</label>
					<textarea
						placeholder={description}
						value={description}
						id="description"
						name="description"
						class="max-h-48 min-h-20"
						maxlength="400"
					/>
				</div>

				<Sheet.Footer>
					<Sheet.Close asChild let:builder>
						<Button builders={[builder]} type="submit">Save changes</Button>
					</Sheet.Close>
				</Sheet.Footer>
			</form>
		</div>
	</Sheet.Content>
</Sheet.Root>

<style>
	.button {
		@apply flex h-[58px] w-[58px] cursor-pointer rounded transition-all duration-300  hover:bg-primary hover:dark:bg-primary;
	}
</style>
