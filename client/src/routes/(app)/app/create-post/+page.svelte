<script lang="ts">
	export let data;

	import PostFormComponent from './post-form.svelte';

</script>

<p>Create post</p>

<PostFormComponent data={data.form} />
