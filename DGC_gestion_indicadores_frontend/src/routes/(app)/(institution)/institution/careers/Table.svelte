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

	import type { UpdateCarrerSchema } from './schema';
	import type { Career } from '$lib/api/model/api/career';
	import type { Message } from '$lib/components/combobox/combobox';

	export let careers: Career[];
	export let comboMessages: Message[];
	export let formData: SuperValidated<Infer<UpdateCarrerSchema>>;
	let career: Career;

	const filterFields = ['name', 'abbreviation'];

	const table = createTable(readable(careers), {
		page: addPagination({
			initialPageSize: 4
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
			const res = await deleteCareer(detail.id.toString());
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
		const detail: { status: boolean; id: number } = event.detail;
		if (detail.status) {
			career = careers.find((career) => career.ID === detail.id)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		}
	}

	async function deleteCareer(id: string) {
		const url = `/api/career/` + id;
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
	modalTitle="Actualizar información carrera"
	{formData}
	formComponent={UpdateForm}
	bind:updateEntity={career}
	comboMessages={[comboMessages]}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/>

<div class="w-full">
	<Table {table} {columns} {filterFields} />
</div>
