<script lang="ts">
	import Post from '$lib/components/Post.svelte';
	import type { PageData } from './$types';

	export let data: PageData;

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

	const demoData = [
		{
			imgURL: 'https://c.files.bbci.co.uk/E909/production/_112375695_crucible976.jpg',
			title: 'The Evolution of Gaming Consoles',
			body: 'The evolution of gaming consoles has been a fascinating journey, marked by technological advancements, innovative design, and fierce competition among industry giants. From the humble beginnings of the Magnavox Odyssey to the groundbreaking release of the PlayStation 5 and Xbox Series X, gaming consoles have continuously pushed the boundaries of gaming experiences. This article explores the key milestones, breakthroughs, and cultural impact of gaming consoles throughout history.'
		},
		{
			imgURL:
				'https://assetsio.reedpopcdn.com/best-fps-header.jpg?width=1200&height=1200&fit=bounds&quality=70&format=jpg&auto=webp',
			title: 'Top 10 Must-Play Indie Games of the Year',
			body: 'Indie games continue to captivate players with their unique art styles, compelling narratives, and innovative gameplay mechanics. In this article, we showcase the top 10 must-play indie games of the year, spanning various genres and platforms. From emotionally charged adventures to mind-bending puzzles, these indie gems offer unforgettable experiences that rival even the biggest AAA titles. Get ready to discover your next gaming obsession!'
		},
		{
			imgURL: 'https://www.premiumbeat.com/blog/wp-content/uploads/2022/08/Best-VR-Film.jpg',
			title: 'Exploring the World of Virtual Reality Gaming',
			body: 'Virtual reality gaming has emerged as one of the most immersive and exhilarating forms of entertainment, transporting players to fantastical worlds and pushing the boundaries of reality. In this article, we delve into the exciting world of virtual reality gaming, exploring the latest hardware advancements, groundbreaking titles, and the transformative impact of VR technology on the gaming industry. Strap on your headset and get ready to experience gaming like never before!'
		},
		{
			imgURL: 'https://dossiers.dhnet.be/esport/media/08583030-photo-esport-mr.jpg',
			title: 'The Rise of Esports: Competitive Gaming Takes the World by Storm',
			body: "Esports has evolved from niche hobby to global phenomenon, captivating millions of viewers and offering lucrative opportunities for skilled players around the world. In this article, we examine the meteoric rise of esports, tracing its humble beginnings in local arcades to sold-out arenas and multi-million dollar tournaments. From the strategic depth of MOBAs to the lightning-fast reflexes of first-person shooters, esports encompasses a diverse array of competitive gaming experiences that continue to push the boundaries of what's possible in gaming."
		}
	];
</script>

<svelte:head>
	<title>u/{data.username}</title>
</svelte:head>

<!-- user profile page -->

<main class=" overflow-scroll ">
	<div class=" m-auto h-full w-full max-w-[1096px]">
		<!-- profile info header -->
		<div class="profile-info relative">
			<!-- banner img  -->
			<div class="m-auto h-24 sm:h-40 max-w-[1096px] p-0 sm:p-4">
				<img
					class="h-full w-full rounded-none sm:rounded-xl object-cover"
					src={data.cover}
					alt="cover"
				/>
			</div>

			<div class="h-8 relative mx-0 sm:mx-4">
				<img
					src={data.avatar}
					alt="avatar"
					class="absolute bottom-[1px] left-12 h-20 w-20 rounded-full border-4 border-white object-cover dark:border-slate-950"
				/>

				<div
					class="absolute bottom-1 sm:bottom-3 left-[140px] bg-white dark:bg-slate-950 rounded-2xl flex"
				>
					<p class=" md:text-2xl font-bold mr-2">{data.username}</p>
					<!-- <p class="text-xs">
						<span class="bg-slate-500 px-1 rounded-sm">AKA</span>
						{data.first_name}
						{data.last_name}
					</p> -->
					<button class="text-sm px-5 rounded-md bg-sky-500"> follow </button>
				</div>

				<div class="hidden md:block">
					<div class="flex dark:bg-slate-900 w-72 rounded-sm absolute right-1 bottom-4">
						<div
							class="text-xs w-1/3 border-r text-center hover:bg-slate-200 dark:hover:bg-slate-800 p-1"
						>
							<span class="font-bold">{demoData.length}</span> posts
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
			{#each demoData as data}
				<Post {...data} />
			{/each}
		</div>
	</div>
</main>

<style>
</style>
