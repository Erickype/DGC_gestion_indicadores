<script lang="ts">
	import type { CommonError } from '$lib/api/model/errors';
	import type { PageServerData } from './$types';
	import { goto } from '$app/navigation';

	import Icon from 'lucide-svelte/icons/circle-gauge';
	import Activity from 'lucide-svelte/icons/activity';

	import { Button } from '$lib/components/ui/button/index.js';
	import { Progress } from '$lib/components/ui/progress';
	import * as Card from '$lib/components/ui/card';

	import type { IndicatorEvaluationPeriodJoined } from '$lib/api/model/api/indicators/evaluationPeriod';
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';

	import IndicatorPie from '$lib/components/indicators/IndicatorPie.svelte';

	export let data: PageServerData;

	const evaluationPeriodsData = data.evaluationPeriodsData;
	let selectedEvaluationPeriod: number = evaluationPeriodsData.periods.at(0)!.ID;

	let indicatorsPromise: Promise<IndicatorEvaluationPeriodJoined> =
		fetchGetIndicatorsByTypeIDAndEvaluationPeriodID(26);

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
			throw errorData;
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
			fetchGetIndicatorsByTypeIDAndEvaluationPeriodID(26);
		}}
	></AcademicPeriodCombo>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	{#await indicatorsPromise}
		cargando...
	{:then indicator}
		<Card.Root class="bg-secondary/50">
			<Card.Header class="pb-2">
				<Card.Title class="flex items-center justify-between gap-2 text-xl">
					<p>{indicator.indicator_type}</p>
					<Button variant="ghost" size="icon">
						<Activity class="h-4 w-4" />
					</Button>
				</Card.Title>
				<Card.Description>Obetivo {indicator.target_value}</Card.Description>
			</Card.Header>
			<Card.Content>
				<IndicatorPie {indicator} isPercentaje={false} />
			</Card.Content>
		</Card.Root>
	{/await}
</div>
