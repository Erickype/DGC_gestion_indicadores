<script lang="ts">
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import Alert from '$lib/components/alert/alert.svelte';
	import type { PageData } from './$types';

	import Presentation from 'lucide-svelte/icons/presentation';

	import type { FilterTeachersRequest, FilterTeachersResponse } from '$lib/api/model/api/teacher';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import PeopleTable from './Table.svelte';
	import AddForm from './AddForm.svelte';
	import {
		fetchFilterTeachers,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		newFilterTeachersRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/teachers';

	export let data: PageData;
	const addPersonFormData = data.addPersonForm;
	const updatePersonFormData = data.updatePersonForm;

	let filterTeachersRequest: FilterTeachersRequest = newFilterTeachersRequest(5, 1);
	let filterTeachersResponsePromise: Promise<FilterTeachersResponse> =
		fetchFilterTeachers(filterTeachersRequest);
	let popoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function fetchOnSuccess(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			filterTeachersResponsePromise = fetchFilterTeachers(filterTeachersRequest);
		}
	}

	function handleOnFilterChanged(event: CustomEvent) {
		const data: { filter: string } = event.detail;
		filterTeachersResponsePromise = fetchOnFilterChanged(
			data.filter.trim(),
			filterTeachersRequest,
			popoverFilterDataMap
		);
	}

	function handleOnDetailedFilter() {
		filterTeachersResponsePromise = fetchOnDetailedFilter(
			filterTeachersRequest,
			popoverFilterDataMap
		).then(({ request, response }) => {
			filterTeachersRequest = request;
			return response;
		});
	}

	function handlePaginationChanged() {
		filterTeachersResponsePromise = fetchFilterTeachers(filterTeachersRequest);
	}
</script>

<svelte:head>
	<title>Profesores</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex gap-1 items-center">
		<Presentation class="h-8 w-8" />
	<h2 class="text-2xl font-bold">Profesores</h2>
	</div>
	<AddModal
		modalTitle="Seleccionar a una persona para profesor"
		formComponent={AddForm}
		formData={addPersonFormData}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await filterTeachersResponsePromise}
		cargando...
	{:then filterTeachersResponse}
		{#if filterTeachersResponse.teachers.length > 0}
			<PeopleTable
				bind:filterTeachersRequest
				formData={updatePersonFormData}
				{filterTeachersResponse}
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
