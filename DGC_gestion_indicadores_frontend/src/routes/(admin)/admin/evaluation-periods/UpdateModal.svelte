<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';

	import UpdateForm from './UpdateForm.svelte';
	import { createEventDispatcher } from 'svelte';

	export let formData;

	export let open: boolean;
	const dispatch = createEventDispatcher();

	function handleCreated(event: any) {
		const data: { created: boolean } = event.detail;
		if (data.created) {
			open = false;
		}
		dispatch('created', {
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
		<UpdateForm data={formData} on:message={handleCreated} />
	</Dialog.Content>
</Dialog.Root>
