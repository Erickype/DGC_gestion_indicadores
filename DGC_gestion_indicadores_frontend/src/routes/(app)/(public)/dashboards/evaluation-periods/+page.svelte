<script lang="ts">
	import type { PageServerData } from './$types';

	import Icon from 'lucide-svelte/icons/circle-gauge';

	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';

	export let data: PageServerData;

	const evaluationPeriodsData = data.evaluationPeriodsData;
	let selectedEvaluationPeriod: number = evaluationPeriodsData.periods.at(0)!.ID;

	async function fetchGetIndicatorsByEvaluationPeriodID() {
		/* const url = `/api/indicators/academicPeriod/${selectedAcademicPeriod}`;
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
		return (indicatorsPromise = response.json()); */
	}
</script>

<svelte:head>
	<title>Dashboards Periodos Evaluación</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Dashboards Periodos Evaluación</h2>
	</div>

	<AcademicPeriodCombo
		messages={evaluationPeriodsData.messages}
		bind:selectedValue={selectedEvaluationPeriod}
		on:message={fetchGetIndicatorsByEvaluationPeriodID}
	></AcademicPeriodCombo>
</div>
