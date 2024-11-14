<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdateScienceMagazineSchema } from './schema';

	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import Table from '$lib/components/table/tablePaginated.svelte';
	import UpdateForm from './UpdateForm.svelte';

	import { generateInitialFilterValue } from '$lib/components/filters/indicatorsInformation/postgraduate/programs';
	import type {
		ScienceMagazineJoined,
		FilterScienceMagazinesRequest,
		FilterScienceMagazinesResponse
	} from '$lib/api/model/api/academicProduction/scienceMagazines/scienceMagazine';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type { Message } from '$lib/components/combobox/combobox';

	export let filterScienceMagazinesResponse: FilterScienceMagazinesResponse;
	let filterPostgraduateResponse: ScienceMagazineJoined[] =
		filterScienceMagazinesResponse.science_magazines;

	export let comboMessages: Message[][] | undefined;
	export let formData: SuperValidated<Infer<UpdateScienceMagazineSchema>, App.Superforms.Message>;
	let scienceMagazineJoined: ScienceMagazineJoined;

	let filterValue = '';
	let pageIndex: number = 0;
	let pageSize: number = 0;
	export let filterScienceMagazinesRequest: FilterScienceMagazinesRequest;
	let initialFilterValue: string | undefined = generateInitialFilterValue(
		filterScienceMagazinesRequest
	);

	export let popoverFilterDataMap: PopoverFilterDataMap = new Map();

	const filterFields = [
		'academic_database',
		'science_magazine',
		'science_magazine_abbreviation',
		'science_magazine_description'
	];

	const table = createTable(readable(filterPostgraduateResponse), {
		page: addPagination({
			initialPageSize: filterScienceMagazinesResponse.page_size,
			initialPageIndex: filterScienceMagazinesResponse.page - 1,
			serverSide: true,
			serverItemCount: readable(filterScienceMagazinesResponse.count)
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			initialFilterValue: initialFilterValue?.trim(),
			serverSide: true
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'academic_database',
			header: 'Base de datos'
		}),
		table.column({
			accessor: 'science_magazine',
			header: 'Revista'
		}),
		table.column({
			accessor: 'science_magazine_abbreviation',
			header: 'Abbreviación'
		}),
		table.column({
			accessor: 'science_magazine_description',
			header: 'Descripción'
		}),
		table.column({
			accessor: ({ science_magazine_id }) => science_magazine_id,
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
	async function handleDeleteConfirmation(event: any) {}

	let updateFormOpen = false;
	function handleUpdateAction(event: any) {
		const detail: { status: boolean; id: string } = event.detail;
		if (detail.status) {
			scienceMagazineJoined = filterPostgraduateResponse.find(
				(scienceMagazineJoined) =>
					scienceMagazineJoined.science_magazine_id!.toString() === detail.id
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
		filterScienceMagazinesRequest.page = pageIndex + 1;
		dispatch('paginationChanged');
	}
	function handleOnPageSizeChanged() {
		filterScienceMagazinesRequest.page_size = pageSize;
		filterScienceMagazinesRequest.page = 1;
		dispatch('paginationChanged');
	}
</script>

<UpdateModal
	modalTitle="Actualizar programa posgrado"
	{formData}
	formComponent={UpdateForm}
	{comboMessages}
	bind:updateEntity={scienceMagazineJoined}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/>

<div class="w-full">
	<Table
		tableHeightClass="h-[58vh]"
		{table}
		{columns}
		serverItemCount={filterScienceMagazinesResponse.count}
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
