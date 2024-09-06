<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { readable } from 'svelte/store';

	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import UpdateForm from './updateTeacherForm.svelte';
	import Table from '$lib/components/table/tablePaginated.svelte';

	import { createEventDispatcher } from 'svelte';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdateTeacherSchema } from './scheme';

	import { generateInitialFilterValue } from '$lib/components/filters/indicatorsInformation/teachersLists';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type {
		FilterTeachersListsByAcademicPeriodRequest,
		FilterTeachersListsByAcademicPeriodResponse,
		TeachersListByAcademicPeriodJoined
	} from '$lib/api/model/api/indicatorsInformation/teachersLists';
	import type { Message } from '$lib/components/combobox/combobox';

	export let filterTeachersListsByAcademicPeriodResponse: FilterTeachersListsByAcademicPeriodResponse;
	let filterTeachersListsResponse: TeachersListByAcademicPeriodJoined[] =
		filterTeachersListsByAcademicPeriodResponse.teachers_lists;

	export let comboMessages: Message[][] | undefined = undefined;
	export let formData: SuperValidated<Infer<UpdateTeacherSchema>, App.Superforms.Message>;
	let teacher: TeachersListByAcademicPeriodJoined;

	let filterValue = '';
	let pageIndex: number = 0;
	let pageSize: number = 0;
	export let filterTeachersListsByAcademicPeriodRequest: FilterTeachersListsByAcademicPeriodRequest;
	let initialFilterValue: string | undefined = generateInitialFilterValue(
		filterTeachersListsByAcademicPeriodRequest
	);

	export let popoverFilterDataMap: PopoverFilterDataMap = new Map();

	const filterFields = ['teacher', 'career', 'dedication', 'scaled_grade', 'contract_type'];

	const table = createTable(readable(filterTeachersListsResponse), {
		page: addPagination({
			initialPageSize: filterTeachersListsByAcademicPeriodResponse.page_size,
			initialPageIndex: filterTeachersListsByAcademicPeriodResponse.page - 1,
			serverSide: true,
			serverItemCount: readable(filterTeachersListsByAcademicPeriodResponse.count)
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			initialFilterValue: initialFilterValue?.trim(),
			serverSide: true
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'teacher',
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
			accessor: ({ teacher_id }) => teacher_id,
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
		const detail: { status: boolean; id: string } = event.detail;
		if (detail.status) {
			const segments = detail.id.split('/');
			const teacher_id = parseInt(segments.at(segments.length - 1)!);
			teacher = filterTeachersListsResponse.find((teacher) => teacher.teacher_id === teacher_id)!;
			updateFormOpen = true;
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
	function handleOnFilterChanged() {
		dispatch('filterChanged', {
			filter: filterValue
		});
	}
	function handleOnDetailedFilter() {
		dispatch('detailedFilter');
	}
	function handleOnPageChanged() {
		filterTeachersListsByAcademicPeriodRequest.page = pageIndex + 1;
		dispatch('paginationChanged');
	}
	function handleOnPageSizeChanged() {
		filterTeachersListsByAcademicPeriodRequest.page_size = pageSize;
		filterTeachersListsByAcademicPeriodRequest.page = 1;
		dispatch('paginationChanged');
	}
</script>

<UpdateModal
	modalTitle="Actualizar información profesor"
	{formData}
	formComponent={UpdateForm}
	{comboMessages}
	bind:updateEntity={teacher}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/>

<div class="w-full">
	<Table
		tableHeightClass="h-[50vh]"
		{table}
		{columns}
		serverItemCount={filterTeachersListsByAcademicPeriodResponse.count}
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
