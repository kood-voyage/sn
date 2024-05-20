<script lang="ts">
	import Editor from '$lib/components/Editor.svelte';

	export let editorContent: string;
</script>

<div class="p-4 w-full align-middle ">
	<Editor bind:editorContent />
</div>
