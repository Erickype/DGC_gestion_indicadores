<script lang="ts">
	import { GetAcademicPeriods } from '$lib/api/controller/view/academicPeriod';
	import type { AcademicPeriod } from '$lib/api/model/view/academicPeriod';
	import type { Message } from './select';
	import Autocomplete from './autocomplete.svelte';
	import { onMount } from 'svelte';

	let messages: Message[] = [];
	export let selected: number;
	let periods: AcademicPeriod[] = [];

	onMount(async () => {
		const res = await GetAcademicPeriods();
		if (!res.ok) {
			throw new Error('Fallo cargando periodos');
		}
		periods = await res.json();
		messages = messages.concat(
			periods.map((period) => ({
				id: period.ID,
				name: period.name
			}))
		);
	});

	$: selectedPeriod = periods.find((period) => period.ID === selected);
</script>

<Autocomplete {messages} bind:selected></Autocomplete>
{#if selectedPeriod}
	<div class="flex flex-col text-sm text-accent-content p-2 m-2">
		<p>Selected id: {selectedPeriod?.ID}</p>
		<p>Selected name: {selectedPeriod?.name}</p>
	</div>
{/if}
