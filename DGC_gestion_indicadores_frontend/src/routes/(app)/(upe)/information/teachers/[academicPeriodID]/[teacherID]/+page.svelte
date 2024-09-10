<script lang="ts">
	import type { PageServerData } from './$types';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';

	import AddDegreeAndTeachersListsDegreeForm from './AddDegreeAndTeachersListsDegreeForm.svelte';
	import AddNotAssignedDegreeForm from './AddNotAssignedDegreeForm.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';

	import { Button } from '$lib/components/ui/button/index';

	import MoveLeft from 'lucide-svelte/icons/move-left';
	import SquarePlus from 'lucide-svelte/icons/square-plus';

	import type { GetTeachersListsDegreesJoinedResponse } from '$lib/api/model/api/indicatorsInformation/teachersListsDegree';
	import type { Message } from '$lib/components/combobox/combobox';
	import type { CommonError } from '$lib/api/model/errors';
	import TeachersListsDegreesTable from './Table.svelte';
	import Alert from '$lib/components/alert/alert.svelte';

	export let data: PageServerData;
	const addDegreeAndTeachersListsDegreeForm = data.addDegreeAndTeachersListsDegreeForm;
	const updateDegreeAndTeachersListsDegreeForm = data.updateDegreeAndTeachersListsDegreeForm;
	const addDegreeNotAssignedForm = data.addDegreeNotAssignedForm;

	const comboMessages: Message[][] = [data.degreeLevelsData.messages];
	const comboMessagesNotAssigned: Message[][] = [data.notAssigedDegrees.messages]

	$: {
		addDegreeAndTeachersListsDegreeForm.data.academicPeriodID = data.academicPeriodID;
		addDegreeAndTeachersListsDegreeForm.data.teacherID = data.teacherID;
		addDegreeNotAssignedForm.data.academicPeriodID = data.academicPeriodID
		addDegreeNotAssignedForm.data.teacherID = data.teacherID
	}

	let teachersListsDegreesPromise: Promise<GetTeachersListsDegreesJoinedResponse[]> =
		fetchGetTeachersListsDegreesJoined();

	async function fetchGetTeachersListsDegreesJoined() {
		const url = `/api/indicatorsInformation/teachersLists/degrees/${data.academicPeriodID}/${data.teacherID}`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			throw errorData;
		}
		return (teachersListsDegreesPromise = response.json());
	}

	function returnToTeachers() {
		if (browser) {
			window.history.back();
		}
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			fetchGetTeachersListsDegreesJoined();
		}
	}
</script>

<svelte:head>
	<title>Títulos</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Button variant="ghost" on:click={returnToTeachers}>
			<MoveLeft class="h-4 w-4" />
		</Button>
		<h2 class="text-2xl font-bold">Títulos</h2>
	</div>
	<div class="flex justify-between gap-1">
		<AddModal
			formComponent={AddNotAssignedDegreeForm}
			modalTitle="Agregar título"
			buttonName="Agregar"
			buttonVariant="secondary"
			icon={SquarePlus}
			formData={addDegreeNotAssignedForm}
			comboMessages={comboMessagesNotAssigned}
			on:created={fetchOnSuccess}
		/>
		<AddModal
			formComponent={AddDegreeAndTeachersListsDegreeForm}
			modalTitle="Crear título"
			formData={addDegreeAndTeachersListsDegreeForm}
			{comboMessages}
			on:created={fetchOnSuccess}
		/>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center px-8">
	{#await teachersListsDegreesPromise}
		cargando...
	{:then degrees}
		{#if degrees.length > 0}
			<TeachersListsDegreesTable
				formData={updateDegreeAndTeachersListsDegreeForm}
				{degrees}
				{comboMessages}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
			></TeachersListsDegreesTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay títulos registrados.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
