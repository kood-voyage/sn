<script lang="ts">
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { writable } from 'svelte/store';
	import Face from 'svelte-radix/Face.svelte';
	import FontBold from 'svelte-radix/FontBold.svelte';
	import FontItalic from 'svelte-radix/FontItalic.svelte';

	import pkg from 'lodash';
	const { debounce } = pkg;

	import { data } from '$lib/emojis';
	import Input from './ui/input/input.svelte';

	// Declare searchQuery as a writable store
	const searchQuery = writable('');

	// Function to filter emojis based on search query
	function filterEmojis(value: string) {
		const lowercaseQuery = value.toLowerCase();
		return data.filter((emoji) => emoji.unicodeName.toLowerCase().includes(lowercaseQuery));
	}

	// Debounced version of filterEmojis for performance
	const debouncedFilterEmojis = debounce((value: string) => searchQuery.set(value), 100);

	// Function to handle input change and update search query
	function handleInputChange(event: any) {
		const value = event.target.value;
		debouncedFilterEmojis(value);
	}

	// Subscribe to changes in searchQuery and update filtered emojis
	let filteredEmojis = [];
	searchQuery.subscribe((value) => {
		filteredEmojis = filterEmojis(value);
	});

	function toggleBold() {
		document.execCommand('bold');
	}

	function toggleItalic() {
		document.execCommand('italic');
	}

	function insertEmojiAtEnd(emoji: any) {
		const editor = document.getElementById('editor');
		if (editor) {
			// Create a new text node containing the emoji character
			const emojiNode = document.createTextNode(emoji.character);
			// Append the emoji node to the editor
			editor.appendChild(emojiNode);
		}
	}
</script>

<div>
	<div class="flex">
		<button on:click={toggleBold}><FontBold /></button>
		<button on:click={toggleItalic}><FontItalic /></button>
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
									on:click={() => insertEmojiAtEnd(emoji)}
								>
									{emoji.character}
								</DropdownMenu.Item>
							{/each}
						</div>
					</DropdownMenu.Group>
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		</button>
	</div>

	<div
		id="editor"
		contenteditable="true"
		class="max-h-64 overflow-scroll border rounded-md p-2"
	></div>
</div>
