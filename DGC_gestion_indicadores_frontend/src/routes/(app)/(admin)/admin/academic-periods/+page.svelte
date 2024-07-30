<script lang="ts">
	import type { CommonError } from '$lib/api/model/errors';
	import type { AcademicPeriod } from '$lib/api/model/view/academicPeriod';

	import PeriodsTable from './Table.svelte';

	import Alert from '$lib/components/alert/alert.svelte';
	import type { PageData } from './$types';
	import AddModal from './AddModal.svelte';

	export let data: PageData;

	const addAcademicPeriodFormData = data.addAcademiPeriodForm;
	const updateAcademicPeriodFormData = data.updateAcademicPeriodForm;

	let academicPeriodsPromise: Promise<AcademicPeriod[]> = fetchAcademicPeriods();

	async function fetchAcademicPeriods() {
		const url = `/api/academicPeriod`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			throw errorData;
		}
		return (academicPeriodsPromise = response.json());
	}

	function handleCreated(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			fetchAcademicPeriods();
		}
	}

	function handleDeleted(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			fetchAcademicPeriods();
		}
	}

	function handleUpdated(event: any) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			fetchAcademicPeriods();
		}
	}
</script>

<svelte:head>
	<title>Periodos Académicos</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Periodos Académicos</h2>
	<AddModal formData={addAcademicPeriodFormData} on:created={handleCreated} />
</div>

<div class="mx-auto flex w-full place-content-center px-8">
	{#await academicPeriodsPromise}
		cargando...
	{:then periods}
		{#if periods.length > 0}
			<PeriodsTable
				formData={updateAcademicPeriodFormData}
				{periods}
				on:updated={handleUpdated}
				on:deleted={handleDeleted}
			></PeriodsTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay periodos académicos'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={`${error.Detail}: ${error.Message}`} />
	{/await}
</div>
