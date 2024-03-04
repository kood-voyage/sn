<script lang="ts">
	import type { PageData } from './$types';
	import SettingsForm from './settings-form.svelte';
	export let data: PageData;
</script>

<!-- <div class="w-0 flex hover:-bottom-28">

</div> -->

<div>
	<SettingsForm data={data.form} />

	<p class=" w-full py-2 text-center mt-2">
		already have an account? <a href="/signin"><span class="text-blue-500">sign-in</span></a>
	</p>

	<p>{data.form.data.firstName}</p>
</div>
