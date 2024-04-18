<script lang="ts">
	import Editor from './editor.svelte';
	import Header from './header.svelte';
	import Userlist from './userlist.svelte';

	const user = {
		id: 'asdasdasd',
		username: 'Nikita',
		first_name: 'Aladin',
		last_name: 'Khan'
	};
</script>

<div class="flex w-full h-screen">
	<Userlist />

	<div class="w-full flex flex-col h-screen">
		<Header {user} />

		<div class="h-full overflow-y-scroll">
			<slot />
		</div>

		<div class="w-full">
			<Editor />
		</div>
	</div>
</div>
