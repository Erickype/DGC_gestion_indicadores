<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdateGradeRateListSchema } from './schema';

	import type { GradeRateListJoined } from '$lib/api/model/api/indicatorsInformation/gradeRateLists';
	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import Table from '$lib/components/table/table.svelte';
	/* import UpdateForm from './UpdateForm.svelte'; */
	import { toast } from 'svelte-sonner';


	export let formData: SuperValidated<Infer<UpdateGradeRateListSchema>>;
	export let comboMessages: Message[][] | undefined = undefined;
	export let gradeRateLists: GradeRateListJoined[];
	let academicPeriod: GradeRateListJoined;

	const filterFields = [
		'career',
		'count_graduated_students',
		'count_court_students',
		'count_admitted_matriculated_students',
		'count_admitted_students'
	];

	const table = createTable(readable(gradeRateLists), {
		page: addPagination({
			initialPageSize: 5
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			fn: ({ filterValue, value }) => value.toLowerCase().includes(filterValue.toLowerCase())
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'career',
			header: 'Carrera'
		}),
		table.column({
			accessor: 'count_graduated_students',
			header: 'Graduados'
		}),
		table.column({
			accessor: 'count_court_students',
			header: 'Cohorte'
		}),
		table.column({
			accessor: 'count_admitted_matriculated_students',
			header: 'Matriculados'
		}),
		table.column({
			accessor: 'count_admitted_students',
			header: 'Admitidos'
		}),
		table.column({
			accessor: ({ academic_period_id }) => academic_period_id,
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
			const res = await deleteAcademicPeriod(detail.id.toString());
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
		/* const detail: { status: boolean; id: string } = event.detail;
		if (detail.status) {
			academicPeriod = periods.find((period) => period.ID.toString() === detail.id)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		} */
	}

	async function deleteAcademicPeriod(id: string) {
		const url = `/api/academicPeriod/` + id;
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

<!-- 
<UpdateModal
	modalTitle="Actualizar información de periodo académico"
	{formData}
	formComponent={UpdateForm}
	bind:updateEntity={academicPeriod}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/> -->

<div class="w-full">
	<Table
		tableHeightClass={'h-[45vh]'}
		{table}
		{columns}
		{filterFields}
		itemCount={gradeRateLists.length}
	/>
</div>
