<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;

	const { name } = data;

	// type Data = {
	// 	post: number;
	// 	followers: number;
	// 	following: number;
	// };
</script>

<svelte:head>
	<title>g/{name}</title>
</svelte:head>

<!-- user profile page -->

<main class="flex">
	<div class=" m-auto h-full w-full max-w-[1096px]">
		<!-- profile info header -->
		<div class="profile-info relative">
			<!-- banner img  -->
			<div class="m-auto h-32 sm:h-60 max-w-[1096px] p-0 sm:px-2">
				<img
					class="h-full w-full object-cover sm:rounded-xl"
					src="https://profilemediabucket-voyage.s3.amazonaws.com/group/e9957939-7d54-40b1-bfef-081952574981/coverb"
					alt="banner"
				/>
			</div>

			<div class="max-w-[1096px] sm:px-2 h-16">
				<div class="w-full bg-slate-200/30 p-1 mt-1 h-full flex items-center sm:rounded-xl">
					<div class="h-full align-middle self-center">
						<p class="h-full w-fit align-middle self-center">{name}</p>
					</div>
				</div>
			</div>

			<!-- <div class="w-full relative">
				<img
					src="https://api.contentstack.io/v2/assets/575e4d1c0342dfd738264a1f/download?uid=bltada7771f270d08f6"
					alt="banner"
					class="m-auto h-32 w-32 relative bottom-20 rounded-full border-4 border-white object-cover dark:border-slate-950"
				/>
				<h1 class="sm:bottom-4 left-[140px] text-2xl font-bold">
					{name}
				</h1>
			</div> -->
		</div>

		<!-- group activity / posts -->

		<div class="h-full w-full sm:grid sm:grid-cols-2 md:grid-cols-3 gap-4 p-0 sm:p-4 mt-5 md:mt-0">
			<div class="bg-pink-500 h-56 w-full sm:rounded-lg">group</div>
			<div class="bg-purple-500 h-56 w-full sm:rounded-lg">group</div>
			<div class="bg-red-500 h-56 w-full sm:rounded-lg">123</div>
			<div class="bg-yellow-500 h-56 w-full sm:rounded-lg">123</div>
			<div class="bg-orange-500 h-56 w-full sm:rounded-lg">123</div>
		</div>
	</div>
</main>

<style>
	/* absolute */
</style>
