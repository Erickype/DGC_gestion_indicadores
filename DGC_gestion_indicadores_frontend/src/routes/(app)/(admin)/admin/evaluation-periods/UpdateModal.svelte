<script lang="ts">
	import type { EvaluationPeriod } from '$lib/api/model/view/evaluationPeriod';
	import * as Dialog from '$lib/components/ui/dialog';

	import UpdateForm from './UpdateForm.svelte';
	import { createEventDispatcher } from 'svelte';

	export let formData;

	export let open: boolean;
	export let evaluationPeriod: EvaluationPeriod;
	const dispatch = createEventDispatcher();

	function handleUpdated(event: any) {
		const data: { updated: boolean } = event.detail;
		if (data.updated) {
			open = false;
		}
		dispatch('updated', {
			status: true
		});
	}
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Actualizar periodo de evaluación</Dialog.Title>
			<Dialog.Description>Actualizar la información luego click en Guardar.</Dialog.Description>
		</Dialog.Header>
		<UpdateForm data={formData} {evaluationPeriod} on:updated={handleUpdated} />
	</Dialog.Content>
</Dialog.Root>
