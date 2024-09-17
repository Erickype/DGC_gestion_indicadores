<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addAcademicProductionSchema, type AddAcademicProductionSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';
	import { cn } from '$lib/utils';

	import { buttonVariants } from '$lib/components/ui/button/index.js';
	import { Textarea } from '$lib/components/ui/textarea/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import CalendarIcon from 'lucide-svelte/icons/calendar';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import TeachersServerFormSelect from '$lib/components/filters/teachers/teachersServer.svelte';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import CalendarMY from '$lib/components/calendar/month-year.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import type { Teacher } from '$lib/api/model/api/teacher';
	import {
		DateFormatter,
		type DateValue,
		getLocalTimeZone,
		parseDate,
		today
	} from '@internationalized/date';

	export let data: SuperValidated<Infer<AddAcademicProductionSchema>, App.Superforms.Message>;
	/* 	export let comboMessages: Message[][];
	const careersComboData = comboMessages.at(0)!;
	const dedicationsComboData = comboMessages.at(1)!;
	const scaledGradesComboData = comboMessages.at(2)!;
	const contractTypesComboData = comboMessages.at(3)!; */

	const dispatch = createEventDispatcher();

	function TeacherCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addAcademicProductionSchema),
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

	/*	let formDataCareerID = writable($formData.career);
	formDataCareerID.subscribe((value) => ($formData.career = value));*/

	let placeholderStart = today(getLocalTimeZone()).set({ day: 1, month: 1 });

	const df = new DateFormatter('ec-EC', {
		dateStyle: 'long'
	});

	$: publicationDate = $formData.publication_date
		? parseDate($formData.publication_date)
		: undefined;

	function manageDateChanged(event: any) {
		const data: { date: string } = event.detail;
		const date = data.date;
		$formData.publication_date = date;
	}
</script>

<form action="?/addTeacher" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academicPeriod">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicPeriod} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="doi">
			<Form.Control let:attrs>
				<Form.Label>DOI</Form.Label>
				<Input
					type="url"
					{...attrs}
					placeholder="https://doi.org/10.47366/sabia.v5n1a3..."
					bind:value={$formData.doi}
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="publication_name">
			<Form.Control let:attrs>
				<Form.Label>Nombre publicación</Form.Label>
				<Textarea
					{...attrs}
					placeholder="Saberes emergentes de las artes urbanas..."
					class="resize-none"
					bind:value={$formData.publication_name}
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="publication_date" class="flex flex-col">
			<Form.Control let:attrs>
				<Form.Label>Fecha inicio</Form.Label>
				<Popover.Root>
					<Popover.Trigger
						{...attrs}
						class={cn(
							buttonVariants({ variant: 'outline' }),
							'justify-start pl-4 text-left font-normal',
							!publicationDate && 'text-muted-foreground'
						)}
					>
						{publicationDate
							? df.format(publicationDate.toDate(getLocalTimeZone()))
							: 'Selecciona una fecha'}
						<CalendarIcon class="ml-auto h-4 w-4 opacity-50" />
					</Popover.Trigger>
					<Popover.Content class="w-auto p-0" side="top">
						<CalendarMY
							placeholder={placeholderStart}
							on:keydown={(v) => {
								manageDateChanged(v);
							}}
						/>
					</Popover.Content>
				</Popover.Root>
				<Form.FieldErrors />
				<input hidden value={$formData.publication_date} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<!-- 		<Form.Field {form} name="career" class="flex flex-col">
			<FormSelect
				formLabel="Carreras"
				comboData={careersComboData}
				bind:formDataID={formDataCareerID}
				scrollAreaHeight="h-72"
			/> -->
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
