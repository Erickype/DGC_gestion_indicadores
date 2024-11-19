<script lang="ts">
	import SuperDebug, { superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { filterCohortsAuxSchema } from '$lib/utils';

	import type { PageServerData } from './$types';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';
	import { goto } from '$app/navigation';

	import { Button } from '$lib/components/ui/button/index.js';
	import * as Form from '$lib/components/ui/form';

	import Refresh from 'lucide-svelte/icons/refresh-ccw';
	import Icon from 'lucide-svelte/icons/circle-gauge';

	import CohortsServer from '$lib/components/filters/indicatorsInformation/postgraduate/cohortsServer.svelte';
	import type { IndicatorsPostgraduateJoined } from '$lib/api/model/api/indicators/postgraduateCohortYears';
	import IndicatorCard from '$lib/components/indicators/IndicatorCard.svelte';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';
	import { toast } from 'svelte-sonner';

	export let data: PageServerData;
	const filterCohortsAuxForm = data.filterCohortsAuxForm;

	let indicatorsPromise: Promise<IndicatorsPostgraduateJoined[]>;

	const form = superForm(filterCohortsAuxForm, {
		validators: zodClient(filterCohortsAuxSchema)
	});

	const { form: formData } = form;

	let formDataCohortYear = writable($formData.year);
	formDataCohortYear.subscribe((value) => {
		$formData.year = value;
		indicatorsPromise = fetchGetIndicatorsByCohortYear();
	});

	async function fetchGetIndicatorsByCohortYear() {
		const url = `/api/indicators/postgraduate/${$formData.year}`;
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

	async function fetchGetCalculateIndicatorByCohortYear() {
		const url = `/api/indicators/postgraduate/calculate/${$formData.year}`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			toast.error(errorData.message);
			return fetchGetIndicatorsByCohortYear();
		}
		toast.success('Indicadores actualizados');
		return (indicatorsPromise = response.json());
	}
</script>

<svelte:head>
	<title>Dashboards posgrado</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Dashboards posgrado</h2>
	</div>
	<div class="flex items-center gap-2">
		<Button
			variant="outline"
			role="combobox"
			size="icon"
			on:click={fetchGetCalculateIndicatorByCohortYear}
		>
			<Refresh class="h-4 w-4" />
		</Button>
		<Form.Field {form} name="year" class="flex w-96 justify-between">
			<CohortsServer {formDataCohortYear} />
		</Form.Field>
	</div>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
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
<!-- {#if browser}
	<SuperDebug data={$formData} />
{/if} -->
