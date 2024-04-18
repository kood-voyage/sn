<script lang="ts">
	import Editor from '$lib/components/Editor.svelte';
	import * as Form from '$lib/components/ui/form';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { commentSchema } from '../../comment-schema';
	import { superForm } from 'sveltekit-superforms';
	import { page } from '$app/stores';
	import { editorValue } from '$lib/store/editor-store';
	import { PaperPlane } from 'svelte-radix';

	export let data;
	export let post_id;

	console.log($page.data.data.username);

	const form = superForm(data, {
		validators: zodClient(commentSchema),
		onSubmit: ({ formData }) => {
			formData.set('post_id', post_id);
			formData.set('content', $editorValue);
			formData.set('user_name', $page.data.data.username);
			formData.set('user_avatar', $page.data.data.avatar);
		}
	});

	const { form: formData, enhance } = form;

	// Function to handle form submission
	const handleSubmit = () => {
		formData.submit();
	};

	// Add event listener for Enter key press
	const handleKeyPress = (event: KeyboardEvent) => {
		if (event.key === 'Enter') {
			handleSubmit();
		}
	};
</script>

<form
	method="POST"
	action="?/commentSubmit"
	enctype="multipart/form-data"
	use:enhance
	class="bg-neutral-700/50 rounded-lg relative w-full flex p-2"
>
	<div class="w-[99%] text-wrap ">
		<Editor />
	</div>

	<button type="submit" class="absolute right-2"><PaperPlane /></button>
</form>
