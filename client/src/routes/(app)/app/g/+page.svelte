<script lang="ts">
const groups = [
  { name: "Team Valor", members: 5, description: "Fearless warriors fighting for victory." },
  { name: "Pixel Rangers", members: 8, description: "Guardians of the gaming realm, pixel by pixel." },
  { name: "Quest Keepers", members: 3, description: "Explorers on an epic quest for glory." },
  { name: "Level Legends", members: 6, description: "Ascending to legendary gaming heights." },
  { name: "Game Guardians", members: 4, description: "Protectors of the gaming universe." },
  { name: "Console Crusaders", members: 7, description: "Crusading through realms via console adventures." },
  { name: "Joystick Jugglers", members: 2, description: "Masters of precision and control." },
  { name: "Elite Empire", members: 9, description: "Ruling the gaming empire with elite skills." },
  { name: "Respawn Rebels", members: 1, description: "Rebels with a cause, rising again and again." },
  { name: "Warrior Wizards", members: 5, description: "Wizards wielding gaming magic on the battlefield." }
];

</script>

<div class="p-12">
	<p class="text-md">Groups</p>
	{#each groups as name}
		<div class="flex m-auto w-[420px] p-2 hover:bg-slate-200 dark:hover:bg-slate-900 rounded-md">
			<a href="/app/g/{name.name}" class="flex">
				<img
					src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed={name.name}"
					alt="user-avatar"
					class="w-16 mr-2"
				/>
				<div>
					<p class="font-bold">{name.name}</p>
					<p class="text-sm text-blue-500">Members: <span>{name.members}</span></p>
                    <p class="text-sm text-slate-500">{name.description}</p>
				</div>
			</a>
			<!-- <div class="flex">
				<button class="text-sm p-1 bg-blue-500">follow</button>
			</div> -->
		</div>
	{/each}
</div>
