import { writable } from "svelte/store";
	async function fetchEmojis() {
		let result;
		await fetch('https://emoji-api.com/emojis?access_key=a4a920bbf071e0e845226b1066d53712562b60c6')
			.then((response) => response.json())
			.then((data) => {
				result = data;
			})
			.catch((error) => {
				console.log('error: ' + error);
				result = {};
			});
		return result;
	}

export const emojiStore = writable(fetchEmojis())