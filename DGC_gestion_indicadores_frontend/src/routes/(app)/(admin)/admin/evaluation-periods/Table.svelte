<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { readable } from 'svelte/store';

	import DataTableActions from '$lib/components/table/tableActions.svelte';

	import Table from '$lib/components/table/table.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';

	import type { EvaluationPeriod } from '$lib/api/model/view/evaluationPeriod';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdateEvaluationPeriodSchema } from './schema';
	import { createEventDispatcher } from 'svelte';
	import { toast } from 'svelte-sonner';
	
	import UpdateForm from './UpdateForm.svelte';

	export let periods: EvaluationPeriod[];
	export let formData: SuperValidated<Infer<UpdateEvaluationPeriodSchema>>;
	let evaluationPeriod: EvaluationPeriod;

	const filterFields = ['name', 'abbreviation'];

	const table = createTable(readable(periods), {
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
			accessor: 'start_year',
			header: 'Fecha inicio',
			cell: ({ value }) => {
				const date = new Date(value);
				const formattedDate = date.toLocaleDateString('es-EC', {
					year: 'numeric',
					month: 'numeric',
					day: 'numeric',
					timeZone: 'UTC'
				});
				return formattedDate;
			}
		}),
		table.column({
			accessor: 'end_year',
			header: 'Fecha fin',
			cell: ({ value }) => {
				const date = new Date(value);
				date.setHours(1);
				const formattedDate = date.toLocaleDateString('es-EC', {
					year: 'numeric',
					month: 'numeric',
					day: 'numeric',
					timeZone: 'UTC'
				});
				return formattedDate;
			}
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
			const res = await deleteEvaluationPeriod(detail.id.toString());
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
			evaluationPeriod = periods.find((period) => period.ID === detail.id)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		}
	}

	async function deleteEvaluationPeriod(id: string) {
		const url = `/api/evaluationPeriod/` + id;
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
	modalTitle="Actualizar periodo de evaluación"
	{formData}
	formComponent={UpdateForm}
	bind:updateEntity={evaluationPeriod}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/>

<div class="w-full">
	<Table {table} {columns} {filterFields} itemCount={periods.length} />
</div>
