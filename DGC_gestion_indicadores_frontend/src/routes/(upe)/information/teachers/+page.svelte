<script lang="ts">
	import type { ActionData, PageServerData } from './$types';
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import { Button } from '$lib/components/ui/button/index';

	import CirclePlus from 'lucide-svelte/icons/circle-plus';
	import CircleX from 'lucide-svelte/icons/circle-x';
	import AddTeacherForm from './addTeacherForm.svelte';
	import { onMount } from 'svelte';

	export let data: PageServerData;

	export let actionData: ActionData;
	let form = actionData?.form;

	const academicPeriodsData = data.academicPeriodsData;
	const peopleData = data.peopleData;
	const careersData = data.careersData;
	const dedicationData = data.dedicationsData;
	const scaledGradesData = data.scaledGradesData;
	const teachersByAcademicPeriodID = data.teachersByAcademicPeriod;

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
</script>

<svelte:head>
	<title>Docentes</title>
</svelte:head>

<div class="flex items-center justify-between">
	<AcademicPeriodCombo
		messages={academicPeriodsData.messages}
		bind:selectedValue={selectedAcademicPeriod}
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
		></AddTeacherForm>
	</div>
{/if}

<div class="flex h-full w-full items-center justify-center space-x-4">
	<pre>{JSON.stringify(teachersByAcademicPeriodID)}</pre>
</div>
