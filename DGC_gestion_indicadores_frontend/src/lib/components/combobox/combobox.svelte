<script lang="ts">
	import Check from 'lucide-svelte/icons/check';
	import ChevronsUpDown from 'lucide-svelte/icons/chevrons-up-down';
	import { createEventDispatcher, tick } from 'svelte';
	import * as Command from '$lib/components/ui/command/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { cn } from '$lib/utils.js';
	import type { Message } from './combobox';

	export let data: Message[];
	export let placeholder: string = 'Seleccionar...';
	export let emptyMessage: string = 'No encontrado.';
	export let selectedValue: number | undefined = undefined;
	export let buttonClass: string = 'w-[300px]';
	export let popoverClass: string = 'w-[300px]';

	let open = false;
	let label: string;

	const dispatch = createEventDispatcher()

	function dispatchSelected() {
		dispatch("message",{
			changed: true
		})
	}

	$: {
		if (selectedValue) {			
			label = data.find((f) => f.value === selectedValue)!.label;
		}else{
			label = "Seleccionar una opción"
		}
	}
	// We want to refocus the trigger button when the user selects
	// an item from the list so users can continue navigating the
	// rest of the form with the keyboard.
	function closeAndFocusTrigger(triggerId: string) {
		open = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
</script>

<Popover.Root bind:open let:ids>
	<Popover.Trigger asChild let:builder>
		<Button
			builders={[builder]}
			variant="outline"
			role="combobox"
			aria-expanded={open}
			class="{buttonClass} justify-between"
		>
			{label}
			<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
		</Button>
	</Popover.Trigger>
	<Popover.Content class="{popoverClass} p-0">
		<Command.Root>
			<Command.Input {placeholder} />
			<Command.Empty>{emptyMessage}</Command.Empty>
			<Command.Group>
				{#each data as message}
					<Command.Item
						value={message.label}
						onSelect={() => {
							selectedValue = message.value;
							closeAndFocusTrigger(ids.trigger);
							dispatchSelected()
						}}
					>
						<Check
							class={cn(
								'mr-2 h-4 w-4',
								message.value !== data.find((f) => f.value === selectedValue)?.value &&
									'text-transparent'
							)}
						/>
						{message.label}
					</Command.Item>
				{/each}
			</Command.Group>
		</Command.Root>
	</Popover.Content>
</Popover.Root>
