<script lang="ts">
	// let people: UserRowType[] = [];
	// let chats: ChatsWithUsers = {};
	// if (data.ok) {
	// 	people = data.data.usersData;
	// 	chats = data.data.chatsData;
	// }

	// let filteredPeople: UserModel[] = [];

	// $: chats_arr = Object.values(chats);

	// // console.log(chats);
	// $: if (chats_arr.length != 0 && chats_arr != undefined) {
	// 	filteredPeople = Object.entries(chats).map((chat) => {
	// 		// console.log('CHAT >>>', chat);
	// 		const key = chat[0];
	// 		const value = chat[1];

	// 		const users = value.users;
	// 		const group_id = value.group;
	// 		// console.log(value);

	// 		for (const user of users) {
	// 			// console.log(user);
	// 		}

	// 		// console.log(value.group_id != '' && value.group_id != undefined);
	// 		// if (group_id != '' && group_id != undefined) {
	// 		// }

	// 		// console.log('Key >>>', key);
	// 		// console.log('Value >>>', value);
	// 		// if (chat.length == 0) return false
	// 		// if (chat.length == 1) return true
	// 		// const out = chat.filter
	// 		// return chat[0].username.toLowerCase().includes(searchQuery.toLowerCase());
	// 		return;
	// 	});
	// 	filteredPeople = [];
	// 	// .filter((chat) => {
	// 	// 	// console.log('CHAT >>>', chat);
	// 	// 	// const key = chat[0];
	// 	// 	// const value = chat[1];
	// 	// 	// const users = value.users;
	// 	// 	// const group_id = value.group_id;
	// 	// 	// for (const user of users) {
	// 	// 	// 	console.log(user);
	// 	// 	// }
	// 	// 	// // console.log(value.group_id != '' && value.group_id != undefined);
	// 	// 	// if (value.group_id != '' && value.group_id != undefined) {
	// 	// 	// }
	// 	// 	// console.log('Key >>>', key);
	// 	// 	// console.log('Value >>>', value);
	// 	// 	// if (chat.length == 0) return false
	// 	// 	// if (chat.length == 1) return true
	// 	// 	// const out = chat.filter
	// 	// 	// return chat[0].username.toLowerCase().includes(searchQuery.toLowerCase());
	// 	// })
	// 	// .map((val) => {
	// 	// 	// console.log(val);
	// 	// 	return val[0];
	// 	// });
	// 	// .slice(0, 6);

	// 	// console.log(filteredPeople);
	// }
</script>

<div class="w-14 sm:w-60 bg-neutral-50 dark:bg-neutral-900 h-full relative flex flex-col">
	<!-- search group / friends -->
	<div
		class="border-b-2 border-slate-300 dark:border-slate-950 h-12 hidden sm:block sticky top-0 w-full"
	>
		<div class="p-2">
			<p class="bg-neutral-300 dark:bg-neutral-950 rounded py-1 px-2">find friend</p>
		</div>
	</div>

	<!-- user list -->
	<div class="sm:p-2 overflow-y-scroll">
		<!-- header h1 -->

		<div class="hidden sm:block p-2">
			<p class="text-xs">DIRECT MESSAGES</p>
		</div>

		<div class="block sm:hidden h-12"></div>

		<ol class="">
			<!-- user Component  -->
			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 active rounded">
				<a href="/app/chat/dsadsd">
					<div class="flex p-1">
						<img
							src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
							alt="avatar"
							class="m-auto w-8 h-8 rounded-full sm:mx-2"
						/>
						<div>
							<p class="text-sm font-[600]">Nikita</p>
						</div>
					</div>
				</a>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>

			<li class="user group hover:bg-neutral-300 dark:hover:bg-neutral-800 rounded">
				<div class="flex p-1">
					<img
						src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
						alt="avatar"
						class="m-auto w-8 h-8 rounded-full sm:mx-2"
					/>
					<div>
						<p class="text-sm font-[600]">Nikita</p>
					</div>
				</div>
			</li>
		</ol>
	</div>
</div>

<style>
	.active {
		@apply bg-neutral-400 dark:bg-neutral-700;
	}

	.user .active p {
		@apply text-slate-900 dark:text-slate-100;
	}
</style>
