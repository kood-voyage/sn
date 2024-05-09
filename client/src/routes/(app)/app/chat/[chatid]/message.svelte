<script lang="ts">
	export let user;
	export let msg;
	export let time;
</script>

<div class="w-full flex mb-4 hover:bg-neutral-200 dark:hover:bg-neutral-700 p-1">
	<img src={user.avatar} class="w-8 h-8 rounded-full mx-2" alt={user.display_name + 'image'} />

	<div class="flex flex-col">
		<div class="flex gap-2">
			<p class=" text-xs font-bold">{user.display_name}</p>
			<p class=" text-[11px] opacity-70">{time}</p>
		</div>

		<div contenteditable="false" class="text-sm">{@html msg}</div>
	</div>
</div>
