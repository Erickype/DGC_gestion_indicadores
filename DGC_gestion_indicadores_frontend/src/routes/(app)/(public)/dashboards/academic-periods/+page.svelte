<script lang="ts">
	import Icon from 'lucide-svelte/icons/circle-gauge';

	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import type { PageServerData } from './$types';

	export let data: PageServerData;

	const academicPeriodsData = data.academicPeriodsData;
	let selectedAcademicPeriod: number = academicPeriodsData.periods.at(0)!.ID;
</script>

<svelte:head>
	<title>Dashboards Periodos Académicos</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Dashboards Periodos Académicos</h2>
	</div>

	<AcademicPeriodCombo
		messages={academicPeriodsData.messages}
		bind:selectedValue={selectedAcademicPeriod}
	></AcademicPeriodCombo>
</div>
