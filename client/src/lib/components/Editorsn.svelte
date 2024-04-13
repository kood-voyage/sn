<script lang="ts">
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { writable } from 'svelte/store';
	import Face from 'svelte-radix/Face.svelte';
	import FontBold from 'svelte-radix/FontBold.svelte';
	import FontItalic from 'svelte-radix/FontItalic.svelte';

	const dispatch = createEventDispatcher();

	let editorValue;

	import pkg from 'lodash';
	const { debounce } = pkg;

	import { data } from '$lib/emojis';
	import Input from './ui/input/input.svelte';
	import { createEventDispatcher } from 'svelte';

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

	interface CursorPosition {
		line: number;
		position: number;
	}
	let lastFocusedPosition: CursorPosition | null = null;

	lastFocusedPosition = {
		line: 0,
		position: 0
	};

	function insertEmojiAtLastFocus(emoji: any) {
		if (lastFocusedPosition !== null) {
			const editor = document.getElementById('editor');
			if (editor) {
				const emojiCharacter = emoji.character;
				const lines = editor.innerText.split('\n');
				const currentLine = lines[lastFocusedPosition.line];
				const textBeforeCursor = currentLine.slice(0, lastFocusedPosition.position);
				const textAfterCursor = currentLine.slice(lastFocusedPosition.position);
				lines[lastFocusedPosition.line] = textBeforeCursor + emojiCharacter + textAfterCursor;
				editor.innerText = lines.join('\n');
			}
		}
	}
	function storeLastFocusPosition() {
		const editor = document.getElementById('editor');

		if (editor) {
			const selection = window.getSelection();
			if (selection && selection.anchorNode && selection.focusNode) {
				const range = document.createRange();
				range.setStart(selection.anchorNode, selection.anchorOffset);
				range.setEnd(selection.focusNode, selection.focusOffset);
				const rect = range.getBoundingClientRect();
				const lineHeight = parseInt(getComputedStyle(editor).lineHeight, 10);
				const line = Math.floor((rect.top - editor.getBoundingClientRect().top) / lineHeight);
				const position = range.startOffset;
				lastFocusedPosition = { line, position };
			}
		}
	}

	function handleEditorInput(event: InputEvent) {
		const editor = event.target as HTMLElement;
		const content = editor.innerText;
		const lines = content.split('\n');
		const cursorPosition = getCaretPosition(editor);

		// console.log(cursorPosition);

		// Check if a newline character is inserted
		if (lines.length > lastFocusedPosition.line + 1) {
			// Newline inserted, handle the event here
			// For example, you could update the lastFocusedPosition:
			lastFocusedPosition = {
				line: cursorPosition.line,
				position: cursorPosition.position
			};
		}

		// const html = Array.prototype.reduce.call(
		// 	editorValue,
		// 	function (html, node) {
		// 		return html + (node.outerHTML || node.nodeValue);
		// 	},
		// 	''
		// );

		editorValue = event.target;
		console.log('Editor val >>>', editorValue);
		dispatch('valueChange', editorValue);
	}

	function getCaretPosition(element: HTMLElement): CursorPosition {
		const selection = window.getSelection();
		if (selection) {
			const range = selection.getRangeAt(0);
			const preRange = range.cloneRange();
			preRange.selectNodeContents(element);
			preRange.setEnd(range.startContainer, range.startOffset);
			const line = preRange.toString().split('\n').length - 1;
			const position = preRange.toString().length - preRange.toString().lastIndexOf('\n') - 1;
			return { line, position };
		}
		return { line: 0, position: 0 };
	}
</script>

<div>
	<div class="flex">
		<button on:click={toggleBold}><FontBold /></button>
		<button on:click={toggleItalic}><FontItalic /></button>
		<button id="emoji" on:click|stopPropagation={storeLastFocusPosition}>
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
									on:click={() => insertEmojiAtLastFocus(emoji)}
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
		on:keydown={handleEditorInput}
		on:focus={handleEditorInput}
	/>
</div>
