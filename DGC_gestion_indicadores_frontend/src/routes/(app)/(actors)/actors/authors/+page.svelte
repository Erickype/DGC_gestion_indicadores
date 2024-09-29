<script lang="ts">
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import Alert from '$lib/components/alert/alert.svelte';
	import type { PageData } from './$types';

	import Presentation from 'lucide-svelte/icons/square-user-round';

	import type {
		FilterAuthorsRequest,
		FilterAuthorsResponse
	} from '$lib/api/model/api/academicProduction/authors/authorsFilter';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import AddFromPersonForm from './AddFromPersonForm.svelte';
	import AuthorsTable from './Table.svelte';
	import AddForm from './AddForm.svelte';
	import {
		fetchFilterAuthors,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		newFilterAuthorsRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/authors/authors';

	export let data: PageData;
	const addAuthorFormData = data.addAuthorForm;
	const addAuthorFromPersonFormData = data.addAuthorFromPersonForm;
	const updateAuthorPersonFormData = data.updateAuthorPersonForm;

	let filterAuthorsRequest: FilterAuthorsRequest = newFilterAuthorsRequest(5, 1);
	let filterAuthorsResponsePromise: Promise<FilterAuthorsResponse> =
		fetchFilterAuthors(filterAuthorsRequest);
	let popoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function fetchOnSuccess(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			filterAuthorsResponsePromise = fetchFilterAuthors(filterAuthorsRequest);
		}
	}

	function handleOnFilterChanged(event: CustomEvent) {
		const data: { filter: string } = event.detail;
		filterAuthorsResponsePromise = fetchOnFilterChanged(
			data.filter.trim(),
			filterAuthorsRequest,
			popoverFilterDataMap
		);
	}

	function handleOnDetailedFilter() {
		filterAuthorsResponsePromise = fetchOnDetailedFilter(
			filterAuthorsRequest,
			popoverFilterDataMap
		).then(({ request, response }) => {
			filterAuthorsRequest = request;
			return response;
		});
	}

	function handlePaginationChanged() {
		filterAuthorsResponsePromise = fetchFilterAuthors(filterAuthorsRequest);
	}
</script>

<svelte:head>
	<title>Autores</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Presentation class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Autores</h2>
	</div>
	<div class="flex justify-between gap-1">
		<AddModal
			modalTitle="Seleccionar a una persona para autor"
			formComponent={AddForm}
			buttonName="Agregar"
			buttonVariant="secondary"
			formData={addAuthorFormData}
			comboMessages={undefined}
			on:created={fetchOnSuccess}
		/>
		<AddModal
			modalTitle="Crear a una persona para autor"
			formComponent={AddFromPersonForm}
			formData={addAuthorFromPersonFormData}
			comboMessages={undefined}
			on:created={fetchOnSuccess}
		/>
	</div>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await filterAuthorsResponsePromise}
		<TableSkeleton />
	{:then filterAuthorsResponse}
		{#if filterAuthorsResponse.authors.length > 0}
			<AuthorsTable
				bind:filterAuthorsRequest
				formData={updateAuthorPersonFormData}
				{filterAuthorsResponse}
				bind:popoverFilterDataMap
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
				on:filterChanged={handleOnFilterChanged}
				on:detailedFilter={handleOnDetailedFilter}
				on:paginationChanged={handlePaginationChanged}
			></AuthorsTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay personas registradas.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={`${error.Detail}: ${error.Message}`} />
	{/await}
</div>
