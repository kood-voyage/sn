<script lang="ts">
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { writable } from 'svelte/store';

	import Face from 'svelte-radix/Face.svelte';
	import pkg from 'lodash';

	import { data } from '$lib/emojis';
	import Input from './ui/input/input.svelte';
	import { Link2 } from 'svelte-radix';
	import DropdownMenuLabel from './ui/dropdown-menu/dropdown-menu-label.svelte';

	export let editorContent;

	let linkInput;

	const { debounce } = pkg;
	const searchQuery = writable('');
	const debouncedFilterEmojis = debounce((value: string) => searchQuery.set(value), 100);
	let filteredEmojis = [];
	searchQuery.subscribe((value) => {
		filteredEmojis = filterEmojis(value);
	});

	// Function to filter emojis based on search query
	function filterEmojis(value: string) {
		const lowercaseQuery = value.toLowerCase();
		return data.filter((emoji) => emoji.unicodeName.toLowerCase().includes(lowercaseQuery));
	}

	// Debounced version of filterEmojis for performance

	// Function to handle input change and update search query
	function handleInputChange(event: any) {
		const value = event.target.value;
		debouncedFilterEmojis(value);
	}

	function handlePaste(event: ClipboardEvent) {
		event.preventDefault();
		const clipboardData = event.clipboardData || window.clipboardData;
		const paste = clipboardData.getData('text/html');
		const sanitizedPaste = removeInlineStyles(paste);
		document.execCommand('insertHTML', false, sanitizedPaste);
	}

	function removeInlineStyles(html: string): string {
		// Remove inline style attributes
		return html.replace(/<[^>]+style\s*=\s*['"][^'"]*['"][^>]*>/gi, '');
	}

	function insertEmoji(emojiCharacter: string) {
		editorContent += emojiCharacter;
	}

	function insertLink() {
		const link = linkInput;
		if (link) {
			const linkHtml = `<img src="${link}" alt"conent image" class="max-w-[400px] w-full rounded-lg" />`;
			editorContent += linkHtml;
			linkInput = '';
		}
	}
</script>

<div>
	<div class="flex">
		<button id="emoji">
			<DropdownMenu.Root>
				<DropdownMenu.Trigger class="w-[32px] h-[32px]"><Face /></DropdownMenu.Trigger>
				<DropdownMenu.Content>
					<DropdownMenu.Group>
						<DropdownMenu.Label>Emoji</DropdownMenu.Label>
						<Input
							type="text"
							id="search"
							placeholder="Search..."
							bind:value={$searchQuery}
							on:input={handleInputChange}
						/>

						<DropdownMenu.Separator />
						<div class="grid grid-cols-8 min-h-4 max-h-64 overflow-scroll">
							{#each filteredEmojis as emoji}
								<DropdownMenu.Item
									class="w-[32px] h-[32px]"
									on:click={() => insertEmoji(emoji.character)}
								>
									{emoji.character}
								</DropdownMenu.Item>
							{/each}
						</div>
					</DropdownMenu.Group>
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		</button>

		<button id="insertLink">
			<DropdownMenu.Root>
				<DropdownMenu.Trigger class="w-[32px] h-[32px]"><Link2 /></DropdownMenu.Trigger>
				<DropdownMenu.Content class="p-2">
					<DropdownMenuLabel>Insert Link</DropdownMenuLabel>
					<Input type="text" bind:value={linkInput} placeholder="Enter image link..." />
					<button on:click={insertLink}>Insert</button>
					<!-- Insert button for link -->
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		</button>
	</div>

	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div
		id="editor"
		contenteditable="true"
		class="max-h-64 overflow-scroll h-full outline-none w-full border-b"
		on:paste={handlePaste}
		bind:innerHTML={editorContent}
	/>
</div>
