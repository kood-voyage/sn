<script lang="ts">
	import type { PageData } from './$types';
	import SettingsForm from './settings-form.svelte';
	export let data: PageData;

	console.log(data);
	console.log('HI');
</script>

<div>
	<SettingsForm data={data.form} />

	<p class=" w-full py-2 text-center mt-2">
		already have an account? <a href="/signin"><span class="text-blue-500">sign-in</span></a>
	</p>

	<p>{data.form.data.firstName}</p>
</div>
