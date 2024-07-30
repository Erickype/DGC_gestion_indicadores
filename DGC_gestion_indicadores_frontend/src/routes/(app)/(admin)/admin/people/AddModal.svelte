<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button/index.js';
	import Plus from 'lucide-svelte/icons/plus';

	/* 	import AddForm from './AddForm.svelte';
	 */ import { createEventDispatcher } from 'svelte';
	import AddForm from './AddForm.svelte';

	export let formData;

	let open: boolean;
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
	<Dialog.Trigger>
		<Button variant="default" class="ml-auto">
			Nuevo <Plus class="ml-2 h-4 w-4" />
		</Button>
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Añadir datos de persona</Dialog.Title>
			<Dialog.Description>Resgistrar la información luego click en Guardar.</Dialog.Description>
		</Dialog.Header>
		<AddForm data={formData} on:message={handleCreated} />
	</Dialog.Content>
</Dialog.Root>
