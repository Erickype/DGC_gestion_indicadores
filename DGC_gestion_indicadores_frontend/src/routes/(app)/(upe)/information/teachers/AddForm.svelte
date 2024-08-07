<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher, tick } from 'svelte';
	import { browser } from '$app/environment';
	import { cn } from '$lib/utils.js';

	import ChevronsUpDown from 'lucide-svelte/icons/chevrons-up-down';
	import Check from 'lucide-svelte/icons/check';

	import { buttonVariants } from '$lib/components/ui/button/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import * as Command from '$lib/components/ui/command/index.js';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import { addTeacherSchema, type AddTeacherSchema } from './scheme';
	import type { Message } from '$lib/components/combobox/combobox';
	import type { Teacher } from '$lib/api/model/api/teacher';

	export let data: SuperValidated<Infer<AddTeacherSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	const peopleComboData = comboMessages.at(0)!;
	const careersComboData = comboMessages.at(1)!;
	const dedicationsComboData = comboMessages.at(2)!;
	const scaledGradesComboData = comboMessages.at(3)!;
	const contractTypesComboData = comboMessages.at(4)!;

	const dispatch = createEventDispatcher();

	function TeacherCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addTeacherSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const teacher = message.data as Teacher;
				TeacherCreated();
				return toast.success(`Docente creado: ${teacher.person_id}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	let openPerson = false;
	let openCareer = false;
	let openDedication = false;
	let openScaledGrade = false;
	let openContractType = false;

	function closeAndFocusTriggerPerson(triggerId: string) {
		openPerson = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
	function closeAndFocusTriggerCareer(triggerId: string) {
		openCareer = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
	function closeAndFocusTriggerDedication(triggerId: string) {
		openDedication = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
	function closeAndFocusTriggerScaledGrade(triggerId: string) {
		openScaledGrade = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
	function closeAndFocusTriggerContractType(triggerId: string) {
		openContractType = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
</script>

<form action="?/addTeacher" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academicPeriod">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicPeriod} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="person" class="flex flex-col">
			<Popover.Root bind:open={openPerson} let:ids>
				<Form.Control let:attrs>
					<Form.Label>Persona</Form.Label>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'w-full justify-between',
							!$formData.person && 'text-muted-foreground'
						)}
						role="combobox"
						{...attrs}
					>
						{peopleComboData.find((f) => f.value === $formData.person)?.label ??
							'Seleccionar persona'}
						<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
					</Popover.Trigger>
					<input hidden value={$formData.person} name={attrs.name} />
				</Form.Control>
				<Popover.Content class="w-full p-0">
					<Command.Root>
						<Command.Input autofocus placeholder="Buscar persona..." class="h-9" />
						<Command.Empty>No se encontró la cédula.</Command.Empty>
						<Command.Group>
							{#each peopleComboData as person}
								<Command.Item
									value={person.label}
									onSelect={() => {
										$formData.person = person.value;
										closeAndFocusTriggerPerson(ids.trigger);
									}}
								>
									{person.label}
									<Check
										class={cn(
											'ml-auto h-4 w-4',
											person.value !== $formData.person && 'text-transparent'
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
		<Form.Field {form} name="career" class="flex flex-col">
			<Popover.Root bind:open={openCareer} let:ids>
				<Form.Control let:attrs>
					<Form.Label>Carrera</Form.Label>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'w-full justify-between',
							!$formData.career && 'text-muted-foreground'
						)}
						role="combobox"
						{...attrs}
					>
						{careersComboData.find((f) => f.value === $formData.career)?.label ??
							'Seleccionar carrera'}
						<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
					</Popover.Trigger>
					<input hidden value={$formData.career} name={attrs.name} />
				</Form.Control>
				<Popover.Content class="w-full p-0">
					<Command.Root>
						<Command.Input autofocus placeholder="Buscar carrera..." class="h-9" />
						<Command.Empty>No se encontó la carrera.</Command.Empty>
						<Command.Group>
							{#each careersComboData as career}
								<Command.Item
									value={career.label}
									onSelect={() => {
										$formData.career = career.value;
										closeAndFocusTriggerCareer(ids.trigger);
									}}
								>
									{career.label}
									<Check
										class={cn(
											'ml-auto h-4 w-4',
											career.value !== $formData.career && 'text-transparent'
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
		<Form.Field {form} name="dedication" class="flex flex-col">
			<Popover.Root bind:open={openDedication} let:ids>
				<Form.Control let:attrs>
					<Form.Label>Dedicación</Form.Label>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'w-full justify-between',
							!$formData.dedication && 'text-muted-foreground'
						)}
						role="combobox"
						{...attrs}
					>
						{dedicationsComboData.find((f) => f.value === $formData.dedication)?.label ??
							'Seleccionar dedicación'}
						<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
					</Popover.Trigger>
					<input hidden value={$formData.dedication} name={attrs.name} />
				</Form.Control>
				<Popover.Content class="w-full p-0">
					<Command.Root>
						<Command.Input autofocus placeholder="Buscar dedicación..." class="h-9" />
						<Command.Empty>No se encontó la dedicación.</Command.Empty>
						<Command.Group>
							{#each dedicationsComboData as dedication}
								<Command.Item
									value={dedication.label}
									onSelect={() => {
										$formData.dedication = dedication.value;
										closeAndFocusTriggerDedication(ids.trigger);
									}}
								>
									{dedication.label}
									<Check
										class={cn(
											'ml-auto h-4 w-4',
											dedication.value !== $formData.dedication && 'text-transparent'
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
		<Form.Field {form} name="scaledGrade" class="flex flex-col">
			<Popover.Root bind:open={openScaledGrade} let:ids>
				<Form.Control let:attrs>
					<Form.Label>Grado Escalafonado</Form.Label>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'w-full justify-between',
							!$formData.scaledGrade && 'text-muted-foreground'
						)}
						role="combobox"
						{...attrs}
					>
						{scaledGradesComboData.find((f) => f.value === $formData.scaledGrade)?.label ??
							'Seleccionar grado escalafonado'}
						<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
					</Popover.Trigger>
					<input hidden value={$formData.scaledGrade} name={attrs.name} />
				</Form.Control>
				<Popover.Content class="w-[90%] p-0">
					<Command.Root>
						<Command.Input autofocus placeholder="Buscar grado escalafonado..." class="h-9" />
						<Command.Empty>No se encontó el grado escalafonado.</Command.Empty>
						<Command.Group>
							{#each scaledGradesComboData as scaledGrade}
								<Command.Item
									value={scaledGrade.label}
									onSelect={() => {
										$formData.scaledGrade = scaledGrade.value;
										closeAndFocusTriggerScaledGrade(ids.trigger);
									}}
								>
									{scaledGrade.label}
									<Check
										class={cn(
											'ml-auto h-4 w-4',
											scaledGrade.value !== $formData.scaledGrade && 'text-transparent'
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
		<Form.Field {form} name="contractType" class="flex flex-col">
			<Popover.Root bind:open={openContractType} let:ids>
				<Form.Control let:attrs>
					<Form.Label>Tipo contrato</Form.Label>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'w-full justify-between',
							!$formData.contractType && 'text-muted-foreground'
						)}
						role="combobox"
						{...attrs}
					>
						{contractTypesComboData.find((f) => f.value === $formData.contractType)?.label ??
							'Seleccionar grado escalafonado'}
						<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
					</Popover.Trigger>
					<input hidden value={$formData.contractType} name={attrs.name} />
				</Form.Control>
				<Popover.Content class="w-[90%] p-0">
					<Command.Root>
						<Command.Input autofocus placeholder="Buscar grado tipo contrato..." class="h-9" />
						<Command.Empty>No se encontó el tipo contrato.</Command.Empty>
						<Command.Group>
							{#each contractTypesComboData as contractType}
								<Command.Item
									value={contractType.label}
									onSelect={() => {
										$formData.contractType = contractType.value;
										closeAndFocusTriggerContractType(ids.trigger);
									}}
								>
									{contractType.label}
									<Check
										class={cn(
											'ml-auto h-4 w-4',
											contractType.value !== $formData.contractType && 'text-transparent'
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
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
