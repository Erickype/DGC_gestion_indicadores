<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { readable } from 'svelte/store';

	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import Table from '$lib/components/table/tablePaginated.svelte';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import { createEventDispatcher } from 'svelte';
	import { toast } from 'svelte-sonner';



	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type { UpdateAuthorPersonSchema } from './schema';
	import UpdateForm from './UpdateForm.svelte';
	import {
		generateInitialFilterValue,
		newFilterAuthorsRequest
	} from '$lib/components/filters/authors/authors';
	import type { AuthorPerson, FilterAuthorsRequest, FilterAuthorsResponse } from '$lib/api/model/api/academicProduction/authors/authorsFilter';

	export let filterAuthorsResponse: FilterAuthorsResponse;
	let people: AuthorPerson[] = filterAuthorsResponse.authors;

	export let formData: SuperValidated<Infer<UpdateAuthorPersonSchema>>;
	let person: AuthorPerson;

	let filterValue = '';
	let pageIndex: number = 0;
	let pageSize: number = 0;
	export let filterAuthorsRequest: FilterAuthorsRequest = newFilterAuthorsRequest(
		pageSize,
		pageIndex
	);
	let initialFilterValue: string | undefined = generateInitialFilterValue(filterAuthorsRequest);

	export let popoverFilterDataMap: PopoverFilterDataMap = new Map();

	const filterFields = ['identity', 'name', 'lastname', 'email'];

	const table = createTable(readable(people), {
		page: addPagination({
			initialPageSize: filterAuthorsResponse.page_size,
			initialPageIndex: filterAuthorsResponse.page - 1,
			serverSide: true,
			serverItemCount: readable(filterAuthorsResponse.count)
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			initialFilterValue: initialFilterValue?.trim(),
			serverSide: true
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'identity',
			header: 'Identificación'
		}),
		table.column({
			accessor: 'name',
			header: 'Nombre'
		}),
		table.column({
			accessor: 'lastname',
			header: 'Apellido'
		}),
		table.column({
			accessor: 'email',
			header: 'Correo'
		}),
		table.column({
			accessor: ({ ID }) => ID,
			header: '',
			cell: ({ value }) => {
				const actions = createRender(DataTableActions, { id: value.toString() });
				actions.on('delete-confirmation', handleDeleteConfirmation);
				actions.on('update-action', handleUpdateAction);
				return actions;
			}
		})
	]);

	const dispatch = createEventDispatcher();
	async function handleDeleteConfirmation(event: any) {
		const detail: { status: boolean; id: number } = event.detail;
		if (detail.status) {
			const res = await deletePerson(detail.id.toString());
			if (!res.ok) {
				return toast.error('Error eliminando el registro');
			}
			toast.success('Se eliminó el registro');
			return dispatch('deleted', {
				status: true
			});
		}
	}

	let updateFormOpen = false;
	function handleUpdateAction(event: any) {
		const detail: { status: boolean; id: string } = event.detail;
		if (detail.status) {
			person = people.find((person) => person.ID.toString() === detail.id)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		}
	}

	async function deletePerson(id: string) {
		const url = `/api/person/` + id;
		const response = await fetch(url, {
			method: 'DELETE'
		});
		return response;
	}

	function handleUpdated(event: any) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			dispatch('updated', {
				status: true
			});
		}
	}

	function handleOnFilterChanged() {
		dispatch('filterChanged', {
			filter: filterValue
		});
	}
	function handleOnDetailedFilter() {
		dispatch('detailedFilter');
	}
	function handleOnPageChanged() {
		filterAuthorsRequest.page = pageIndex + 1;
		dispatch('paginationChanged');
	}
	function handleOnPageSizeChanged() {
		filterAuthorsRequest.page_size = pageSize;
		filterAuthorsRequest.page = 1;
		dispatch('paginationChanged');
	}
</script>

<UpdateModal
	modalTitle="Actualizar información de persona"
	{formData}
	formComponent={UpdateForm}
	bind:updateEntity={person}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/>

<div class="w-full">
	<Table
		tableHeightClass="h-[55vh]"
		{table}
		{columns}
		serverItemCount={filterAuthorsResponse.count}
		{filterFields}
		bind:popoverFilterDataMap
		bind:filter_value={filterValue}
		bind:page_index={pageIndex}
		bind:page_size={pageSize}
		on:filterChanged={handleOnFilterChanged}
		on:detailedFilter={handleOnDetailedFilter}
		on:pageChanged={handleOnPageChanged}
		on:pageSizeChanged={handleOnPageSizeChanged}
	/>
</div>
