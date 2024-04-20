<script lang="ts">
	import type { ActionData, PageServerData } from './$types';
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import { Button } from '$lib/components/ui/button/index';

	import CirclePlus from 'lucide-svelte/icons/circle-plus';
	import CircleX from 'lucide-svelte/icons/circle-x';
	import AddTeacherForm from './addTeacherForm.svelte';
	import { onMount } from 'svelte';
	import TeachersTable from './teachersTable.svelte';

	export let data: PageServerData;

	export let actionData: ActionData;
	let form = actionData?.form;

	const academicPeriodsData = data.academicPeriodsData;
	const peopleData = data.peopleData;
	const careersData = data.careersData;
	const dedicationData = data.dedicationsData;
	const scaledGradesData = data.scaledGradesData;
	let teachersPromise = data.teachersByAcademicPeriod;

	let selectedAcademicPeriod: number;

	let addTeacherAction = false;

	let teacherHasBeenCreated = false;

	onMount(() => {
		selectedAcademicPeriod = academicPeriodsData.periods.at(
			academicPeriodsData.periods.length - 1
		)!.ID;
	});

	$: {
		if (teacherHasBeenCreated) {
			addTeacherAction = false;
			teacherHasBeenCreated = false;
		}
	}

	async function fetchTeachers() {
		const url = `/api/teacher/byAcademicPeriodID/${selectedAcademicPeriod}`;
		const response = await fetch(url, {
			method: 'GET',
			credentials: 'include'
		});
		if (response.ok) {
			teachersPromise = response.json();
		} else {
			console.error('Failed to fetch teachers:', response.status);
		}
	}

	async function updateTeachersTable() {
		fetchTeachers();
	}
</script>

<svelte:head>
	<title>Docentes</title>
</svelte:head>

<div class="flex items-center justify-between">
	<AcademicPeriodCombo
		messages={academicPeriodsData.messages}
		bind:selectedValue={selectedAcademicPeriod}
		on:message={updateTeachersTable}
	></AcademicPeriodCombo>

	{#if !addTeacherAction}
		<Button
			size="sm"
			class="h-8 gap-1"
			on:click={() => {
				addTeacherAction = !addTeacherAction;
			}}
		>
			<CirclePlus class="h-3.5 w-3.5" />
			<span class="sr-only sm:not-sr-only sm:whitespace-nowrap"> Agregar docente </span>
		</Button>
	{:else}
		<Button
			variant="secondary"
			size="sm"
			class="h-8 gap-1"
			on:click={() => {
				addTeacherAction = !addTeacherAction;
			}}
		>
			<CircleX class="h-3.5 w-3.5" />
			<span class="sr-only sm:not-sr-only sm:whitespace-nowrap"> Cancelar </span>
		</Button>
	{/if}
</div>

{#if addTeacherAction}
	<div class="min-h-1/3 bg-muted/30 max-w-full rounded-md p-6">
		<AddTeacherForm
			data={data.addTeacherForm}
			bind:academicPeriod={selectedAcademicPeriod}
			people={peopleData.messages}
			careers={careersData.messages}
			dedications={dedicationData.messages}
			scaledGrades={scaledGradesData.messages}
			bind:teacherHasBeenCreated
			on:teacher-created={updateTeachersTable}
		></AddTeacherForm>
	</div>
{/if}

<div class="container mx-auto">
	{#await teachersPromise}
		loading...
	{:then teachers}
		{#if teachers.length > 0}
			<TeachersTable {teachers}></TeachersTable>
		{:else}
			No users here.
		{/if}
	{/await}
</div>
