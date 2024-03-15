<script lang="ts">
	let previewAvatar = '';
	let previewBanner = '';

	let fileInputAvatar: HTMLInputElement;
	function PreviewAvatar() {
		let file = fileInputAvatar.files[0];
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

	let fileInputBanner: HTMLInputElement;
	function PreviewBanner() {
		let file = fileInputBanner.files[0];
		let reader = new FileReader();

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

<div class="w-[420px] m-auto p-8">
	<p class="text-2xl">Edit profile</p>

	<p class="text-sky-500 text-sm">
		Make changes to your profile here. Click save when you're done.
	</p>
	<form action="?/profile" method="POST" enctype="multipart/form-data" class="p-2">
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
				bind:this={fileInputAvatar}
				on:change={PreviewAvatar}
				accept="image/png, image/jpeg"
			/>
			<img
				alt="description"
				src="https://profilemediabucket-voyage.s3.amazonaws.com/profile/user1/avatar.png"
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

			<label for="fileInputBanner" class="text-right">Banner</label>

			<input
				id="fileInputBanner"
				name="fileInputBanner"
				type="file"
				class="col-span-3 text-red-500"
				bind:this={fileInputBanner}
				on:change={PreviewBanner}
				accept="image/png, image/jpeg"
			/>
		</div>

		<div class="grid w-full gap-1.5">
			<label for="description">Your bio</label>
			<textarea
				placeholder="Type your message here."
				id="description"
				name="description"
				class="max-h-48 min-h-20"
				maxlength="400"
			/>
		</div>

		<input type="submit" value="save" class="w-20 bg-slate-600 p-2 rounded-xl" />
	</form>
</div>
