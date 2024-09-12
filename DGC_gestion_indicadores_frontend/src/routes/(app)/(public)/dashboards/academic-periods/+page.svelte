<script lang="ts">
	import Icon from 'lucide-svelte/icons/circle-gauge';
	import Activity from 'lucide-svelte/icons/activity';

	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import type { PageServerData } from './$types';
	import type { IndicatorsAcademicPeriod } from '$lib/api/model/api/indicators/academicPeriods';
	import type { CommonError } from '$lib/api/model/errors';
	import { goto } from '$app/navigation';
	import Alert from '$lib/components/alert/alert.svelte';

	import * as Card from '$lib/components/ui/card';
	import { Progress } from '$lib/components/ui/progress';

	export let data: PageServerData;

	const academicPeriodsData = data.academicPeriodsData;
	let selectedAcademicPeriod: number = academicPeriodsData.periods.at(0)!.ID;

	let indicatorsPromise: Promise<IndicatorsAcademicPeriod[]> =
		FetchGetIndicatorsByAcademicPeriodID();

	async function FetchGetIndicatorsByAcademicPeriodID() {
		const url = `/api/indicators/academicPeriod/${selectedAcademicPeriod}`;
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
	<title>Dashboards Periodos Académicos</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Dashboards Periodos Académicos</h2>
	</div>

	<AcademicPeriodCombo
		messages={academicPeriodsData.messages}
		bind:selectedValue={selectedAcademicPeriod}
		on:message={FetchGetIndicatorsByAcademicPeriodID}
	></AcademicPeriodCombo>
</div>

<div class="mx-auto flex h-[70vh] w-full place-content-center px-8">
	{#await indicatorsPromise}
		cargando...
	{:then indicators}
		{#if indicators.length > 0}
			<div class="my-auto grid min-h-40 w-full grid-cols-4 gap-6">
				{#each indicators as indicator}
					<Card.Root class="bg-secondary/50">
						<Card.Header class="pb-2">
							<Card.Description>Indicador {indicator.indicator_type_id}</Card.Description>
							<Card.Title class="flex justify-between text-4xl">
								{indicator.actual_value}%
								<Activity class="text-muted-foreground h-4 w-4" />
							</Card.Title>
						</Card.Header>
						<Card.Content>
							<div class="text-muted-foreground text-xs">Obetivo {indicator.target_value}%</div>
						</Card.Content>
						<Card.Footer>
							<Progress
								value={(indicator.actual_value * 100) / indicator.target_value}
								aria-label="25% increase"
							/>
						</Card.Footer>
					</Card.Root>
				{/each}
			</div>
		{:else}
			<Alert title="Sin registros" description={'Ups, aún no hay datos aquí.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
