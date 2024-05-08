<script lang="ts">
	import Editorsn from '$lib/components/Editor.svelte';

	export let editorContent: string;
</script>

<div class="p-4 w-full align-middle">
	<Editorsn bind:editorContent />
</div>
