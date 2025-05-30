<script lang="ts">
	import { filterAcademicPeriodsAuxSchema } from '$lib/utils';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { superForm } from 'sveltekit-superforms';

	import type { PageServerData } from './$types';
	import { writable } from 'svelte/store';
	import { goto } from '$app/navigation';

	import * as Form from '$lib/components/ui/form';

	import Icon from 'lucide-svelte/icons/graduation-cap';

	import AcademicPeriodsServer from '$lib/components/filters/academicPeriods/academicPeriodsServer.svelte';
	import type { GradeRateListJoined } from '$lib/api/model/api/indicatorsInformation/gradeRateLists';
	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';
	import AddForm from './AddForm.svelte';
	import Table from './Table.svelte';

	export let data: PageServerData;
	const filterAcademicPeriodsAuxForm = data.filterAcademicPeriodsAuxForm;
	const addGradeRateListForm = data.addGradeRateListForm;
	const updateGradeRateListForm = data.updateGradeRateListForm;

	const comboMessages: Message[][] = [data.careersData.messages];

	const form = superForm(filterAcademicPeriodsAuxForm, {
		validators: zodClient(filterAcademicPeriodsAuxSchema)
	});

	const { form: formData } = form;

	$formData.academic_period_id = data.academicPeriodsData.periods.at(0)!.ID;

	let gradeRateListsPromise: Promise<GradeRateListJoined[]>;

	let formDataAcademicPeriodID = writable($formData.academic_period_id);
	formDataAcademicPeriodID.subscribe((value) => {
		$formData.academic_period_id = value;
		addGradeRateListForm.data.academic_period_id = $formData.academic_period_id;
		updateGradeRateListForm.data.academic_period_id = $formData.academic_period_id;
		gradeRateListsPromise = fetchGradeRateListsByAcademicPeriod();
	});

	async function fetchGradeRateListsByAcademicPeriod(): Promise<GradeRateListJoined[]> {
		const url = `/api/indicatorsInformation/gradeRateLists/${$formData.academic_period_id}`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			throw errorData;
		}
		return response.json();
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			$formData.academic_period_id = $formDataAcademicPeriodID;
			gradeRateListsPromise = fetchGradeRateListsByAcademicPeriod();
		}
	}
</script>

<svelte:head>
	<title>Tasas de grado</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Tasas de grado</h2>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<Form.Field {form} name="academic_period_id" class="w-1/3">
		<AcademicPeriodsServer {formDataAcademicPeriodID} />
	</Form.Field>

	<AddModal
		formComponent={AddForm}
		modalTitle="Crear tasa de grado"
		formData={addGradeRateListForm}
		{comboMessages}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await gradeRateListsPromise}
		<TableSkeleton tableHeightClass="h-[55vh]" />
	{:then gradeRateLists}
		{#if gradeRateLists && gradeRateLists.length > 0}
			<Table
				formData={updateGradeRateListForm}
				{gradeRateLists}
				{comboMessages}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
			></Table>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay tasas de grado registradas.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
