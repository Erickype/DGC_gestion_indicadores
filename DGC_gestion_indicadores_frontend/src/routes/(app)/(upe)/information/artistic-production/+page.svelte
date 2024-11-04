<script lang="ts">
	import type { PageServerData } from './$types';
	import { goto } from '$app/navigation';

	import Icon from 'lucide-svelte/icons/brush';

	import type { ArtisticProductionListJoined } from '$lib/api/model/api/indicatorsInformation/artisticProductionLists';
	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';
	import AddForm from './AddForm.svelte';
	import Table from './Table.svelte';

	export let data: PageServerData;
	const addResearchInnovationProjectListForm = data.addResearchInnovationProjectListForm;
	const updateResearchInnovationProjectListForm = data.updateResearchInnovationProjectListForm;

	let artisticProductionListsPromise: Promise<ArtisticProductionListJoined[]> =
		fetchArtisticProductionListsByAcademicPeriod();

	async function fetchArtisticProductionListsByAcademicPeriod(): Promise<
		ArtisticProductionListJoined[]
	> {
		const url = `/api/indicatorsInformation/artisticProductionLists`;
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
			artisticProductionListsPromise = fetchArtisticProductionListsByAcademicPeriod();
		}
	}
</script>

<svelte:head>
	<title>Producción artística</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Producción artística</h2>
	</div>
	<AddModal
		formComponent={AddForm}
		modalTitle="Crear Producción artística"
		formData={addResearchInnovationProjectListForm}
		comboMessages={undefined}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await artisticProductionListsPromise}
		<TableSkeleton tableHeightClass="h-[65vh]" />
	{:then artisticProductionLists}
		{#if artisticProductionLists && artisticProductionLists.length > 0}
			<Table
				formData={updateResearchInnovationProjectListForm}
				{artisticProductionLists}
				comboMessages={undefined}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
			></Table>
		{:else}
			<Alert
				title="Sin registros"
				description={'Ups, no hay producciones artísticas registradas.'}
			/>
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
