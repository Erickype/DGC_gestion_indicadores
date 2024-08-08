<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { readable } from 'svelte/store';

	import DataTableActions from '$lib/components/table/tableActions.svelte';

	import Table from '$lib/components/table/table.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';

	/* 	import UpdateForm from './UpdateForm.svelte';
	 */
	import { createEventDispatcher } from 'svelte';
	import { toast } from 'svelte-sonner';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdateTeacherSchema } from './scheme';

	import type { GetTeachersByAcademicPeriodResponse, Teacher } from '$lib/api/model/api/teacher';
	import type { CommonDeleteResponse } from '$lib/api/model/common';
	import { goto } from '$app/navigation';

	export let teachers: GetTeachersByAcademicPeriodResponse[];
	export let formData: SuperValidated<Infer<UpdateTeacherSchema>, App.Superforms.Message>;
	let teacher: GetTeachersByAcademicPeriodResponse;

	const filterFields = ['name', 'abbreviation'];

	const table = createTable(readable(teachers), {
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
			accessor: 'person',
			header: 'Nombre'
		}),
		table.column({
			accessor: 'career',
			header: 'Carrera'
		}),
		table.column({
			accessor: 'dedication',
			header: 'Dedicación'
		}),
		table.column({
			accessor: 'scaled_grade',
			header: 'Grado Escalafonado'
		}),
		table.column({
			accessor: 'contract_type',
			header: 'Tipo contrato'
		}),
		table.column({
			accessor: ({ person_id }) => person_id,
			header: '',
			cell: ({ value }) => {
				const actions = createRender(DataTableActions, {
					id: `${formData.data.academicPeriod}/${value}`,
					moreAction: true,
					moreActionURL: `/information/teachers/`
				});
				actions.on('delete-confirmation', handleDeleteConfirmation);
				actions.on('update-action', handleUpdateAction);
				return actions;
			}
		})
	]);

	const dispatch = createEventDispatcher();
	async function handleDeleteConfirmation(event: any) {
		/* const detail: { status: boolean; id: number } = event.detail;
		if (detail.status) {
			const response = await deleteFeacher(detail.id.toString());
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
		} */
	}

	let updateFormOpen = false;
	function handleUpdateAction(event: any) {
		const detail: { status: boolean; id: number } = event.detail;
		if (detail.status) {
			/* teacher = teachers.find((teacher) => teacher.ID === detail.id)!;
			updateFormOpen = true; */
		} else {
			updateFormOpen = false;
		}
	}

	async function deleteTeacher(id: string) {
		/* const url = `/api/teacher/` + id;
		const response = await fetch(url, {
			method: 'DELETE'
		});
		return response; */
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
	modalTitle="Actualizar información facultad"
	{formData}
	formComponent={UpdateForm}
	bind:updateEntity={teacher}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/> -->

<div class="max-h-[45%] w-full">
	<Table {table} {columns} {filterFields} tableHeightClass="h-[50vh]" />
</div>
