<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdatePostgraduateCohortListSchema } from './schema';

	import type { PostgraduateCohortList } from '$lib/api/model/api/indicatorsInformation/postgraduate';
	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import Table from '$lib/components/table/table.svelte';
	/* import UpdateForm from './UpdateForm.svelte'; */
	import { toast } from 'svelte-sonner';

	export let formData: SuperValidated<Infer<UpdatePostgraduateCohortListSchema>>;
	export let comboMessages: Message[][] | undefined;
	comboMessages = [];
	export let postgraduateCohortLists: PostgraduateCohortList[];
	let postgraduateCohortList: PostgraduateCohortList;

	const filterFields = ['name', 'year', 'graduated_students', 'total_students'];

	const table = createTable(readable(postgraduateCohortLists), {
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
			accessor: 'name',
			header: 'Nombre'
		}),
		table.column({
			accessor: 'year',
			header: 'Año'
		}),
		table.column({
			accessor: 'graduated_students',
			header: 'Estudiantes graduados'
		}),
		table.column({
			accessor: 'total_students',
			header: 'Total estudiantes'
		}),
		table.column({
			accessor: ({ year }) => year,
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
			const res = await deletePostgraduateCohortList(detail.id.toString());
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
			postgraduateCohortList = postgraduateCohortLists.find(
				(postgraduateCohortList) =>
					postgraduateCohortList.year.toString() === detail.id &&
					postgraduateCohortList.postgraduate_program_id === formData.data.postgraduate_program_id
			)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		}
	}

	async function deletePostgraduateCohortList(id: string) {
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

<!-- <UpdateModal
	modalTitle="Actualizar información tasas de grado"
	{formData}
	{comboMessages}
	formComponent={UpdateForm}
	bind:updateEntity={postgraduateCohortList}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/> -->

<div class="w-full">
	<Table
		tableHeightClass={'h-[48vh]'}
		{table}
		{columns}
		{filterFields}
		itemCount={postgraduateCohortLists.length}
	/>
</div>
