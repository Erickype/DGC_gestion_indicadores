<script lang="ts">
	import { filterAcademicPeriodsAuxSchema } from '$lib/utils';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { superForm } from 'sveltekit-superforms';

	import type { PageServerData } from './$types';
	import { writable } from 'svelte/store';

	import * as Form from '$lib/components/ui/form';

	import Icon from 'lucide-svelte/icons/users-round';

	import AcademicPeriodsServer from '$lib/components/filters/academicPeriods/academicPeriodsServer.svelte';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import Alert from '$lib/components/alert/alert.svelte';
	import AddForm from './AddForm.svelte';
	import Table from './Table.svelte';

	import {
		fetchFilterSocialProjectLists,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		newFilterSocialProjectListsByAcademiPeriodRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/indicatorsInformation/socialProjectLists/socialProjectLists';
	import type {
		FilterSocialProjectListsByAcademicPeriodRequest,
		FilterSocialProjectListsByAcademicPeriodResponse
	} from '$lib/api/model/api/indicatorsInformation/socialProjectLists';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: PageServerData;
	const filterAcademicPeriodsAuxForm = data.filterAcademicPeriodsAuxForm;
	const addSocialProjectListForm = data.addSocialProjectListForm;
	const updateSocialProjectListForm = data.updateSocialProjectListForm;

	const comboMessages: Message[][] = [data.careersData.messages];

	const form = superForm(filterAcademicPeriodsAuxForm, {
		validators: zodClient(filterAcademicPeriodsAuxSchema)
	});

	const { form: formData } = form;

	$formData.academic_period_id = data.academicPeriodsData.periods.at(0)!.ID;

	let filterSocialProjectListsByAcademicPeriodRequest: FilterSocialProjectListsByAcademicPeriodRequest =
		newFilterSocialProjectListsByAcademiPeriodRequest(5, 1, $formData.academic_period_id);
	let filterSocialProjectListsPromise: Promise<FilterSocialProjectListsByAcademicPeriodResponse>;
	let popoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	let formDataAcademicPeriodID = writable($formData.academic_period_id);
	formDataAcademicPeriodID.subscribe((value) => {
		$formData.academic_period_id = value;
		addSocialProjectListForm.data.academic_period_id = $formData.academic_period_id;
		updateSocialProjectListForm.data.academic_period_id = $formData.academic_period_id;
		fetchFilterSocialProjecListsOnAcademicPeriodChange();
	});

	function fetchFilterSocialProjecListsOnAcademicPeriodChange() {
		filterSocialProjectListsByAcademicPeriodRequest.academic_period_id =
			$formData.academic_period_id;
		filterSocialProjectListsPromise = fetchFilterSocialProjectLists(
			filterSocialProjectListsByAcademicPeriodRequest
		);
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			filterSocialProjectListsPromise = fetchFilterSocialProjectLists(
				filterSocialProjectListsByAcademicPeriodRequest
			);
		}
	}

	function handleOnFilterChanged(event: CustomEvent) {
		const data: { filter: string } = event.detail;
		filterSocialProjectListsPromise = fetchOnFilterChanged(
			data.filter.trim(),
			filterSocialProjectListsByAcademicPeriodRequest,
			popoverFilterDataMap
		);
	}

	function handleOnDetailedFilter() {
		filterSocialProjectListsPromise = fetchOnDetailedFilter(
			filterSocialProjectListsByAcademicPeriodRequest,
			popoverFilterDataMap
		).then(({ request, response }) => {
			filterSocialProjectListsByAcademicPeriodRequest = request;
			return response;
		});
	}

	function handlePaginationChanged() {
		filterSocialProjectListsPromise = fetchFilterSocialProjectLists(
			filterSocialProjectListsByAcademicPeriodRequest
		);
	}
</script>

<svelte:head>
	<title>Proyectos Vinculación</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Proyectos Vinculación</h2>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<Form.Field {form} name="academic_period_id" class="w-1/3">
		<AcademicPeriodsServer {formDataAcademicPeriodID} />
	</Form.Field>

	<AddModal
		formComponent={AddForm}
		modalTitle="Crear proyecto vinculación"
		formData={addSocialProjectListForm}
		{comboMessages}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await filterSocialProjectListsPromise}
		<TableSkeleton tableHeightClass="h-[55vh]" />
	{:then filterSocialProjectListsByAcademicPeriodResponse}
		{#if filterSocialProjectListsByAcademicPeriodResponse.social_project_lists}
			<Table
				bind:filterSocialProjectListsByAcademicPeriodRequest
				formData={updateSocialProjectListForm}
				{filterSocialProjectListsByAcademicPeriodResponse}
				bind:popoverFilterDataMap
				{comboMessages}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
				on:filterChanged={handleOnFilterChanged}
				on:detailedFilter={handleOnDetailedFilter}
				on:paginationChanged={handlePaginationChanged}
			></Table>
		{:else}
			<Alert
				title="Sin registros"
				description={'Ups, no hay proyectos de vinvulación registrados.'}
			/>
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
