<script lang="ts">
	import type { EvaluationPeriod } from '$lib/api/model/view/evaluationPeriod';
	import type { PageData } from './$types';
	import AddModal from './AddModal.svelte';
	import PeriodsTable from './Table.svelte';
	import Alert from '$lib/components/alert/alert.svelte';

	export let data: PageData;

	const formData = data.addEvaluationPeriodForm;

	let evaluationPeriodsPromise: Promise<EvaluationPeriod[]> = fetchEvaluationPeriods();

	async function fetchEvaluationPeriods() {
		const url = `/api/evaluationPeriod`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			throw new Error(response.statusText);
		}
		return (evaluationPeriodsPromise = response.json());
	}

	function handleCreated(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			fetchEvaluationPeriods();
		}
	}

	function handleDeleted(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			fetchEvaluationPeriods();
		}
	}
</script>

<svelte:head>
	<title>Periodos Evaluación</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Periodos de Evaluación</h2>
	<AddModal {formData} on:created={handleCreated} />
</div>
<div class="mx-auto flex w-full place-content-center px-8">
	{#await evaluationPeriodsPromise}
		cargando...
	{:then periods}
		{#if periods.length > 0}
			<PeriodsTable {formData} {periods} on:deleted={handleDeleted}></PeriodsTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay periodos de evaluación'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
