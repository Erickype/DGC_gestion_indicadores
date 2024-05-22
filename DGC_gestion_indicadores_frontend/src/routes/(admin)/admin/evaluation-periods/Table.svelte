<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { readable } from 'svelte/store';

	import DataTableActions from '$lib/components/table/tableActions.svelte';

	import Table from '$lib/components/table/table.svelte';

	import type { EvaluationPeriod } from '$lib/api/model/view/evaluationPeriod';
	import { createEventDispatcher } from 'svelte';
	import { toast } from 'svelte-sonner';

	export let periods: EvaluationPeriod[];

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

	function handleUpdateAction(event: any) {
		const detail: { status: boolean; id: number } = event.detail;
		console.log(detail);
	}

	async function deleteEvaluationPeriod(id: string) {
		const url = `/api/evaluationPeriod/` + id;
		const response = await fetch(url, {
			method: 'DELETE'
		});
		return response;
	}
</script>

<div class="w-full">
	<Table {table} {columns} {filterFields} />
</div>
