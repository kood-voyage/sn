<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;

	const { username } = data;
</script>

<svelte:head>
	<title>u/{username}</title>
</svelte:head>

<!-- user profile page -->

<main class="flex w-full">
	<div class="w-screen h-screen">
		<div class="profile">
			<!-- profile info header -->
			<div class="profile-info relative">
				<!-- banner img  -->
				<div class="banner">
					<img
						class="image"
						src="https://letsenhance.io/static/8f5e523ee6b2479e26ecc91b9c25261e/1015f/MainAfter.jpg"
						alt="banner"
					/>
				</div>

				<div class="h-8 relative">
					<div class="avatar">
						<img
							src="https://cdn.create.vista.com/api/media/small/490310088/stock-photo-green-basilisk-or-jesus-christ-lizard-in-the-zoo-terrarium"
							alt="banner"
						/>
					</div>

					<h1 class="">{username}</h1>

					<!-- <button>
                        follow
                    </button> -->
				</div>
			</div>

			<!-- profile activity / posts -->

			<div></div>
		</div>
	</div>

	<!-- <aside>fsdf</aside> -->

	<span class="icon-[meteocons--clear-day-fill]"></span>
</main>

<style>
	.profile {
		@apply m-auto h-full max-w-[1096px];
	}

	/* .profile-info {
		@apply;
	} */

	.banner {
		@apply m-auto h-40 max-w-[1096px] p-4;
	}

	.banner img {
		@apply h-full w-full rounded-xl object-cover;
	}

	.avatar img {
		@apply absolute bottom-[1px] left-12 h-20 w-20 rounded-full border-4  object-cover dark:border-slate-950;
	}

	aside {
		@apply h-screen w-[320px] bg-orange-600;
	}

	h1 {
		@apply absolute bottom-4 left-[140px] text-2xl font-bold;
	}

	main {
		@apply flex;
	}
</style>
