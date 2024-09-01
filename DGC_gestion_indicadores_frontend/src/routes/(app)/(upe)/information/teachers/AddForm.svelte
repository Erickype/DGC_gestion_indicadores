<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addTeacherSchema, type AddTeacherSchema } from './scheme';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import TeachersServerFormSelect from '$lib/components/filters/teachers/teachersServer.svelte';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import type { Teacher } from '$lib/api/model/api/teacher';

	export let data: SuperValidated<Infer<AddTeacherSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	const dedicationsComboData = comboMessages.at(0)!;
	const scaledGradesComboData = comboMessages.at(1)!;
	const contractTypesComboData = comboMessages.at(2)!;

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

	let formDataTeacherID = writable($formData.teacher);
	formDataTeacherID.subscribe((value) => ($formData.teacher = value));
	let formDataDedicationID = writable($formData.dedication);
	formDataDedicationID.subscribe((value) => ($formData.dedication = value));
	let formDataScaledGradeID = writable($formData.scaledGrade);
	formDataScaledGradeID.subscribe((value) => ($formData.scaledGrade = value));
	let formDataContractTypeID = writable($formData.contractType);
	formDataContractTypeID.subscribe((value) => ($formData.contractType = value));
</script>

<form action="?/addTeacher" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academicPeriod">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicPeriod} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="teacher" class="flex flex-col">
			<TeachersServerFormSelect bind:formDataTeacherID />
			<Form.FieldErrors />
		</Form.Field>
		<!-- <Form.Field {form} name="career" class="flex flex-col">
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
				<Popover.Content class="w-[90%] p-0">
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
		</Form.Field> -->
		<Form.Field {form} name="dedication" class="flex flex-col">
			<FormSelect
				formLabel="Dedicación"
				comboData={dedicationsComboData}
				bind:formDataID={formDataDedicationID}
			/>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="scaledGrade" class="flex flex-col">
			<FormSelect
				formLabel="Grado escalafonado"
				comboData={scaledGradesComboData}
				bind:formDataID={formDataScaledGradeID}
			/>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="contractType" class="flex flex-col">
			<FormSelect
				formLabel="Tipo contrato"
				comboData={contractTypesComboData}
				bind:formDataID={formDataContractTypeID}
			/>
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	{#if browser}
		<SuperDebug data={$formData} />
	{/if}
</form>
