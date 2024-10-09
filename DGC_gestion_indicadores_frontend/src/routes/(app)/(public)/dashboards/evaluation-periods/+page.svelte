<script lang="ts">
	import type { PageServerData } from './$types';

	export let data: PageServerData;

	const evaluationPeriodsData = data.evaluationPeriodsData;
	let selectedEvaluationPeriod: number = evaluationPeriodsData.periods.at(0)!.ID;
</script>

{selectedEvaluationPeriod}