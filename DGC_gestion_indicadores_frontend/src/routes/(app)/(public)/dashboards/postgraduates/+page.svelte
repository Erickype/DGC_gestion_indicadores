<script lang="ts">
	import SuperDebug, { superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { filterCohortsAuxSchema } from '$lib/utils';

	import type { PageServerData } from './$types';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';
	import { goto } from '$app/navigation';

	import * as Form from '$lib/components/ui/form';

	import Icon from 'lucide-svelte/icons/circle-gauge';

	import CohortsServer from '$lib/components/filters/indicatorsInformation/postgraduate/cohortsServer.svelte';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';

	export let data: PageServerData;
	const filterCohortsAuxForm = data.filterCohortsAuxForm;

	const form = superForm(filterCohortsAuxForm, {
		validators: zodClient(filterCohortsAuxSchema)
	});

	const { form: formData } = form;

	let formDataCohortYear = writable($formData.year);
	formDataCohortYear.subscribe((value) => {
		$formData.year = value;
	});
</script>

<svelte:head>
	<title>Dashboards posgrado</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Dashboards posgrado</h2>
	</div>
	<Form.Field {form} name="year" class="w-1/3">
		<CohortsServer {formDataCohortYear} />
	</Form.Field>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8"></div>
<!-- {#if browser}
	<SuperDebug data={$formData} />
{/if} -->
