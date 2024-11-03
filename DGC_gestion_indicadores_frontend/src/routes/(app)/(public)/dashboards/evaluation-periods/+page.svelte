<script lang="ts">
	import type { CommonError } from '$lib/api/model/errors';
	import type { PageServerData } from './$types';
	import { goto } from '$app/navigation';

	import Icon from 'lucide-svelte/icons/circle-gauge';

	import type { IndicatorEvaluationPeriodJoined } from '$lib/api/model/api/indicators/evaluationPeriod';
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';

	import IndicatorCard from '$lib/components/indicators/IndicatorCard.svelte';
	import { toast } from 'svelte-sonner';

	export let data: PageServerData;

	const evaluationPeriodsData = data.evaluationPeriodsData;
	let selectedEvaluationPeriod: number = evaluationPeriodsData.periods.at(0)!.ID;

	let indicatorsPromise: Promise<IndicatorEvaluationPeriodJoined> =
		fetchGetIndicatorsByTypeIDAndEvaluationPeriodID(25);

	async function fetchGetIndicatorsByTypeIDAndEvaluationPeriodID(indicatorTypeID: number) {
		const url = `/api/indicators/evaluationPeriod/${selectedEvaluationPeriod + '/' + indicatorTypeID}`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			return toast.error(errorData.message);
		}
		return (indicatorsPromise = response.json());
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
		on:message={() => {
			fetchGetIndicatorsByTypeIDAndEvaluationPeriodID(25);
		}}
	></AcademicPeriodCombo>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	{#await indicatorsPromise}
		cargando...
	{:then indicator}
		<IndicatorCard {indicator} />
	{/await}
</div>
