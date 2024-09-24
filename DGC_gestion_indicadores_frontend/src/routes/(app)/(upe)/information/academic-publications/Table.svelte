<script lang="ts">
	import type {
		AcademicProductionListByAcademicPeriodJoined,
		FilterAcademicProductionListsByAcademicPeriodRequest,
		FilterAcademicProductionListsByAcademicPeriodResponse
	} from '$lib/api/model/api/indicatorsInformation/academicProductionLists';

	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';

	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import Table from '$lib/components/table/tablePaginated.svelte';

	import { generateInitialFilterValue } from '$lib/components/filters/indicatorsInformation/teachersLists/teachersLists';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';

	export let filterAcademicProductionListsByAcademicPeriodResponse: FilterAcademicProductionListsByAcademicPeriodResponse;
	let filterAcademicProductionListsResponse: AcademicProductionListByAcademicPeriodJoined[] =
		filterAcademicProductionListsByAcademicPeriodResponse.academic_production_lists;

	let academicProduction: AcademicProductionListByAcademicPeriodJoined;

	let filterValue = '';
	let pageIndex: number = 0;
	let pageSize: number = 0;
	export let filterAcademicProductionListRequest: FilterAcademicProductionListsByAcademicPeriodRequest;
	let initialFilterValue: string | undefined = generateInitialFilterValue(
		filterAcademicProductionListRequest
	);

	export let popoverFilterDataMap: PopoverFilterDataMap = new Map();

	const filterFields = [
		'doi',
		'publication_date',
		'publication_name',
		'detailed_field',
		'science_magazine',
		'impact_coefficient',
		'career',
		'intercultural_component'
	];

	const table = createTable(readable(filterAcademicProductionListsResponse), {
		page: addPagination({
			initialPageSize: filterAcademicProductionListsByAcademicPeriodResponse.page_size,
			initialPageIndex: filterAcademicProductionListsByAcademicPeriodResponse.page - 1,
			serverSide: true,
			serverItemCount: readable(filterAcademicProductionListsByAcademicPeriodResponse.count)
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			initialFilterValue: initialFilterValue?.trim(),
			serverSide: true
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'publication_date',
			header: 'Fecha',
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
			accessor: 'publication_name',
			header: 'Nombre'
		}),
		table.column({
			accessor: 'detailed_field',
			header: 'Campo detallado'
		}),
		table.column({
			accessor: 'science_magazine',
			header: 'Revista'
		}),
		table.column({
			accessor: 'impact_coefficient',
			header: 'Coeficiente'
		}),
		table.column({
			accessor: 'intercultural_component',
			header: 'Componente',
			cell: ({ value }) => {
				if (value) {
					return 'Sí';
				}
				return 'No';
			}
		}),
		table.column({
			accessor: ({ doi }) => doi,
			header: '',
			cell: ({ value }) => {
				const actions = createRender(DataTableActions, {
					id: `${value}`,
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
	async function handleDeleteConfirmation(event: any) {}

	let updateFormOpen = false;
	function handleUpdateAction(event: any) {
		const detail: { status: boolean; id: string } = event.detail;
		if (detail.status) {
			academicProduction = filterAcademicProductionListsResponse.find(
				(academicProduction) => academicProduction.doi === detail.id
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
		filterAcademicProductionListRequest.page = pageIndex + 1;
		dispatch('paginationChanged');
	}
	function handleOnPageSizeChanged() {
		filterAcademicProductionListRequest.page_size = pageSize;
		filterAcademicProductionListRequest.page = 1;
		dispatch('paginationChanged');
	}
</script>

<div class="w-full">
	<Table
		tableHeightClass="h-[50vh]"
		{table}
		{columns}
		serverItemCount={filterAcademicProductionListsByAcademicPeriodResponse.count}
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
