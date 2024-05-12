<script lang="ts">
	export let user;

	// const { id, last_name, first_name, username } = user;
</script>

<div class="h-12 bg-neutral-50 dark:bg-neutral-800 px-4 w-full">
	<div class="flex p-2 h-full">
		<img src={user.avatar} alt="avatar" class="w-8 rounded-full mr-2" />
		<p class="h-full text-center">{user.display_name}</p>
		<!-- <p class="">{user.}</p> -->
	</div>
</div>
