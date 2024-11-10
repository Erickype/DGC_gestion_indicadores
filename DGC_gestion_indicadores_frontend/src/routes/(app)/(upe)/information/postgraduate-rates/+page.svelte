<script lang="ts">
	import { filterAcademicPeriodsAuxSchema, filterPostgraduateProgramAuxSchema } from '$lib/utils';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import SuperDebug, { superForm } from 'sveltekit-superforms';

	import type { PageServerData } from './$types';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';
	import { goto } from '$app/navigation';

	import * as Form from '$lib/components/ui/form';

	import Icon from 'lucide-svelte/icons/graduation-cap';

	import PostgraduateProgramsServer from '$lib/components/filters/indicatorsInformation/postgraduate/programsServer.svelte';
	import type { GradeRateListJoined } from '$lib/api/model/api/indicatorsInformation/gradeRateLists';
	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';
	/* import AddForm from './AddForm.svelte';
	import Table from './Table.svelte'; */

	export let data: PageServerData;
	const filterPostgraduateProgramAuxForm = data.filterPostgraduateProgramAuxForm;
	const addPostgraduateCohortListForm = data.addPostgraduateCohortListForm;
	const updatePostgraduateCohortListForm = data.updatePostgraduateCohortListForm;

	const form = superForm(filterPostgraduateProgramAuxForm, {
		validators: zodClient(filterPostgraduateProgramAuxSchema)
	});

	const { form: formData } = form;

	let gradeRateListsPromise: Promise<GradeRateListJoined[]>;

	let formDataPostgraduateProgramID = writable($formData.programID);
	formDataPostgraduateProgramID.subscribe((value) => {
		$formData.programID = value;
		addPostgraduateCohortListForm.data.postgraduate_program_id = $formData.programID
		updatePostgraduateCohortListForm.data.postgraduate_program_id = $formData.programID
		gradeRateListsPromise = fetchGradeRateListsByAcademicPeriod();
	});

	async function fetchGradeRateListsByAcademicPeriod(): Promise<GradeRateListJoined[]> {
		const url = `/api/indicatorsInformation/gradeRateLists/${$formData.programID}`;
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
			$formData.programID = $formDataPostgraduateProgramID;
			gradeRateListsPromise = fetchGradeRateListsByAcademicPeriod();
		}
	}
</script>

<svelte:head>
	<title>Tasas de posgrado</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Tasas de posgrado</h2>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<Form.Field {form} name="programID" class="w-2/3">
		<PostgraduateProgramsServer {formDataPostgraduateProgramID} />
	</Form.Field>

	<!-- <AddModal
		formComponent={AddForm}
		modalTitle="Crear tasa de grado"
		formData={addGradeRateListForm}
		{comboMessages}
		on:created={fetchOnSuccess}
	/> -->
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await gradeRateListsPromise}
		<TableSkeleton tableHeightClass="h-[55vh]" />
	{:then gradeRateLists}
		{#if gradeRateLists && gradeRateLists.length > 0}
			<!-- <Table
				formData={updateGradeRateListForm}
				{gradeRateLists}
				{comboMessages}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
			></Table> -->
		{:else}
			<Alert
				title="Sin registros"
				description={'Ups, no hay tasas de posgrado registradas para el programa.'}
			/>
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
<!-- {#if browser}
	<SuperDebug data={$formData} />
{/if} -->
