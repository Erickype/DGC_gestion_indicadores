<script lang="ts">
	import type { IndicatorsPostgraduateJoined } from '$lib/api/model/api/indicators/postgraduateCohortYears';
	import type { IndicatorEvaluationPeriodJoined } from '$lib/api/model/api/indicators/evaluationPeriod';
	import type { IndicatorAcademicPeriodJoined } from '$lib/api/model/api/indicators/academicPeriods';

	import { Chart, Svg, Group, LinearGradient, Arc, Text } from 'layerchart';

	export let indicator:
		| IndicatorAcademicPeriodJoined
		| IndicatorEvaluationPeriodJoined
		| IndicatorsPostgraduateJoined;
	export let isPercentaje: boolean = true;

	let value = (indicator.actual_value * 100) / indicator.target_value;
	if (indicator.actual_value > indicator.target_value) {
		value = 100;
	}
	let valueText = `${indicator.actual_value.toFixed(2)} ${isPercentaje ? '%' : ''}`;
</script>

<div class="h-[120px] rounded p-4">
	<Chart>
		<Svg center>
			<Group y={16}>
				<LinearGradient class="from-red-500 to-green-500" let:url>
					<Arc
						{value}
						range={[-120, 120]}
						outerRadius={60}
						innerRadius={50}
						cornerRadius={5}
						spring
						fill={url}
						track={{ class: 'fill-none stroke-surface-content/10' }}
					>
						<Text
							value={valueText}
							textAnchor="middle"
							verticalAnchor="middle"
							class="text-3xl tabular-nums"
						/>
					</Arc>
				</LinearGradient>
			</Group>
		</Svg>
	</Chart>
</div>
