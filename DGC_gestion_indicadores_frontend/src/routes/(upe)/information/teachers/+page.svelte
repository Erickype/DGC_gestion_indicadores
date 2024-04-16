<script lang="ts">
	import type { PageServerData } from './$types';
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import { Button } from '$lib/components/ui/button/index';

	import CirclePlus from 'lucide-svelte/icons/circle-plus';
	import AddTeacherForm from './addTeacherForm.svelte';
	import { onMount } from 'svelte';

	export let data: PageServerData;

	const academicPeriodsData = data.academicPeriodsData;
	const peopleData = data.peopleData;
	const careersData = data.careersData;

	let selectedAcademicPeriod: number;

	onMount(() => {
		selectedAcademicPeriod = academicPeriodsData.periods.at(
			academicPeriodsData.periods.length - 1
		)!.ID;
	});
</script>

<svelte:head>
	<title>Docentes</title>
</svelte:head>
<div class="flex items-center justify-between">
	<AcademicPeriodCombo
		messages={academicPeriodsData.messages}
		bind:selectedValue={selectedAcademicPeriod}
	></AcademicPeriodCombo>

	<Button size="sm" class="h-8 gap-1">
		<CirclePlus class="h-3.5 w-3.5" />
		<span class="sr-only sm:not-sr-only sm:whitespace-nowrap"> Agregar docente </span>
	</Button>
</div>

<div class="min-h-1/3 container max-w-full">
	<AddTeacherForm
		data={data.addTeacherForm}
		bind:academicPeriod={selectedAcademicPeriod}
		people={peopleData.messages}
		careers={careersData.messages}
	></AddTeacherForm>
</div>

<div class="flex h-full w-full items-center justify-center space-x-4"></div>
