<script lang="ts">
	import * as Sheet from '$lib/components/ui/sheet/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Gear } from 'svelte-radix';
	import * as Tooltip from '../ui/tooltip';

	import { handleImageCopression } from '$lib/client/image-compression.js';
	import { currentUserStore } from '$lib/store/user-store';
	import { imageStore } from '$lib/client/api/image-requests';
	import { updateDescription } from '$lib/client/api/user-requests';

	$: previewAvatar = $currentUserStore.avatar;
	$: previewCover = $currentUserStore.cover;

	let description = $currentUserStore.description;

	let avatar: File;
	let cover: File;

	async function handleImageChange(file: File, type: string) {
		if (file === undefined) {
			return;
		}
		if (file.length === 1) {
			const fileResp = await handleImageCopression(file[0]);
			let reader = new FileReader();
			reader.onloadend = function () {
				if (type === 'avatar') {
					previewAvatar = reader.result;
				}

				if (type === 'cover') {
					previewCover = reader.result;
				}
			};
			reader.readAsDataURL(fileResp.file as File);

			const fileFormData = new FormData();
			fileFormData.append('path', `profile/${$currentUserStore.id}`);
			fileFormData.append('images', fileResp.file as File);
			const resp = await imageStore(fileFormData, $currentUserStore.id, type);

			if ((resp.data = 'Successfully uploaded to S3')) {
				if (type === 'avatar') {
					$currentUserStore.avatar = previewAvatar;
				}

				if (type === 'cover') {
					$currentUserStore.cover = previewCover;
				}
			}
		}
	}

	$: handleImageChange(avatar, 'avatar');
	$: handleImageChange(cover, 'cover');
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

	<Sheet.Content side="right" class="overflow-y-scroll h-screen">
		<Sheet.Header>
			<Sheet.Title>Settings</Sheet.Title>
			<Sheet.Description class="text-sky-500 text-sm">
				Make changes to your profile here. Click save when you're done.
			</Sheet.Description>
		</Sheet.Header>

		<div class=" m-auto">
			<form id="settingForm" enctype="multipart/form-data">
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
						id="fileInputAvatar"
						name="fileInputAvatar"
						type="file"
						class="col-span-3 text-red-500"
						bind:files={avatar}
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
						id="fileInputCover"
						name="fileInputCover"
						type="file"
						class="col-span-3 text-red-500"
						bind:files={cover}
						accept="image/png, image/jpeg"
					/>
				</div>

				<div class="grid w-full gap-1.5">
					<label for="description">Your bio</label>
					<textarea
						placeholder={description}
						bind:value={description}
						id="description"
						name="description"
						class="max-h-48 min-h-20"
						maxlength="400"
						minlength="4"
					/>
				</div>

				<Sheet.Footer>
					<Sheet.Close asChild let:builder>
						<Button
							builders={[builder]}
							type="submit "
							class="w-full mt-4"
							on:click={() => updateDescription(description)}>Save changes</Button
						>
					</Sheet.Close>
				</Sheet.Footer>
				<img src={'../pretty.png'} alt="login" class="p-20" />
			</form>
		</div>
	</Sheet.Content>
</Sheet.Root>

<style>
	.button {
		@apply hover:bg-primary hover:dark:bg-primary flex h-[58px] w-[58px] cursor-pointer rounded  transition-all duration-300;
	}
</style>
