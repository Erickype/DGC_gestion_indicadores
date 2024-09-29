<script lang="ts">
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import AddForm from './AddForm.svelte';
	import Table from './Table.svelte';

	import Icon from 'lucide-svelte/icons/book-marked';

	import type { PageServerData } from './$types';
	import type {
		FilterAcademicProductionListsByAcademicPeriodRequest,
		FilterAcademicProductionListsByAcademicPeriodResponse
	} from '$lib/api/model/api/indicatorsInformation/academicProductionLists';
	import {
		fetchFilterAcademicProductionLists,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		newFilterAcademicProductionListsByAcademiPeriodRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/indicatorsInformation/academicProductionLists/academicProductionLists';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type { Message } from '$lib/components/combobox/combobox';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import Alert from '$lib/components/alert/alert.svelte';

	export let data: PageServerData;
	const addAcademicProductionForm = data.academicProductionForm;

	const academicPeriodsData = data.academicPeriodsData;

	const comboMessages: Message[][] = [
		data.scienceMagazinesData.messages,
		data.impactCoefficientsData.messages,
	];

	let selectedAcademicPeriod: number = academicPeriodsData.periods.at(0)!.ID;

	$: {
		addAcademicProductionForm.data.academicPeriod = selectedAcademicPeriod;
	}

	let filterAcademicProductionListRequest: FilterAcademicProductionListsByAcademicPeriodRequest =
		newFilterAcademicProductionListsByAcademiPeriodRequest(5, 1, selectedAcademicPeriod);
	let filterAcademicProductionListPromise: Promise<FilterAcademicProductionListsByAcademicPeriodResponse> =
		fetchFilterAcademicProductionLists(filterAcademicProductionListRequest);
	let popoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function fetchOnAcademicPeriodChange() {
		filterAcademicProductionListRequest.academic_period_id = selectedAcademicPeriod;
		filterAcademicProductionListPromise = fetchFilterAcademicProductionLists(
			filterAcademicProductionListRequest
		);
	}
	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			filterAcademicProductionListPromise = fetchFilterAcademicProductionLists(
				filterAcademicProductionListRequest
			);
		}
	}

	function handleOnFilterChanged(event: CustomEvent) {
		const data: { filter: string } = event.detail;
		filterAcademicProductionListPromise = fetchOnFilterChanged(
			data.filter.trim(),
			filterAcademicProductionListRequest,
			popoverFilterDataMap
		);
	}

	function handleOnDetailedFilter() {
		filterAcademicProductionListPromise = fetchOnDetailedFilter(
			filterAcademicProductionListRequest,
			popoverFilterDataMap
		).then(({ request, response }) => {
			filterAcademicProductionListRequest = request;
			return response;
		});
	}

	function handlePaginationChanged() {
		filterAcademicProductionListPromise = fetchFilterAcademicProductionLists(
			filterAcademicProductionListRequest
		);
	}
</script>

<svelte:head>
	<title>Publicaciones académicas</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Publicaciones académicas</h2>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<AcademicPeriodCombo
		messages={academicPeriodsData.messages}
		bind:selectedValue={selectedAcademicPeriod}
		on:message={fetchOnAcademicPeriodChange}
	></AcademicPeriodCombo>

	<AddModal
		formComponent={AddForm}
		modalTitle="Crear publicación académica"
		formData={addAcademicProductionForm}
		{comboMessages}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full place-content-center px-8">
	{#await filterAcademicProductionListPromise}
		cargando...
	{:then filterAcademicProductionListsByAcademicPeriodResponse}
		{#if filterAcademicProductionListsByAcademicPeriodResponse.academic_production_lists}
			<Table
				bind:filterAcademicProductionListRequest
				{filterAcademicProductionListsByAcademicPeriodResponse}
				bind:popoverFilterDataMap
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
				on:filterChanged={handleOnFilterChanged}
				on:detailedFilter={handleOnDetailedFilter}
				on:paginationChanged={handlePaginationChanged}
			></Table>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay artículos registrados.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
