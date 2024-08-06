<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { readable } from 'svelte/store';

	import DataTableActions from '$lib/components/table/tableActions.svelte';

	import Table from '$lib/components/table/table.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';

	import UpdateForm from './UpdateForm.svelte';

	import { createEventDispatcher } from 'svelte';
	import { toast } from 'svelte-sonner';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdateFacultySchema } from './schema';

	import type { Faculty } from '$lib/api/model/api/faculty';
	import type { CommonDeleteResponse } from '$lib/api/model/common';
	import { goto } from '$app/navigation';

	export let faculties: Faculty[];
	export let formData: SuperValidated<Infer<UpdateFacultySchema>>;
	let faculty: Faculty;

	const filterFields = ['name', 'abbreviation'];

	const table = createTable(readable(faculties), {
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
			accessor: 'name',
			header: 'Nombre'
		}),
		table.column({
			accessor: 'abbreviation',
			header: 'Abreviación'
		}),
		table.column({
			accessor: 'description',
			header: 'Descripción'
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
		const detail: { status: boolean; id: number } = event.detail;
		if (detail.status) {
			faculty = faculties.find((faculty) => faculty.ID === detail.id)!;
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
	modalTitle="Actualizar información facultad"
	{formData}
	formComponent={UpdateForm}
	bind:updateEntity={faculty}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/>

<div class="w-full max-h-[50%]">
	<Table {table} {columns} {filterFields} />
</div>
