<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addAcademicProductionSchema, type AddAcademicProductionSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { cn, toISO8601 } from '$lib/utils';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { buttonVariants } from '$lib/components/ui/button/index.js';
	import { Checkbox } from '$lib/components/ui/checkbox/index.js';
	import { Textarea } from '$lib/components/ui/textarea/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import CalendarIcon from 'lucide-svelte/icons/calendar';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import type { AcademicProductionList } from '$lib/api/model/api/indicatorsInformation/academicProductionLists';
	import ScienceMagazinesServer from '$lib/components/filters/scienceMagazines/scienceMagazinesServer.svelte';
	import DetailedFieldsServer from '$lib/components/filters/detailedFields/detailedFieldsServer.svelte';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import { DateFormatter, getLocalTimeZone, parseDate, today } from '@internationalized/date';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import CalendarMY from '$lib/components/calendar/month-year.svelte';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: SuperValidated<Infer<AddAcademicProductionSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	const impactCoefficientsComboData = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function AcademicProductionCreated() {
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
				const academicProduction = message.data as AcademicProductionList;
				AcademicProductionCreated();
				return toast.success(`Publicación creada: ${academicProduction.publication_name}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	let formDataScienceMagazineID = writable($formData.science_magazine_id);
	formDataScienceMagazineID.subscribe((value) => ($formData.science_magazine_id = value));
	let formImpactCoefficientID = writable($formData.impact_coefficient_id);
	formImpactCoefficientID.subscribe((value) => ($formData.impact_coefficient_id = value));
	let formDataDetailedFieldID = writable($formData.detailed_field_id);
	formDataDetailedFieldID.subscribe((value) => ($formData.detailed_field_id = value));

	let startYear = new Date($formData.startDate).getFullYear();
	let endYear = new Date($formData.endDate).getFullYear();
	let placeholderStart = parseDate(toISO8601(startYear.toString())).set({ day: 1, month: 1 });

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

<form action="?/postAcademicProductionList" use:enhance>
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
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="publication_date" class="flex flex-col">
				<Form.Control let:attrs>
					<Form.Label>Fecha publicación</Form.Label>
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
								{startYear}
								{endYear}
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
			<Form.Field class="flex flex-col" {form} name="science_magazine_id">
				<ScienceMagazinesServer {formDataScienceMagazineID} />
			</Form.Field>
		</div>
		<Form.Field {form} name="detailed_field_id">
			<DetailedFieldsServer {formDataDetailedFieldID}></DetailedFieldsServer>
		</Form.Field>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="impact_coefficient_id" class="flex flex-col">
				<FormSelect
					formLabel="Coefficiente impacto"
					formSelectWidth="w-[45%]"
					comboData={impactCoefficientsComboData}
					bind:formDataID={formImpactCoefficientID}
				/>
			</Form.Field>
			<Form.Field
				{form}
				name="intercultural_component"
				class="flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4"
			>
				<Form.Control let:attrs>
					<Checkbox {...attrs} bind:checked={$formData.intercultural_component} />
					<div class="space-y-1 leading-none">
						<Form.Label
							>{!$formData.intercultural_component
								? 'Sin componente intercultural'
								: 'Con componente intercultural'}</Form.Label
						>
					</div>
					<input name={attrs.name} value={$formData.intercultural_component} hidden />
				</Form.Control>
			</Form.Field>
		</div>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
