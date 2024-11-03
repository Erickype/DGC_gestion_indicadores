<script lang="ts">
	import type { CommonError } from '$lib/api/model/errors';
	import type { PageServerData } from './$types';
	import { goto } from '$app/navigation';

	import Refresh from 'lucide-svelte/icons/refresh-ccw';
	import Icon from 'lucide-svelte/icons/circle-gauge';

	import type { IndicatorEvaluationPeriodJoined } from '$lib/api/model/api/indicators/evaluationPeriod';
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';

	import IndicatorCard from '$lib/components/indicators/IndicatorCard.svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import Alert from '$lib/components/alert/alert.svelte';
	import { toast } from 'svelte-sonner';

	export let data: PageServerData;

	const evaluationPeriodsData = data.evaluationPeriodsData;
	let selectedEvaluationPeriod: number = evaluationPeriodsData.periods.at(0)!.ID;

	let indicatorsPromise: Promise<IndicatorEvaluationPeriodJoined[]> =
		fetchGetIndicatorsByEvaluationPeriod();

	async function fetchGetIndicatorsByEvaluationPeriod() {
		const url = `/api/indicators/evaluationPeriod/${selectedEvaluationPeriod}`;
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

	async function fetchGetCalculateIndicatorByEvaluationPeriod() {
		const url = `/api/indicators/evaluationPeriod/calculate/${selectedEvaluationPeriod}`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			toast.error(errorData.message);
			return fetchGetIndicatorsByEvaluationPeriod();
		}
		toast.success('Indicadores actualizados');
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
	<div class="flex items-center gap-2">
		<Button
			variant="outline"
			role="combobox"
			size="icon"
			on:click={fetchGetCalculateIndicatorByEvaluationPeriod}
		>
			<Refresh class="h-4 w-4" />
		</Button>

		<AcademicPeriodCombo
			messages={evaluationPeriodsData.messages}
			bind:selectedValue={selectedEvaluationPeriod}
			on:message={fetchGetIndicatorsByEvaluationPeriod}
		></AcademicPeriodCombo>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	{#await indicatorsPromise}
		cargando...
	{:then indicators}
		{#if indicators.length > 0}
			<div class="my-auto grid min-h-40 w-full grid-cols-3 gap-6">
				{#each indicators as indicator}
					<IndicatorCard {indicator} />
				{/each}
			</div>
		{:else}
			<Alert title="Sin registros" description={'Ups, aún no hay datos aquí.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
