<script lang="ts">
	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import Alert from '$lib/components/alert/alert.svelte';
	import AddForm from './AddForm.svelte';
	import Table from './Table.svelte';

	import Icon from 'lucide-svelte/icons/newspaper';

	import type { PageServerData } from './$types';

	import {
		fetchFilteScienceMagazines,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		newFilterScienceMagazinesRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/scienceMagazines/scienceMagazines';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type {
		FilterScienceMagazinesRequest,
		FilterScienceMagazinesResponse
	} from '$lib/api/model/api/academicProduction/scienceMagazines/scienceMagazine';

	export let data: PageServerData;
	const addScienceMagazineForm = data.addScienceMagazineForm;
	const updateScienceMagazineForm = data.updateScienceMagazineForm;

	let filterScienceMagazinesRequest: FilterScienceMagazinesRequest =
		newFilterScienceMagazinesRequest(5, 1);
	let filterPostGraduateProgramsPromise: Promise<FilterScienceMagazinesResponse> =
		fetchFilteScienceMagazines(filterScienceMagazinesRequest);
	let popoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			filterPostGraduateProgramsPromise = fetchFilteScienceMagazines(filterScienceMagazinesRequest);
		}
	}

	function handleOnFilterChanged(event: CustomEvent) {
		const data: { filter: string } = event.detail;
		filterPostGraduateProgramsPromise = fetchOnFilterChanged(
			data.filter.trim(),
			filterScienceMagazinesRequest,
			popoverFilterDataMap
		);
	}

	function handleOnDetailedFilter() {
		filterPostGraduateProgramsPromise = fetchOnDetailedFilter(
			filterScienceMagazinesRequest,
			popoverFilterDataMap
		).then(({ request, response }) => {
			filterScienceMagazinesRequest = request;
			return response;
		});
	}

	function handlePaginationChanged() {
		filterPostGraduateProgramsPromise = fetchFilteScienceMagazines(filterScienceMagazinesRequest);
	}
</script>

<svelte:head>
	<title>Revistas científicas</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Revistas científicas</h2>
	</div>
	<AddModal
		formComponent={AddForm}
		modalTitle="Crear revista científica"
		formData={addScienceMagazineForm}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await filterPostGraduateProgramsPromise}
		<TableSkeleton tableHeightClass="h-[65vh]" />
	{:then filterScienceMagazinesResponse}
		{#if filterScienceMagazinesResponse.science_magazines}
			<Table
				bind:filterScienceMagazinesRequest
				formData={updateScienceMagazineForm}
				{filterScienceMagazinesResponse}
				bind:popoverFilterDataMap
				comboMessages={undefined}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
				on:filterChanged={handleOnFilterChanged}
				on:detailedFilter={handleOnDetailedFilter}
				on:paginationChanged={handlePaginationChanged}
			/>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay revistas ceintíficas registradas.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
