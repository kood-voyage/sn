<script lang="ts">
	import { handleSubmit } from './page.js';
	import { handleImageCopression } from '$lib/client/image-compression.js';

	export let data;

	const { username, email, first_name, last_name, avatar, cover, description } = data;

	$: previewAvatar = avatar;
	$: previewCover = cover;

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

<div class="w-[420px] m-auto p-8">
	<p class="text-2xl">Edit profile</p>

	<p class="text-sky-500 text-sm">
		Make changes to your profile here. Click save when you're done.
	</p>
	<form id="imageForm" enctype="multipart/form-data" class="p-2">
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

		<input
			type="button"
			value="save"
			on:click={handleSubmit}
			class="w-20 bg-neutral-600 p-2 rounded-xl"
		/>
	</form>
</div>
