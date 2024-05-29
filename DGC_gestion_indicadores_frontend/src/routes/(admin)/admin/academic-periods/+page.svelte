<script lang="ts">
	import type { CommonError } from '$lib/api/model/errors';
	import type { AcademicPeriod } from '$lib/api/model/view/academicPeriod';

	import Alert from '$lib/components/alert/alert.svelte';

	let academicPeriodsPromise: Promise<AcademicPeriod[]> = fetchAcademicPeriods();

	async function fetchAcademicPeriods() {
		const url = `/api/academicPeriod`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = await response.json() as CommonError
			throw errorData
		}
		return (academicPeriodsPromise = response.json());
	}
</script>

<svelte:head>
	<title>Periodos Académicos</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Periodos Académicos</h2>
</div>

{#await academicPeriodsPromise}
	cargando...
{:then periods}
	{#if periods.length > 0}
		{#each periods as period}
			{period.name}
		{/each}
	{:else}
		<Alert title="Sin registros" description={'Ups, no hay periodos académicos'} />
	{/if}
{:catch error}
	<Alert variant="destructive" description={`${error.Detail}: ${error.Message}`} />
{/await}
