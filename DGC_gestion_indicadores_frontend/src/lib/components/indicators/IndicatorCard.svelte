<script lang="ts">
	import type { CommonError } from '$lib/api/model/errors';
	import { goto } from '$app/navigation';

	import { Button } from '$lib/components/ui/button/index.js';
	import * as Card from '$lib/components/ui/card';

	import Activity from 'lucide-svelte/icons/activity';
	import Loader from 'lucide-svelte/icons/loader';

	import type { IndicatorAcademicPeriodJoined } from '$lib/api/model/api/indicators/academicPeriods';
	import IndicatorPie from '$lib/components/indicators/IndicatorPie.svelte';
	import type {
		IndicatorEvaluationPeriodJoined,
		IndicatorsEvaluationPeriod
	} from '$lib/api/model/api/indicators/evaluationPeriod';

	export let indicator: IndicatorEvaluationPeriodJoined | IndicatorAcademicPeriodJoined;

	let isAcademicPeriodIndicator: boolean;
	if ((indicator as IndicatorAcademicPeriodJoined).academic_period_id) {
		isAcademicPeriodIndicator = true;
	} else {
		isAcademicPeriodIndicator = false;
	}

	let calculateIndicatorPromise: Promise<IndicatorsEvaluationPeriod>;

	async function fetchGetCalculateIndicatorByTypeIDAndEvaluationPeriod() {
		const actualIndicator = indicator as IndicatorEvaluationPeriodJoined;
		const url = `/api/indicators/evaluationPeriod/calculate/${actualIndicator.evaluation_period_id}/${actualIndicator.indicator_type_id}`;
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
		return (calculateIndicatorPromise = response.json());
	}
</script>

{#if isAcademicPeriodIndicator}
	<Card.Root class="bg-secondary/50">
		<Card.Header class="pb-2">
			<Card.Title class="flex items-center justify-between gap-2 text-xl">
				<p>{indicator.indicator_type}</p>
				<Button variant="ghost" size="icon">
					<Activity class="h-4 w-4" />
				</Button>
			</Card.Title>
			<Card.Description>Objetivo {indicator.target_value}</Card.Description>
		</Card.Header>
		<Card.Content>
			<IndicatorPie {indicator} isPercentaje={false} />
		</Card.Content>
	</Card.Root>
{:else}
	<Card.Root class="bg-secondary/50">
		<Card.Header class="pb-2">
			<Card.Title class="flex items-center justify-between gap-2 text-xl">
				<p>{indicator.indicator_type}</p>
				{#await calculateIndicatorPromise}
					<Button variant="ghost" size="icon" disabled>
						<Loader class="h-4 w-4" />
					</Button>
				{:then _}
					<Button
						variant="ghost"
						size="icon"
						on:click={fetchGetCalculateIndicatorByTypeIDAndEvaluationPeriod}
					>
						<Activity class="h-4 w-4" />
					</Button>
				{/await}
			</Card.Title>
			<Card.Description>Objetivo {indicator.target_value}</Card.Description>
		</Card.Header>
		<Card.Content>
			{#await calculateIndicatorPromise}
				cargando...
			{:then calculateIndicator}
				{#if calculateIndicator}
					<p class="hidden">
						{(indicator.actual_value = calculateIndicator.actual_value)}
					</p>
				{/if}
				<IndicatorPie {indicator} isPercentaje={false} />
			{/await}
		</Card.Content>
	</Card.Root>
{/if}
