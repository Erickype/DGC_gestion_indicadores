<script lang="ts">
	import type { PageServerData } from './$types';
	import { goto } from '$app/navigation';

	import { toast } from 'svelte-sonner';

	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import AddForm from './AddForm.svelte';

	import type { Message } from '$lib/components/combobox/combobox';
	import type { Teacher } from '$lib/api/model/api/teacher';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';

	export let data: PageServerData;
	const addTeacherForm = data.addTeacherForm;

	const academicPeriodsData = data.academicPeriodsData;

	const comboMessages: Message[][] = [
		data.peopleData.messages,
		data.careersData.messages,
		data.dedicationsData.messages,
		data.scaledGradesData.messages,
		data.contractTypesData.messages
	];
	
	let selectedAcademicPeriod: number = academicPeriodsData.periods.at(
		academicPeriodsData.periods.length - 1
	)!.ID;
	
	$: addTeacherForm.data.academicPeriod = selectedAcademicPeriod;
	
	let teachersPromise: Promise<Teacher[]> = fetchTeachers();

	async function fetchTeachers() {
		const url = `/api/teacher/byAcademicPeriodID/${selectedAcademicPeriod}`;
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
		return (teachersPromise = response.json());
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			fetchTeachers();
		}
	}
</script>

<svelte:head>
	<title>Docentes</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Docentes</h2>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<AcademicPeriodCombo
		messages={academicPeriodsData.messages}
		bind:selectedValue={selectedAcademicPeriod}
		on:message={fetchTeachers}
	></AcademicPeriodCombo>

	<AddModal
		formComponent={AddForm}
		modalTitle="Crear docente"
		formData={addTeacherForm}
		{comboMessages}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full place-content-center px-8">
	{#await teachersPromise}
		cargando...
	{:then teachers}
		{#if teachers.length > 0}
			<!-- <FacultiesTable
				formData={updateFacultyFormData}
				{faculties}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
			></FacultiesTable> -->
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay docentes registrados.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
