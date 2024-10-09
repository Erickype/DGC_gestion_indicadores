<script lang="ts">
	import type { CommonError } from '$lib/api/model/errors';
	import type { PageServerData } from './$types';
	import { goto } from '$app/navigation';

	import Icon from 'lucide-svelte/icons/circle-gauge';
	import Activity from 'lucide-svelte/icons/activity';

	import { Progress } from '$lib/components/ui/progress';
	import * as Card from '$lib/components/ui/card';

	import type { IndicatorEvaluationPeriodJoined } from '$lib/api/model/api/indicators/evaluationPeriod';
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';

	export let data: PageServerData;

	const evaluationPeriodsData = data.evaluationPeriodsData;
	let selectedEvaluationPeriod: number = evaluationPeriodsData.periods.at(0)!.ID;

	let indicatorsPromise: Promise<IndicatorEvaluationPeriodJoined> = fetchGetIndicatorsByTypeIDAndEvaluationPeriodID(26);

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

{#await indicatorsPromise}
	cargando...
{:then indicator}
	<Card.Root class="bg-secondary/50">
		<Card.Header class="pb-2">
			<Card.Description>Indicador {indicator.indicator_type}</Card.Description>
			<Card.Title class="flex justify-between text-4xl">
				{indicator.actual_value.toFixed(2)}
				<Activity class="text-muted-foreground h-4 w-4" />
			</Card.Title>
		</Card.Header>
		<Card.Content>
			<div class="text-muted-foreground text-xs">
				Obetivo {indicator.target_value} per cápita
			</div>
		</Card.Content>
		<Card.Footer>
			<Progress
				value={(indicator.actual_value * 100) / indicator.target_value}
				aria-label="25% increase"
			/>
		</Card.Footer>
	</Card.Root>
{/await}
