<script lang="ts">
	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import * as Command from '$lib/components/ui/command/index.js';
	import { buttonVariants } from '$lib/components/ui/button';
	import * as Form from '$lib/components/ui/form';
	import { cn } from '$lib/utils.js';

	import ChevronsUpDown from 'lucide-svelte/icons/chevrons-up-down';
	import Check from 'lucide-svelte/icons/check';
	import type { Message } from './combobox';
	import { writable } from 'svelte/store';
	import { tick } from 'svelte';

	export let comboData: Message[];
	export let openCombo = false;
	export let formDataID = writable();
	export let emptyLabel: string = 'Seleccionar';
	export let formLabel: string = 'Object';
	export let scrollAreaHeight: string = '';
	export let formSelectWidth = 'w-[90%]';

	function closeAndFocusTrigger(triggerId: string) {
		openCombo = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
</script>

<Popover.Root bind:open={openCombo} let:ids>
	<Form.Control let:attrs>
		<Form.Label>{formLabel}</Form.Label>
		<Popover.Trigger
			class={cn(
				buttonVariants({ variant: 'outline' }),
				'w-full justify-between',
				!$formDataID && 'text-muted-foreground'
			)}
			role="combobox"
			{...attrs}
		>
			{comboData.find((f) => f.value === $formDataID)?.label ?? emptyLabel}
			<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
		</Popover.Trigger>
		<input hidden value={$formDataID} name={attrs.name} />
	</Form.Control>
	<Popover.Content class="{formSelectWidth} p-0">
		<Command.Root loop>
			<Command.Input autofocus placeholder="Filtrar..." class="h-9" />
			<Command.Empty>Sin coincidencias!</Command.Empty>
			<Command.Group>
				<ScrollArea class="{scrollAreaHeight} w-full">
					{#each comboData as comboMessage}
						<Command.Item
							value={comboMessage.label}
							onSelect={() => {
								$formDataID = comboMessage.value;
								closeAndFocusTrigger(ids.trigger);
							}}
						>
							{comboMessage.label}
							<Check
								class={cn(
									'ml-auto h-4 w-4',
									comboMessage.value !== $formDataID && 'text-transparent'
								)}
							/>
						</Command.Item>
					{/each}
				</ScrollArea>
			</Command.Group>
		</Command.Root>
	</Popover.Content>
</Popover.Root>
