<script lang="ts">
	import { filterAcademicPeriodsAuxSchema } from '$lib/utils';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { superForm } from 'sveltekit-superforms';

	import type { PageServerData } from './$types';
	import { writable } from 'svelte/store';
	import { goto } from '$app/navigation';

	import Icon from 'lucide-svelte/icons/projector';

	import type { ResearchInnovationProjectListJoined } from '$lib/api/model/api/indicatorsInformation/researchInnovationProjectLists';
	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';
	import Table from './Table.svelte';

	let researchInnovationProjectListsPromise: Promise<ResearchInnovationProjectListJoined[]> =
		fetchResearchInnovationProjectListsByAcademicPeriod();

	async function fetchResearchInnovationProjectListsByAcademicPeriod(): Promise<
		ResearchInnovationProjectListJoined[]
	> {
		const url = `/api/indicatorsInformation/researchInnovationProjectLists`;
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
			researchInnovationProjectListsPromise = fetchResearchInnovationProjectListsByAcademicPeriod();
		}
	}
</script>

<svelte:head>
	<title>Proyectos de investigación e innovación</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Proyectos de investigación e innovación</h2>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<!-- <AddModal
		formComponent={AddForm}
		modalTitle="Crear tasa de grado"
		formData={addGradeRateListForm}
		{comboMessages}
		on:created={fetchOnSuccess}
	/> -->
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await researchInnovationProjectListsPromise}
		<TableSkeleton tableHeightClass="h-[55vh]" />
	{:then researchInnovationProjectLists}
		{#if researchInnovationProjectLists && researchInnovationProjectLists.length > 0}
			<Table
				{researchInnovationProjectLists}
				comboMessages={undefined}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
			></Table>
		{:else}
			<Alert
				title="Sin registros"
				description={'Ups, no hay proyectos de investigación e innovación registrados.'}
			/>
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
