<script lang="ts">
	import type { PageServerData } from './$types';
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import { Button } from '$lib/components/ui/button/index';

	import CirclePlus from 'lucide-svelte/icons/circle-plus';
	import CircleX from 'lucide-svelte/icons/circle-x';
	import AddTeacherForm from './addTeacherForm.svelte';
	import { onMount } from 'svelte';
	import TeachersTable from './teachersTable.svelte';
	import { hasTeacherDeleted, updateTeacherAction } from '../../../../../stores';
	import UpdateTeacherForm from './updateTeacherForm.svelte';
	import type { GetTeachersByAcademicPeriodResponse } from '$lib/api/model/api/teacher';
	import { goto, invalidateAll } from '$app/navigation';
	import { toast } from 'svelte-sonner';

	export let data: PageServerData;

	const academicPeriodsData = data.academicPeriodsData;
	const peopleData = data.peopleData;
	const careersData = data.careersData;
	const dedicationData = data.dedicationsData;
	const scaledGradesData = data.scaledGradesData;
	let teachersPromise = data.teachersByAcademicPeriod;

	let selectedAcademicPeriod: number;

	let addTeacherAction = false;
	let updateTeacherActionValue: { status: boolean; teacherID: number };

	let selectedTeacherToUpdate: Promise<GetTeachersByAcademicPeriodResponse>;

	let teacherHasBeenCreated = false;

	let hasTeacherDeletedValue: boolean;

	const unsubscribe = hasTeacherDeleted.subscribe((value) => {
		hasTeacherDeletedValue = value;
	});

	const unsubscribeUpdateTeacher = updateTeacherAction.subscribe((value) => {
		updateTeacherActionValue = value;
	});

	$: {
		if (hasTeacherDeletedValue) {
			console.log('Refreshing after delete');
			fetchTeachers();
			hasTeacherDeleted.set(false);
		}
	}

	$: {
		if (updateTeacherActionValue.status) {
			const ID = updateTeacherActionValue.teacherID;
			selectedTeacherToUpdate = teachersPromise.then((teachers) => {
				const foundTeacher = teachers.find((teacher) => teacher.ID === ID);
				if (foundTeacher) {
					return Promise.resolve(foundTeacher);
				} else {
					return Promise.reject(new Error('Teacher not found'));
				}
			});
		}
	}

	onMount(() => {
		invalidateAll()
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
			return teachersPromise = response.json();
		} 
		
		if(response.status === 401){
			toast.warning("No está autenticado.")
			return goto("/login")
		}
	}

	async function updateTeachersTable() {
		updateTeacherAction.set({
			status: false,
			teacherID: -1
		});
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
				updateTeacherAction.set({
					status: false,
					teacherID: -1
				});
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
{:else if updateTeacherActionValue.status}
	{#await selectedTeacherToUpdate}
		loading...
	{:then teacher}
		<div class="min-h-1/3 bg-muted/30 max-w-full rounded-md p-6">
			<UpdateTeacherForm
				data={data.updateTeacherForm}
				people={peopleData.messages}
				careers={careersData.messages}
				dedications={dedicationData.messages}
				scaledGrades={scaledGradesData.messages}
				selectedTeacherToUpdate={teacher}
				on:teacher-updated={updateTeachersTable}
			></UpdateTeacherForm>
		</div>
	{/await}
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
