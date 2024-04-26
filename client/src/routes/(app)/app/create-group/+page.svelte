<script lang="ts">
	export let data;

	import GroupFormComponent from './group-form.svelte';

	// console.log(data);
</script>

<p>Create Group</p>

<GroupFormComponent data={data.form} />
