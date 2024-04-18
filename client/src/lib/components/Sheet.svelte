<script lang="ts">
	let previewAvatar = '';
	let previewBanner = '';

	function PreviewAvatar() {
		const fileInput = document.getElementById('fileInputAvatar');

		const file = fileInput.files[0];
		const reader = new FileReader();

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
		const fileInput = document.getElementById('fileInputBanner');

		const file = fileInput.files[0];
		const reader = new FileReader();

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

<p>Edit profile</p>

<p>Make changes to your profile here. Click save when you're done.</p>

<form action="?/profile" >
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

			<label for="fileInputAvatar" class="text-right">Avatar Upload</label>

			<input
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

			<label for="fileInputBanner" class="text-right">Banner Upload</label>

			<input
				id="fileInputBanner"
				type="file"
				class="col-span-3 text-red-500"
				on:change={PreviewBanner}
			/>
		</div>

		<div class="grid w-full gap-1.5">
			<label for="description">Your bio</label>
			<textarea placeholder="Type your message here." id="description" class="max-h-48" />
		</div>

		<button type="submit">submit</button>

</form>
