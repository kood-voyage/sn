<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;

	const { username, id, email, first_name, last_name, date_of_birth } = data;

	type Data = {
		post: number;
		followers: number;
		following: number;
	};

	const profileData: Data = {
		post: 0,
		followers: 1,
		following: 0
	};
</script>

<svelte:head>
	<title>u/{username}</title>
</svelte:head>

<!-- user profile page -->

<main class="flex overflow-scroll">
	<div class=" m-auto h-full w-full max-w-[1096px]">
		<!-- profile info header -->
		<div class="profile-info relative">
			<!-- banner img  -->
			<div class="m-auto h-24 sm:h-40 max-w-[1096px] p-0 sm:p-4">
				<img
					class="h-full w-full rounded-none sm:rounded-xl object-cover"
					src="https://letsenhance.io/static/8f5e523ee6b2479e26ecc91b9c25261e/1015f/MainAfter.jpg"
					alt="banner"
				/>
			</div>

			<div class="h-8 relative mx-0 sm:mx-4">
				<img
					src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed={username}"
					alt="banner"
					class="absolute bottom-[1px] left-12 h-20 w-20 rounded-full border-4 border-white object-cover dark:border-slate-950"
				/>

				<div
					class="absolute bottom-1 sm:bottom-3 left-[140px] bg-white dark:bg-slate-950 rounded-2xl flex"
				>
					<p class=" md:text-2xl font-bold mr-2">{username}</p>
					<!-- <p class="text-xs">
						<span class="bg-slate-500 px-1 rounded-sm">AKA</span>
						{first_name}
						{last_name}
					</p> -->
					<button class="text-sm px-5 rounded-md bg-sky-500"> follow </button>
				</div>

				<div class="hidden md:block">
					<div class="flex dark:bg-slate-900 w-72 rounded-sm absolute right-1 bottom-4">
						<div
							class="text-xs w-1/3 border-r text-center hover:bg-slate-200 dark:hover:bg-slate-800 p-1"
						>
							<span class="font-bold">{profileData.post}</span> posts
						</div>
						<div
							class="text-xs w-1/3 border-r text-center hover:bg-slate-200 dark:hover:bg-slate-800 p-1"
						>
							<span class="font-bold">{profileData.followers}</span> followers
						</div>
						<div class="text-xs w-1/3 text-center hover:bg-slate-200 dark:hover:bg-slate-800 p-1">
							<span class="font-bold">{profileData.following}</span> following
						</div>
					</div>
				</div>
			</div>

			<div class="sm:mx-4 h-8 block md:hidden">
				<div class="">
					<div class="flex bg-slate-100 dark:bg-slate-900 w-full rounded-lg">
						<div
							class="text-xs w-1/3 border-r text-center p-4 hover:bg-slate-200 dark:hover:bg-slate-800 hover:rounded-l-lg"
						>
							<span class="font-bold">{profileData.post}</span> posts
						</div>
						<div
							class="text-xs w-1/3 border-r text-center p-4 hover:bg-slate-200 dark:hover:bg-slate-800"
						>
							<span class="font-bold">{profileData.followers}</span> followers
						</div>
						<div
							class="text-xs w-1/3 text-center p-4 hover:rounded-r-lg hover:bg-slate-200 dark:hover:bg-slate-800"
						>
							<span class="font-bold">{profileData.following}</span> following
						</div>
					</div>
				</div>
			</div>
		</div>

		<!-- profile activity / posts -->

		<div class="h-full w-full sm:grid sm:grid-cols-2 md:grid-cols-3 gap-1 p-0 sm:p-4 mt-5 md:mt-0">
			<div class="bg-pink-500 h-96 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-purple-500 h-96 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-red-500 h-96 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-yellow-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-orange-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-pink-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-purple-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-red-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-yellow-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-orange-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-pink-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-purple-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-red-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-yellow-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
			<div class="bg-orange-500 h-56 w-full sm:rounded-lg mb-1 sm:mb-0">123</div>
		</div>
	</div>
</main>

<style>
</style>
