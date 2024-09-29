<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';
	import { goto } from '$app/navigation';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdateDegreeAndTeachersListsDegreeSchema } from './schema';

	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateForm from './UpdateDegreeAndTeachersListsDegreeForm.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import Table from '$lib/components/table/table.svelte';
	import { toast } from 'svelte-sonner';

	import type { GetTeachersListsDegreesJoinedResponse } from '$lib/api/model/api/indicatorsInformation/teachersListsDegree';
	import type { CommonDeleteResponse } from '$lib/api/model/common';
	import type { Message } from '$lib/components/combobox/combobox';

	export let degrees: GetTeachersListsDegreesJoinedResponse[];

	export let comboMessages: Message[][] | undefined = undefined;
	export let formData: SuperValidated<Infer<UpdateDegreeAndTeachersListsDegreeSchema>>;
	let degree: GetTeachersListsDegreesJoinedResponse;

	const filterFields = ['degree_level_name', 'name'];

	const table = createTable(readable(degrees), {
		page: addPagination({
			initialPageSize: 10
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			fn: ({ filterValue, value }) => value.toLowerCase().includes(filterValue.toLowerCase())
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'degree_level_name',
			header: 'Nivel título'
		}),
		table.column({
			accessor: 'name',
			header: 'Nombre'
		}),
		table.column({
			accessor: ({ teachers_degree_id }) => teachers_degree_id,
			header: '',
			cell: ({ value }) => {
				const actions = createRender(DataTableActions, {
					id: `${value}`
				});
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
			const response = await deleteFaculty(detail.id.toString());
			if (!response.ok) {
				const error: App.Error = await response.json();
				toast.error(error.message);
				if (response.status === 401) {
					throw goto('/login');
				}
				return;
			}
			const deleteResponse: CommonDeleteResponse = await response.json();
			toast.success(`Facultad eliminada: ${deleteResponse.id_deleted}`);
			return dispatch('deleted', {
				status: true
			});
		}
	}

	let updateFormOpen = false;
	function handleUpdateAction(event: any) {
		const detail: { status: boolean; id: string } = event.detail;

		if (detail.status) {
			degree = degrees.find((degree) => degree.teachers_degree_id.toString() === detail.id)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		}
	}

	async function deleteFaculty(id: string) {
		const url = `/api/faculty/` + id;
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
</script>

<UpdateModal
	modalTitle="Actualizar información título"
	{formData}
	formComponent={UpdateForm}
	{comboMessages}
	bind:updateEntity={degree}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/>

<div class="h-[60vh] w-full">
	<Table {table} {columns} {filterFields} itemCount={degrees.length} />
</div>
