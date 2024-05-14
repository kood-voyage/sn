<script lang="ts">
	import type { PageData } from './$types';
	import SettingsForm from '../../../lib/components/forms/LoginForm.svelte';
	export let data: PageData;
</script>

<div>
	<SettingsForm data={data.form} />

	<p class=" w-full py-2 text-center mt-2">
		Don't have an account? <a href="/signup"><span class="text-blue-500">sign-up</span></a>
	</p>
</div>
