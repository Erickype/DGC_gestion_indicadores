<script lang="ts">
	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import Alert from '$lib/components/alert/alert.svelte';
	/* mport AddForm from './AddForm.svelte';
	import Table from './Table.svelte'; */

	import Icon from 'lucide-svelte/icons/brain-circuit';

	import type { PageServerData } from './$types';

	import {
		fetchFilterPostgraduatePrograms,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		newFilterPostgraduateProgramsRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/indicatorsInformation/postgraduate/programs';
	import type {
		FilterPostgraduateProgramsRequest,
		FilterPostGraduateProgramsResponse
	} from '$lib/api/model/api/indicatorsInformation/postgraduate';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';

	export let data: PageServerData;
	const addPostGraduateProgramForm = data.addPostGraduateProgramForm;
	const updatePostGraduateProgramForm = data.updatePostGraduateProgramForm;

	let filterPostgraduateProgramsRequest: FilterPostgraduateProgramsRequest =
		newFilterPostgraduateProgramsRequest(5, 1);
	let filterPostGraduateProgramsPromise: Promise<FilterPostGraduateProgramsResponse> =
		fetchFilterPostgraduatePrograms(filterPostgraduateProgramsRequest);
	let popoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			filterPostGraduateProgramsPromise = fetchFilterPostgraduatePrograms(
				filterPostgraduateProgramsRequest
			);
		}
	}

	function handleOnFilterChanged(event: CustomEvent) {
		const data: { filter: string } = event.detail;
		filterPostGraduateProgramsPromise = fetchOnFilterChanged(
			data.filter.trim(),
			filterPostgraduateProgramsRequest,
			popoverFilterDataMap
		);
	}

	function handleOnDetailedFilter() {
		filterPostGraduateProgramsPromise = fetchOnDetailedFilter(
			filterPostgraduateProgramsRequest,
			popoverFilterDataMap
		).then(({ request, response }) => {
			filterPostgraduateProgramsRequest = request;
			return response;
		});
	}

	function handlePaginationChanged() {
		filterPostGraduateProgramsPromise = fetchFilterPostgraduatePrograms(
			filterPostgraduateProgramsRequest
		);
	}
</script>

<svelte:head>
	<title>Programas de posgrado</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Programas de posgrado</h2>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<!-- <AddModal
		formComponent={AddForm}
		modalTitle="Crear libro o capítulo"
		formData={addPostGraduateProgramForm}
		on:created={fetchOnSuccess}
	/> -->
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await filterPostGraduateProgramsPromise}
		<TableSkeleton />
	{:then filterPostGraduateProgramsResponse}
		{#if filterPostGraduateProgramsResponse.postgraduate_programs}
			<!-- <Table
				bind:filterPostgraduateProgramsRequest
				formData={updatePostGraduateProgramForm}
				{filterPostGraduateProgramsResponse}
				bind:popoverFilterDataMap
				comboMessages={undefined}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
				on:filterChanged={handleOnFilterChanged}
				on:detailedFilter={handleOnDetailedFilter}
				on:paginationChanged={handlePaginationChanged}
			></Table> -->
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay programas posgrado registrados.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
