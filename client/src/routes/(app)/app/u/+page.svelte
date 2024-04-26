<script lang="ts">
	// import type { RowType } from '$lib/server/db/user';

	export let data;
	// const userList = data.data as RowType[];

	console.log(data);
</script>

<svelte:head>
	<title>users</title>
</svelte:head>

<div class="p-4">
	<input
		type="text"
		class="flex mx-auto my-2 w-[420px] p-2 rounded-md border dark:bg-neutral-900"
		placeholder="Search..."
	/>
	<hr class="w-[440px] m-auto" />


</div>
