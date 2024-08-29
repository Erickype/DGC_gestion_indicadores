<script lang="ts">
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import Alert from '$lib/components/alert/alert.svelte';
	import type { PageData } from './$types';

	import type { FilterPeopleRequest, FilterPeopleResponse } from '$lib/api/model/api/person';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import PeopleTable from './Table.svelte';
	import AddForm from './AddForm.svelte';
	import {
		fetchFilterPeople,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		newFilterPeopleRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/people';

	export let data: PageData;
	const addPersonFormData = data.addPersonForm;
	const updatePersonFormData = data.updatePersonForm;

	let filterPeopleRequest: FilterPeopleRequest = newFilterPeopleRequest(5, 1);
	let filterPeopleResponsePromise: Promise<FilterPeopleResponse> =
		fetchFilterPeople(filterPeopleRequest);
	let popoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function fetchOnSuccess(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			filterPeopleResponsePromise = fetchFilterPeople(filterPeopleRequest);
		}
	}

	function handleOnFilterChanged(event: CustomEvent) {
		const data: { filter: string } = event.detail;
		filterPeopleResponsePromise = fetchOnFilterChanged(
			data.filter.trim(),
			filterPeopleRequest,
			popoverFilterDataMap
		);
	}

	function handleOnDetailedFilter() {
		filterPeopleResponsePromise = fetchOnDetailedFilter(
			filterPeopleRequest,
			popoverFilterDataMap
		).then(({ request, response }) => {
			filterPeopleRequest = request;
			return response;
		});
	}

	function handlePaginationChanged() {
		filterPeopleResponsePromise = fetchFilterPeople(filterPeopleRequest);
	}
</script>

<svelte:head>
	<title>Profesores</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Profesores</h2>
	<AddModal
		modalTitle="Seleccionar a una persona para profesor"
		formComponent={AddForm}
		formData={addPersonFormData}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await filterPeopleResponsePromise}
		cargando...
	{:then filterPeopleResponse}
		{#if filterPeopleResponse.people.length > 0}
			<PeopleTable
				bind:filterPeopleRequest
				formData={updatePersonFormData}
				{filterPeopleResponse}
				bind:popoverFilterDataMap
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
				on:filterChanged={handleOnFilterChanged}
				on:detailedFilter={handleOnDetailedFilter}
				on:paginationChanged={handlePaginationChanged}
			></PeopleTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay personas registradas.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={`${error.Detail}: ${error.Message}`} />
	{/await}
</div>
