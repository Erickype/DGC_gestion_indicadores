<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Card from '$lib/components/ui/card';

	import Activity from 'lucide-svelte/icons/activity';

	import type { IndicatorEvaluationPeriodJoined } from '$lib/api/model/api/indicators/evaluationPeriod';
	import type { IndicatorAcademicPeriodJoined } from '$lib/api/model/api/indicators/academicPeriods';
	import IndicatorPie from '$lib/components/indicators/IndicatorPie.svelte';

	export let indicator: IndicatorEvaluationPeriodJoined | IndicatorAcademicPeriodJoined;

	let isAcademicPeriodIndicator: boolean;
	if ((indicator as IndicatorAcademicPeriodJoined).academic_period_id) {
		isAcademicPeriodIndicator = true;
	} else {
		isAcademicPeriodIndicator = false;
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
			<Card.Description>Obetivo {indicator.target_value}</Card.Description>
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
{/if}
