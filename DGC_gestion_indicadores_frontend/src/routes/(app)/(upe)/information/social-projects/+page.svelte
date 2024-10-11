<script lang="ts">
	import { filterAcademicPeriodsAuxSchema } from '$lib/utils';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { superForm } from 'sveltekit-superforms';

	import type { PageServerData } from './$types';
	import { writable } from 'svelte/store';

	import * as Form from '$lib/components/ui/form';

	import Icon from 'lucide-svelte/icons/users-round';

	import AcademicPeriodsServer from '$lib/components/filters/academicPeriods/academicPeriodsServer.svelte';

	export let data: PageServerData;
	const filterAcademicPeriodsAuxForm = data.filterAcademicPeriodsAuxForm;

	const form = superForm(filterAcademicPeriodsAuxForm, {
		validators: zodClient(filterAcademicPeriodsAuxSchema)
	});

	const { form: formData } = form;

	$formData.academic_period_id = 7

	let formDataAcademicPeriodID = writable($formData.academic_period_id);
	formDataAcademicPeriodID.subscribe((value) => {
		$formData.academic_period_id = value;
		fetchPostFilterSocialProjectLists();
	});

	function fetchPostFilterSocialProjectLists() {
		console.log('loading...');
	}
</script>

<svelte:head>
	<title>Proyectos Vinculación</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Proyectos Vinculación</h2>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<Form.Field {form} name="academic_period_id">
		<AcademicPeriodsServer {formDataAcademicPeriodID} />
	</Form.Field>
</div>
