<script lang="ts">
	export let data;
	import GroupFormComponent from '../../../../lib/components/forms/GroupfForm.svelte';

</script>

<p>Create Group</p>

<GroupFormComponent data={data.form} />
