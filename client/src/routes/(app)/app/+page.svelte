<script lang="ts">
	import { currentUserStore } from '$lib/store/user-store';
	import { onMount } from 'svelte';

	let webSocket;
	let messages = [];

	// Function to establish WebSocket connection
	function connectWebSocket() {
		webSocket = new WebSocket('wss://your-websocket-server.com');

		webSocket.onopen = () => {
			console.log('WebSocket connection established');
		};

		webSocket.onmessage = (event) => {
			console.log('Message from server ', event.data);
			messages = [...messages, event.data]; // Update messages to render
		};

		webSocket.onerror = (error) => {
			console.error('WebSocket error: ', error);
		};

		webSocket.onclose = () => {
			console.log('WebSocket connection closed');
			// Optionally, implement reconnection logic here
		};
	}

	// Establish connection when component mounts
	onMount(() => {
		connectWebSocket();
	});

	function sendMessage(message) {
		if (webSocket.readyState === WebSocket.OPEN) {
			webSocket.send(message);
		} else {
			console.error('WebSocket is not open.');
		}
	}

	import { onDestroy } from 'svelte';

	onDestroy(() => {
		if (webSocket) {
			webSocket.close();
		}
	});
</script>

<ul>
	{#each messages as message (message)}
		<li>{message}</li>
	{/each}
</ul>

<!-- Example send message form -->
<form on:submit|preventDefault={() => sendMessage('Hello WebSocket!')}>
	<button type="submit">Send Message</button>
</form>

<h1>HERE WILL BE HOME PAGE</h1>
<h1>{$currentUserStore.username}</h1>
