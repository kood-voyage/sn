<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import type { PageData } from './$types';
	import PeopleSearch from './people-search.svelte';

	export let data: PageData;
	console.log(data.data);
</script>

<svelte:head>
	<title>Chat</title>
</svelte:head>

<!-- chat page -->
<div class="flex w-full">
	<!-- user list -->
	<div class="overflow-scroll h-screen w-14 sm:w-60 bg-slate-50 dark:bg-slate-900">
		<!-- search group / friends -->
		<div class="border-b-2 border-slate-300 dark:border-slate-950 h-22 hidden sm:block">
			<Dialog.Root>
				<div class="p-2 pb-1">
					<Dialog.Trigger
						class="text-sm rounded-md h-fit w-full  py-1 border dark:hover:bg-slate-800 bg-slate-300 dark:bg-slate-950 "
						>New Chat</Dialog.Trigger
					>
				</div>

				<Dialog.Content class="w-fit h-fit max-h-96">
					<PeopleSearch userInfo={data.data} />
					<!-- <Dialog.Header class=""> -->
					<!-- <Dialog.Title>{title}</Dialog.Title>

														<Dialog.Description>
															{content}
														</Dialog.Description>

									

														<div class="py-2 text-center text-sm text-muted-foreground">
															Slide {current} of {count}
														</div>
													</Dialog.Header>
													<div class="w-full h-full"></div>

													<a href={`/app/post/${id}`} class="w-full h-4">to post</a> -->
				</Dialog.Content>
			</Dialog.Root>
			<div class="p-2 pt-1">
				<p class="bg-slate-300 dark:bg-slate-950 rounded py-1 px-2">find friend</p>
			</div>
		</div>

		<!-- user list -->
		<div class=" sm:p-2">
			<!-- header h1 -->

			<div class="hidden sm:block p-2">
				<p class="text-xs">DIRECT MESSAGES</p>
			</div>

			<div class="block sm:hidden h-12"></div>

			<ol>
				<li class="user group hover:bg-slate-300 dark:hover:bg-slate-800 active">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto sm:mx-2"
					/>
					<p>Nikita</p>
				</li>

				<li class="user hover:bg-slate-300 dark:hover:bg-slate-800">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Markus"
						alt="avatar"
						class="m-auto sm:mx-2"
					/>
					<p>Markus</p>
				</li>

				<li class="user hover:bg-slate-300 dark:hover:bg-slate-800">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Harri"
						alt="avatar"
						class="m-auto sm:mx-2"
					/>
					<p>Harri</p>
				</li>
			</ol>
		</div>
	</div>

	<!-- messenger -->

	<div class="w-full">
		<!-- user chat header -->
		<div class="h-12 bg-slate-50 dark:bg-slate-800 px-4">
			<div class="flex p-2 h-full">
				<img
					src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
					alt="avatar"
					class="w-8 rounded-full mr-2"
				/>
				<p class="h-full text-center">Nikita</p>
			</div>
		</div>

		<!-- direct chat -->
		<div class="p-4"></div>
	</div>
</div>

<style>
	div ol li.user {
		@apply flex h-12 cursor-pointer self-center rounded-sm;
	}

	div ol li.user p {
		@apply hidden items-center text-center text-slate-700 dark:text-slate-300;
	}

	@media (min-width: 640px) {
		div ol li.user p {
			@apply ml-2 block items-center text-center align-middle;
		}
	}

	div ol li.user img {
		@apply h-8 items-center justify-center rounded-full;
	}

	.active {
		@apply bg-slate-400 dark:bg-slate-700;
	}

	.user .active p {
		@apply text-slate-900 dark:text-slate-100;
	}
</style>
