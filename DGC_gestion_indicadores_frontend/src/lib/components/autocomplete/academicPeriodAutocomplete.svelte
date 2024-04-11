<script lang="ts">
	import { GetAcademicPeriods } from '$lib/api/controller/view/academicPeriod';
	import type { AcademicPeriod } from '$lib/api/model/view/academicPeriod';
	import type { Message } from './select';
	import Autocomplete from './autocomplete.svelte';
	import { onMount } from 'svelte';

	export let selected: number;
	export let id: string;
	export let name: string;
	let messages: Message[] = [];
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
</script>

<Autocomplete {id} {name} {messages} bind:selected></Autocomplete>
