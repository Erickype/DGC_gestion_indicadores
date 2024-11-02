<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	/*import type { UpdateGradeRateListSchema } from './schema';*/
	import type {
		ResearchInnovationProjectList,
		ResearchInnovationProjectListJoined
	} from '$lib/api/model/api/indicatorsInformation/researchInnovationProjectLists';
	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import Table from '$lib/components/table/table.svelte';
	/*import UpdateForm from './UpdateForm.svelte';*/
	import { toast } from 'svelte-sonner';

	/*export let formData: SuperValidated<Infer<UpdateGradeRateListSchema>>;*/
	export let comboMessages: Message[][] | undefined = undefined;
	comboMessages = [];
	export let researchInnovationProjectLists: ResearchInnovationProjectListJoined[];
	let researchInnovationProjectList: ResearchInnovationProjectListJoined;

	const filterFields = [
		'academic_period',
		'total_projects_uep',
		'projects_external_funding',
		'international_cooperation_projects',
		'national_cooperation_projects'
	];

	const table = createTable(readable(researchInnovationProjectLists), {
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
			accessor: 'total_projects_uep',
			header: 'Total de UEP'
		}),
		table.column({
			accessor: 'projects_external_funding',
			header: 'Financiamiento externo'
		}),
		table.column({
			accessor: 'international_cooperation_projects',
			header: 'Internacional'
		}),
		table.column({
			accessor: 'national_cooperation_projects',
			header: 'Nacional'
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
		/* const detail: { status: boolean; id: string } = event.detail;
		if (detail.status) {
			researchInnovationProjectList = researchInnovationProjectLists.find(
				(researchInnovationProjectList) =>
					researchInnovationProjectList.career_id.toString() === detail.id &&
					researchInnovationProjectList.academic_period_id === formData.data.academic_period_id
			)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		} */
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
	bind:updateEntity={researchInnovationProjectList}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/> -->

<div class="w-full">
	<Table
		tableHeightClass={'h-[60vh]'}
		{table}
		{columns}
		{filterFields}
		itemCount={researchInnovationProjectLists.length}
	/>
</div>
