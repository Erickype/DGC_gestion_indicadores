<script lang="ts">
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';

	import Icon from 'lucide-svelte/icons/book-marked';
	import type { PageServerData } from './$types';

	export let data: PageServerData;

	const academicPeriodsData = data.academicPeriodsData;
	let selectedAcademicPeriod: number = academicPeriodsData.periods.at(0)!.ID;

	function fetchOnAcademicPeriodChange() {
		/* filterTeachersListsByAcademicPeriodRequest.academic_period_id = selectedAcademicPeriod;
		filterTeachersListsByAcademicPeriodPromise = fetchFilterTeachersLists(
			filterTeachersListsByAcademicPeriodRequest
		); */
	}
</script>

<svelte:head>
	<title>Publicaciones académicas</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Publicaciones académicas</h2>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<AcademicPeriodCombo
		messages={academicPeriodsData.messages}
		bind:selectedValue={selectedAcademicPeriod}
		on:message={fetchOnAcademicPeriodChange}
	></AcademicPeriodCombo>
</div>
