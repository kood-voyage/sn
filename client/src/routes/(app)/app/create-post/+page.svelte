<script lang="ts">
	export let data;

	import PostFormComponent from '$lib/components/post-form.svelte';

	console.log(data);
</script>

<p>Create post</p>

<PostFormComponent data={data.form} />
