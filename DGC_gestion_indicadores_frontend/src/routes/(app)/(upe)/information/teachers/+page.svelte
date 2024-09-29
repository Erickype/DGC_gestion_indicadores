<script lang="ts">
	import type { PageServerData } from './$types';

	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import TeachersListsTable from './Table.svelte';
	import AddForm from './AddForm.svelte';

	import Icon from 'lucide-svelte/icons/list';

	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import Alert from '$lib/components/alert/alert.svelte';
	import type {
		FilterTeachersListsByAcademicPeriodRequest,
		FilterTeachersListsByAcademicPeriodResponse
	} from '$lib/api/model/api/indicatorsInformation/teachersLists';
	import {
		fetchFilterTeachersLists,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		newFilterTeachersListsByAcademiPeriodRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/indicatorsInformation/teachersLists/teachersLists';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';

	export let data: PageServerData;
	const addTeacherForm = data.addTeacherForm;
	const updateTeacherForm = data.updateTeacherForm;

	const academicPeriodsData = data.academicPeriodsData;

	const comboMessages: Message[][] = [
		data.careersData.messages,
		data.dedicationsData.messages,
		data.scaledGradesData.messages,
		data.contractTypesData.messages
	];

	let selectedAcademicPeriod: number = academicPeriodsData.periods.at(0)!.ID;

	$: {
		addTeacherForm.data.academicPeriod = selectedAcademicPeriod;
		updateTeacherForm.data.academicPeriod = selectedAcademicPeriod;
	}

	let filterTeachersListsByAcademicPeriodRequest: FilterTeachersListsByAcademicPeriodRequest =
		newFilterTeachersListsByAcademiPeriodRequest(5, 1, selectedAcademicPeriod);
	let filterTeachersListsByAcademicPeriodPromise: Promise<FilterTeachersListsByAcademicPeriodResponse> =
		fetchFilterTeachersLists(filterTeachersListsByAcademicPeriodRequest);
	let popoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function fetchOnAcademicPeriodChange() {
		filterTeachersListsByAcademicPeriodRequest.academic_period_id = selectedAcademicPeriod;
		filterTeachersListsByAcademicPeriodPromise = fetchFilterTeachersLists(
			filterTeachersListsByAcademicPeriodRequest
		);
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			filterTeachersListsByAcademicPeriodPromise = fetchFilterTeachersLists(
				filterTeachersListsByAcademicPeriodRequest
			);
		}
	}

	function handleOnFilterChanged(event: CustomEvent) {
		const data: { filter: string } = event.detail;
		filterTeachersListsByAcademicPeriodPromise = fetchOnFilterChanged(
			data.filter.trim(),
			filterTeachersListsByAcademicPeriodRequest,
			popoverFilterDataMap
		);
	}

	function handleOnDetailedFilter() {
		filterTeachersListsByAcademicPeriodPromise = fetchOnDetailedFilter(
			filterTeachersListsByAcademicPeriodRequest,
			popoverFilterDataMap
		).then(({ request, response }) => {
			filterTeachersListsByAcademicPeriodRequest = request;
			return response;
		});
	}

	function handlePaginationChanged() {
		filterTeachersListsByAcademicPeriodPromise = fetchFilterTeachersLists(
			filterTeachersListsByAcademicPeriodRequest
		);
	}
</script>

<svelte:head>
	<title>Listas Docentes</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Listas Docentes</h2>
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
		modalTitle="Crear docente"
		formData={addTeacherForm}
		{comboMessages}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await filterTeachersListsByAcademicPeriodPromise}
		<TableSkeleton />
	{:then filterTeachersListsByAcademicPeriodResponse}
		{#if filterTeachersListsByAcademicPeriodResponse.teachers_lists}
			<TeachersListsTable
				bind:filterTeachersListsByAcademicPeriodRequest
				formData={updateTeacherForm}
				{filterTeachersListsByAcademicPeriodResponse}
				bind:popoverFilterDataMap
				{comboMessages}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
				on:filterChanged={handleOnFilterChanged}
				on:detailedFilter={handleOnDetailedFilter}
				on:paginationChanged={handlePaginationChanged}
			></TeachersListsTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay docentes registrados.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
