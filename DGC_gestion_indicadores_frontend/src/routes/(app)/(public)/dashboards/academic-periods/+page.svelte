<script lang="ts">
	import type { PageServerData } from './$types';
	import { goto } from '$app/navigation';

	import Icon from 'lucide-svelte/icons/circle-gauge';
	import Refresh from 'lucide-svelte/icons/refresh-ccw';

	import type { IndicatorAcademicPeriodJoined } from '$lib/api/model/api/indicators/academicPeriods';
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';

	import IndicatorCard from '$lib/components/indicators/IndicatorCard.svelte';
	import { toast } from 'svelte-sonner';

	export let data: PageServerData;

	const academicPeriodsData = data.academicPeriodsData;
	let selectedAcademicPeriod: number = academicPeriodsData.periods.at(0)!.ID;

	let indicatorsPromise: Promise<IndicatorAcademicPeriodJoined[]> =
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
			return toast.error(errorData.message);
		}
		return (indicatorsPromise = response.json());
	}

	async function fetchGetCalculateIndicatorsByAcademicPeriod() {
		const url = `/api/indicators/academicPeriod/calculate/${selectedAcademicPeriod}`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			toast.error(errorData.message);
			return FetchGetIndicatorsByAcademicPeriodID();
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
	<div class="flex items-center gap-2">
		<Button
			variant="outline"
			role="combobox"
			size="icon"
			on:click={fetchGetCalculateIndicatorsByAcademicPeriod}
		>
			<Refresh class="h-4 w-4" />
		</Button>

		<AcademicPeriodCombo
			messages={academicPeriodsData.messages}
			bind:selectedValue={selectedAcademicPeriod}
			on:message={FetchGetIndicatorsByAcademicPeriodID}
		></AcademicPeriodCombo>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center px-8">
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
