<script lang="ts">
	import Message from './message.svelte';

	const user1 = {
		username: 'Surinam',
		image:
			'https://pyxis.nymag.com/v1/imgs/630/6e0/eb215ad90cd826b9e57ff505f54c5c7228-07-avatar.1x.rsquare.w1400.jpg'
	};

	const user2 = {
		username: 'Nikita',
		image:
			'https://static.vecteezy.com/system/resources/thumbnails/002/002/403/small/man-with-beard-avatar-character-isolated-icon-free-vector.jpg'
	};

	const time = '11/12/04 44:30';

	const msg = 'What do you think about our current project crisis? dasdasda sd asd asd asd askjdkjasndjasn dask jdnaskjndjkas ndkanskdjnaskj nakjsnd kandjkanskj dnasjkd nkasnd kas';
</script>

<div class="flex flex-col m-2 overflow-y-scroll h-full">
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />

	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />

	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />

	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
	<Message user={user1} {time} {msg} />

	<Message user={user2} {time} {msg} />
</div>
