<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Dialog from '$lib/components/ui/dialog';
	import Plus from 'lucide-svelte/icons/plus';

	import type { Message } from '../combobox/combobox';
	import { createEventDispatcher } from 'svelte';
	import type { Variant } from '../ui/badge';

	export let formData;
	export let formComponent;
	export let modalTitle = 'Crear';
	export let buttonName = 'Nuevo';
	export let buttonVariant: Variant = 'default';
	export let icon = Plus;
	export let modalDescription = 'Resgistrar la información luego click en Guardar.';
	export let comboMessages: Message[][] | undefined = undefined;

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
		<Button variant={buttonVariant} class="ml-auto">
			{buttonName}<svelte:component this={icon} class="h-4 w-4" />
		</Button>
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>{modalTitle}</Dialog.Title>
			<Dialog.Description>{modalDescription}</Dialog.Description>
		</Dialog.Header>
		<svelte:component
			this={formComponent}
			data={formData}
			on:message={handleCreated}
			{comboMessages}
		></svelte:component>
	</Dialog.Content>
</Dialog.Root>
