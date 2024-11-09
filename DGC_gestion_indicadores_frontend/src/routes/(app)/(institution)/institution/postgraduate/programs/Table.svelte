<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdatePostgraduateProgramSchema } from './schema';

	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import Table from '$lib/components/table/tablePaginated.svelte';
	/* import UpdateForm from './UpdateForm.svelte'; */

	import { generateInitialFilterValue } from '$lib/components/filters/indicatorsInformation/postgraduate/programs';
	import type {
		PostgraduateProgram,
		FilterPostgraduateProgramsRequest,
		FilterPostGraduateProgramsResponse
	} from '$lib/api/model/api/indicatorsInformation/postgraduate';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type { Message } from '$lib/components/combobox/combobox';

	export let filterPostGraduateProgramsResponse: FilterPostGraduateProgramsResponse;
	let filterPostgraduateResponse: PostgraduateProgram[] =
		filterPostGraduateProgramsResponse.postgraduate_programs;

	export let comboMessages: Message[][] | undefined;
	comboMessages = [];
	export let formData: SuperValidated<
		Infer<UpdatePostgraduateProgramSchema>,
		App.Superforms.Message
	>;
	let postgraduateProgram: PostgraduateProgram;

	let filterValue = '';
	let pageIndex: number = 0;
	let pageSize: number = 0;
	export let filterPostgraduateProgramsRequest: FilterPostgraduateProgramsRequest;
	let initialFilterValue: string | undefined = generateInitialFilterValue(
		filterPostgraduateProgramsRequest
	);

	export let popoverFilterDataMap: PopoverFilterDataMap = new Map();

	const filterFields = ['name', 'start_year', 'end_year', 'is_active'];

	const table = createTable(readable(filterPostgraduateResponse), {
		page: addPagination({
			initialPageSize: filterPostGraduateProgramsResponse.page_size,
			initialPageIndex: filterPostGraduateProgramsResponse.page - 1,
			serverSide: true,
			serverItemCount: readable(filterPostGraduateProgramsResponse.count)
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			initialFilterValue: initialFilterValue?.trim(),
			serverSide: true
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'name',
			header: 'Nombre'
		}),
		table.column({
			accessor: 'start_year',
			header: 'Año inicio'
		}),
		table.column({
			accessor: 'end_year',
			header: 'Año fin'
		}),
		table.column({
			accessor: 'is_active',
			header: 'Estado',
			cell: ({ value }) => {
				if (value) {
					return 'Activo';
				}
				return 'Inactivo';
			}
		}),
		table.column({
			accessor: ({ ID }) => ID,
			header: '',
			cell: ({ value }) => {
				const actions = createRender(DataTableActions, {
					id: `${value}`,
					moreAction: true,
					moreActionURL: `/information/books/`
				});
				actions.on('delete-confirmation', handleDeleteConfirmation);
				actions.on('update-action', handleUpdateAction);
				return actions;
			}
		})
	]);

	const dispatch = createEventDispatcher();
	async function handleDeleteConfirmation(event: any) {}

	let updateFormOpen = false;
	function handleUpdateAction(event: any) {
		const detail: { status: boolean; id: string } = event.detail;
		if (detail.status) {
			postgraduateProgram = filterPostgraduateResponse.find(
				(postgraduateProgram) => postgraduateProgram.ID!.toString() === detail.id
			)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		}
	}

	async function deleteTeacher(id: string) {}

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
		filterPostgraduateProgramsRequest.page = pageIndex + 1;
		dispatch('paginationChanged');
	}
	function handleOnPageSizeChanged() {
		filterPostgraduateProgramsRequest.page_size = pageSize;
		filterPostgraduateProgramsRequest.page = 1;
		dispatch('paginationChanged');
	}
</script>

<!-- <UpdateModal
	modalTitle="Actualizar libro o capítulo libro"
	{formData}
	formComponent={UpdateForm}
	{comboMessages}
	bind:updateEntity={postgraduateProgram}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/> -->

<div class="w-full">
	<Table
		tableHeightClass="h-[50vh]"
		{table}
		{columns}
		serverItemCount={filterPostGraduateProgramsResponse.count}
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
