<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addTeacherSchema, type AddTeacherSchema } from './scheme';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import * as Form from '$lib/components/ui/form';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import * as Command from '$lib/components/ui/command/index.js';
	import ChevronsUpDown from 'lucide-svelte/icons/chevrons-up-down';
	import { buttonVariants } from '$lib/components/ui/button/index.js';
	import { cn } from '$lib/utils.js';
	import type { Message } from '$lib/components/combobox/combobox';
	import Check from 'lucide-svelte/icons/check';
	import { createEventDispatcher, tick } from 'svelte';
	import { toast } from 'svelte-sonner';
	import { browser } from '$app/environment';
	import { page } from '$app/stores';

	export let data: SuperValidated<Infer<AddTeacherSchema>>;
	export let academicPeriod: number;
	export let people: Message[];
	export let careers: Message[];
	export let dedications: Message[];
	export let scaledGrades: Message[];

	export let teacherHasBeenCreated = false;

	const form = superForm(data, {
		validators: zodClient(addTeacherSchema),
		onUpdated: ({ form: f }) => {
			if (!f.valid) {
				toast.error('Por favor completa todos los campos.');
			}
			if ($message.success) {
				teacherHasBeenCreated = true;
				toast.success('Registro de docente creado.');
				dispatchTeacherCreated()
			} else {
				toast.error('Fallo creando docente.');
			}
		}
	});

	const dispatch = createEventDispatcher()

	function dispatchTeacherCreated() {
		dispatch("teacher-created",{
			status: true
		})
	}

	const { form: formData, message, enhance } = form;

	$: $formData.academicPeriod = academicPeriod;

	$: if ($page.status == 200 || $page.status >= 400) {
		$formData.academicPeriod = academicPeriod;
	}

	let open = false;
	let openCareer = false;
	let openDedication = false;
	let openScaledGrade = false;

	function closeAndFocusTrigger(triggerId: string) {
		open = false;
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
</script>

<form action="?/addTeacher" method="post" use:enhance>
	<div class="flex w-full items-center justify-between">
		<Form.Field {form} name="academicPeriod">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicPeriod} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="person" class="flex flex-col">
			<Popover.Root bind:open let:ids>
				<Form.Control let:attrs>
					<Form.Label>Persona</Form.Label>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'w-[200px] justify-between',
							!$formData.person && 'text-muted-foreground'
						)}
						role="combobox"
						{...attrs}
					>
						{people.find((f) => f.value === $formData.person)?.label ?? 'Seleccionar persona'}
						<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
					</Popover.Trigger>
					<input hidden value={$formData.person} name={attrs.name} />
				</Form.Control>
				<Popover.Content class="w-[200px] p-0">
					<Command.Root>
						<Command.Input autofocus placeholder="Buscar persona..." class="h-9" />
						<Command.Empty>No se encontró la cédula.</Command.Empty>
						<Command.Group>
							{#each people as person}
								<Command.Item
									value={person.label}
									onSelect={() => {
										$formData.person = person.value;
										closeAndFocusTrigger(ids.trigger);
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
			<Form.Description>Persona para crear un docente.</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="career" class="flex flex-col">
			<Popover.Root bind:open={openCareer} let:ids>
				<Form.Control let:attrs>
					<Form.Label>Carrera</Form.Label>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'w-[200px] justify-between',
							!$formData.career && 'text-muted-foreground'
						)}
						role="combobox"
						{...attrs}
					>
						{careers.find((f) => f.value === $formData.career)?.label ?? 'Seleccionar carrera'}
						<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
					</Popover.Trigger>
					<input hidden value={$formData.career} name={attrs.name} />
				</Form.Control>
				<Popover.Content class="w-[200px] p-0">
					<Command.Root>
						<Command.Input autofocus placeholder="Buscar carrera..." class="h-9" />
						<Command.Empty>No se encontó la carrera.</Command.Empty>
						<Command.Group>
							{#each careers as career}
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
			<Form.Description>Carrera asociada a la persona</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="dedication" class="flex flex-col">
			<Popover.Root bind:open={openDedication} let:ids>
				<Form.Control let:attrs>
					<Form.Label>Dedicación</Form.Label>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'w-[200px] justify-between',
							!$formData.dedication && 'text-muted-foreground'
						)}
						role="combobox"
						{...attrs}
					>
						{dedications.find((f) => f.value === $formData.dedication)?.label ??
							'Seleccionar dedicación'}
						<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
					</Popover.Trigger>
					<input hidden value={$formData.dedication} name={attrs.name} />
				</Form.Control>
				<Popover.Content class="w-[200px] p-0">
					<Command.Root>
						<Command.Input autofocus placeholder="Buscar dedicación..." class="h-9" />
						<Command.Empty>No se encontó la dedicación.</Command.Empty>
						<Command.Group>
							{#each dedications as dedication}
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
			<Form.Description>Dedicación asociada a la persona</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="scaledGrade" class="flex flex-col">
			<Popover.Root bind:open={openScaledGrade} let:ids>
				<Form.Control let:attrs>
					<Form.Label>Grado Escalafonado</Form.Label>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'w-[250px] justify-between',
							!$formData.scaledGrade && 'text-muted-foreground'
						)}
						role="combobox"
						{...attrs}
					>
						{scaledGrades.find((f) => f.value === $formData.scaledGrade)?.label ??
							'Seleccionar grado escalafonado'}
						<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
					</Popover.Trigger>
					<input hidden value={$formData.scaledGrade} name={attrs.name} />
				</Form.Control>
				<Popover.Content class="w-[250px] p-0">
					<Command.Root>
						<Command.Input autofocus placeholder="Buscar grado escalafonado..." class="h-9" />
						<Command.Empty>No se encontó el grado escalafonado.</Command.Empty>
						<Command.Group>
							{#each scaledGrades as scaledGrade}
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
			<Form.Description>Grado escalafonado asociado a la persona</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<div class="flex items-center justify-center">
		<Form.Button>Guardar</Form.Button>
	</div>
	{#if browser}
		<SuperDebug data={$formData} />
	{/if}
</form>
