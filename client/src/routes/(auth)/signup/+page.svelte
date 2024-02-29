<script lang="ts">
	import type { PageData } from './$types';
	import SettingsForm from './settings-form.svelte';
	export let data: PageData;
</script>

<!-- <div class="w-0 flex hover:-bottom-28">

</div> -->


	<SettingsForm data={data.form} />
