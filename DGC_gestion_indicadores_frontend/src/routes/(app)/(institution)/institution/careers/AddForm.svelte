<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher, tick } from 'svelte';
	import { browser } from '$app/environment';
	import { toast } from 'svelte-sonner';

	import { addCarrerSchema, type AddCarrerSchema } from './schema';
	import { Input } from '$lib/components/ui/input';

	import { cn } from '$lib/utils.js';
	import * as Form from '$lib/components/ui/form';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import * as Command from '$lib/components/ui/command/index.js';
	import { buttonVariants } from '$lib/components/ui/button';

	import Check from 'lucide-svelte/icons/check';
	import ChevronsUpDown from 'lucide-svelte/icons/chevrons-up-down';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';

	import type { Message } from '$lib/components/combobox/combobox';
	import type { Career } from '$lib/api/model/api/career';
	
	export let data: SuperValidated<Infer<AddCarrerSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];

	const facultiesComboData = comboMessages.at(0) as Message[];
	let openFaculty = false;

	const dispatch = createEventDispatcher();

	function CareerCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addCarrerSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const career = message.data as Career;
				CareerCreated();
				return toast.success(`Carrera creada: ${career.abbreviation}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	function closeAndFocusTriggerFaculty(triggerId: string) {
		openFaculty = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
</script>

<form action="?/addCareer" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field class="flex flex-col" {form} name="facultyID">
			<Popover.Root bind:open={openFaculty} let:ids>
				<Form.Control let:attrs>
					<Form.Label>Facultad asignada</Form.Label>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'w-full justify-between',
							!$formData.facultyID && 'text-muted-foreground'
						)}
						role="combobox"
						{...attrs}
					>
						{facultiesComboData.find((f) => f.value === $formData.facultyID)?.label ??
							'Seleccionar facultad'}
						<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
					</Popover.Trigger>
					<input hidden value={$formData.facultyID} name={attrs.name} />
				</Form.Control>
				<Popover.Content class="w-[90%] p-0">
					<Command.Root>
						<Command.Input autofocus placeholder="Buscar facultad..." class="h-9" />
						<Command.Empty>Sin coincidencias!</Command.Empty>
						<Command.Group>
							{#each facultiesComboData as faculty}
								<Command.Item
									value={faculty.label}
									onSelect={() => {
										$formData.facultyID = faculty.value;
										closeAndFocusTriggerFaculty(ids.trigger);
									}}
								>
									{faculty.label}
									<Check
										class={cn(
											'ml-auto h-4 w-4',
											faculty.value !== $formData.facultyID && 'text-transparent'
										)}
									/>
								</Command.Item>
							{/each}
						</Command.Group>
					</Command.Root>
				</Popover.Content>
			</Popover.Root>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="name">
			<Form.Control let:attrs>
				<Form.Label>Nombre</Form.Label>
				<Input type="text" {...attrs} bind:value={$formData.name} placeholder="Carrera de ..." />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="abbreviation">
			<Form.Control let:attrs>
				<Form.Label>Abreviación</Form.Label>
				<Input
					type="text"
					{...attrs}
					bind:value={$formData.abbreviation}
					placeholder="Ing. en ..."
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="description">
			<Form.Control let:attrs>
				<Form.Label>Descripción</Form.Label>
				<Input
					type="text"
					{...attrs}
					bind:value={$formData.description}
					placeholder="Dedicada a ..."
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
