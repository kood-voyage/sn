<script lang="ts">
	import Editor from '$lib/components/Editor.svelte';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { commentSchema } from '../../comment-schema';
	import { superForm } from 'sveltekit-superforms';
	import { page } from '$app/stores';
	import { PaperPlane } from 'svelte-radix';
	import { commentsStore } from '$lib/store/comments-store';

	import { v4 as uuidv4 } from 'uuid';
	import { createComment } from '$lib/client/api/post-requests';
	import { invalidate, invalidateAll } from '$app/navigation';

	// import * as Form from '$lib/components/ui/form';

	export let data;
	export let post_id: string;

	let editorContent: string;

	const form = superForm(data, {
		validators: zodClient(commentSchema),
		onSubmit: async ({ formData }) => {
			formData.set('post_id', post_id);
			formData.set('content', editorContent);
			formData.set('user_name', $page.data.data.username);
			formData.set('user_avatar', $page.data.data.avatar);

			let temporary = {
				id: uuidv4(),
				content: editorContent,
				post_id: post_id
			};

			const resp = await createComment(temporary);

			if (resp.ok) {
			}

			editorContent = '';
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
	<div class="w-[99%] text-wrap">
		<Editor bind:editorContent />
	</div>
	<!-- 
	<Form.Field {form} name="content" class="w-[99%] text-wrap">
		<Form.Control let:attrs>
			<Editor bind:editorContent {...attrs} />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field> -->

	<button type="submit" class="absolute right-2"><PaperPlane /></button>
</form>
