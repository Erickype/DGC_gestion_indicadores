<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';

	import type { UpdateArtisticProductionListSchema } from './schema';
	import type { Infer, SuperValidated } from 'sveltekit-superforms';

	import type { ArtisticProductionListJoined } from '$lib/api/model/api/indicatorsInformation/artisticProductionLists';
	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import Table from '$lib/components/table/table.svelte';
	import UpdateForm from './UpdateForm.svelte';

	export let formData: SuperValidated<Infer<UpdateArtisticProductionListSchema>>;
	export let comboMessages: Message[][] | undefined = undefined;
	export let artisticProductionLists: ArtisticProductionListJoined[];
	let artisticProductionListJoined: ArtisticProductionListJoined;
	comboMessages = [];

	const filterFields = ['academic_period', 'international_artistic_work', 'national_artistic_work'];

	const table = createTable(readable(artisticProductionLists), {
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
			accessor: 'academic_period',
			header: 'Periodo académico'
		}),
		table.column({
			accessor: 'international_artistic_work',
			header: 'Obras internacionales'
		}),
		table.column({
			accessor: 'national_artistic_work',
			header: 'Obras nacionales'
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
	async function handleDeleteConfirmation(event: any) {}

	let updateFormOpen = false;
	function handleUpdateAction(event: any) {
		const detail: { status: boolean; id: string } = event.detail;
		if (detail.status) {
			artisticProductionListJoined = artisticProductionLists.find(
				(researchInnovationProjectList) =>
					researchInnovationProjectList.academic_period_id.toString() === detail.id
			)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		}
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
	modalTitle="Actualizar información producción artística"
	{formData}
	{comboMessages}
	formComponent={UpdateForm}
	bind:updateEntity={artisticProductionListJoined}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/>

<div class="w-full">
	<Table
		tableHeightClass={'h-[60vh]'}
		{table}
		{columns}
		{filterFields}
		itemCount={artisticProductionLists.length}
	/>
</div>
