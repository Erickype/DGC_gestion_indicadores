<script lang="ts">
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import Alert from '$lib/components/alert/alert.svelte';
	import AddForm from './AddForm.svelte';
	import Table from './Table.svelte';

	import Icon from 'lucide-svelte/icons/book-open-text';

	import type { PageServerData } from './$types';

	import {
		fetchFilterBooksOrChaptersProductionLists,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		newFilterBooksOrChaptersProductionListsByAcademiPeriodRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/indicatorsInformation/booksOrChaptersProductionLists/booksOrChaptersProductionLists';
	import type {
		FilterBooksOrChaptersProductionListsByAcademicPeriodRequest,
		FilterBooksOrChaptersProductionListsByAcademicPeriodResponse
	} from '$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionLists';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';

	export let data: PageServerData;
	const addBookOrChaptersProductionForm = data.addBookOrChaptersProductionForm;
	const updateBookOrChaptersProductionForm = data.updateBookOrChaptersProductionForm;

	const academicPeriodsData = data.academicPeriodsData;

	let selectedAcademicPeriod: number = academicPeriodsData.periods.at(0)!.ID;

	$: {
		addBookOrChaptersProductionForm.data.academic_period_id = selectedAcademicPeriod;
		updateBookOrChaptersProductionForm.data.academic_period_id = selectedAcademicPeriod;
	}

	let filterBooksOrChaptersProductionListsByAcademicPeriodRequest: FilterBooksOrChaptersProductionListsByAcademicPeriodRequest =
		newFilterBooksOrChaptersProductionListsByAcademiPeriodRequest(5, 1, selectedAcademicPeriod);
	let filterBooksOrChaptersProductionListPromise: Promise<FilterBooksOrChaptersProductionListsByAcademicPeriodResponse> =
		fetchFilterBooksOrChaptersProductionLists(
			filterBooksOrChaptersProductionListsByAcademicPeriodRequest
		);
	let popoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function fetchOnAcademicPeriodChange() {
		filterBooksOrChaptersProductionListsByAcademicPeriodRequest.academic_period_id =
			selectedAcademicPeriod;
		filterBooksOrChaptersProductionListPromise = fetchFilterBooksOrChaptersProductionLists(
			filterBooksOrChaptersProductionListsByAcademicPeriodRequest
		);
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			filterBooksOrChaptersProductionListPromise = fetchFilterBooksOrChaptersProductionLists(
				filterBooksOrChaptersProductionListsByAcademicPeriodRequest
			);
		}
	}

	function handleOnFilterChanged(event: CustomEvent) {
		const data: { filter: string } = event.detail;
		filterBooksOrChaptersProductionListPromise = fetchOnFilterChanged(
			data.filter.trim(),
			filterBooksOrChaptersProductionListsByAcademicPeriodRequest,
			popoverFilterDataMap
		);
	}

	function handleOnDetailedFilter() {
		filterBooksOrChaptersProductionListPromise = fetchOnDetailedFilter(
			filterBooksOrChaptersProductionListsByAcademicPeriodRequest,
			popoverFilterDataMap
		).then(({ request, response }) => {
			filterBooksOrChaptersProductionListsByAcademicPeriodRequest = request;
			return response;
		});
	}

	function handlePaginationChanged() {
		filterBooksOrChaptersProductionListPromise = fetchFilterBooksOrChaptersProductionLists(
			filterBooksOrChaptersProductionListsByAcademicPeriodRequest
		);
	}
</script>

<svelte:head>
	<title>Libros y Capítulos</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Libros y Capítulos</h2>
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
		modalTitle="Crear libro o capítulo"
		formData={addBookOrChaptersProductionForm}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await filterBooksOrChaptersProductionListPromise}
		<TableSkeleton />
	{:then filterBooksOrChaptersProductionListsByAcademicPeriodResponse}
		{#if filterBooksOrChaptersProductionListsByAcademicPeriodResponse.books_or_chapters_production_lists}
			<Table
				bind:filterBooksOrChaptersProductionListsByAcademicPeriodRequest
				formData={updateBookOrChaptersProductionForm}
				{filterBooksOrChaptersProductionListsByAcademicPeriodResponse}
				bind:popoverFilterDataMap
				comboMessages={undefined}
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
