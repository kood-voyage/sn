<script lang="ts">
	import { currentUserStore } from '$lib/store/user-store';

	const socket = new WebSocket('ws://localhost:8080/test/ws');
	socket.onmessage = function (event) {
		console.log('Received message:', event.data);
	};
</script>

<h1>HERE WILL BE HOME PAGE</h1>
<h1>{$currentUserStore.username}</h1>
