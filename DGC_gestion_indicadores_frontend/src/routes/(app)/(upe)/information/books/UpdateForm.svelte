<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import {
		updateBookOrChaptersProductionListSchema,
		type UpdateBookOrChaptersProductionListSchema
	} from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';
	import { cn } from '$lib/utils';

	import { buttonVariants } from '$lib/components/ui/button/index.js';
	import { Textarea } from '$lib/components/ui/textarea/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import { Switch } from '$lib/components/ui/switch/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import CalendarIcon from 'lucide-svelte/icons/calendar';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import type {
		BooksOrChaptersProductionList,
		BooksOrChaptersProductionListsByAcademicPeriodJoined
	} from '$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionLists';
	import DetailedFieldsServer from '$lib/components/filters/detailedFields/detailedFieldsServer.svelte';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import { DateFormatter, getLocalTimeZone, parseDate, today } from '@internationalized/date';
	import CalendarMY from '$lib/components/calendar/month-year.svelte';
	import type { Message } from '$lib/components/combobox/combobox';

	export let updateEntity!: BooksOrChaptersProductionListsByAcademicPeriodJoined;
	export let data: SuperValidated<
		Infer<UpdateBookOrChaptersProductionListSchema>,
		App.Superforms.Message
	>;
	export let comboMessages: Message[][];
	comboMessages = [];

	const dispatch = createEventDispatcher();

	function BooksOrChaptersProductionListCreated() {
		dispatch('updated', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(updateBookOrChaptersProductionListSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const response = message.data as BooksOrChaptersProductionList;
				BooksOrChaptersProductionListCreated();
				return toast.success(
					`${response.is_chapter ? 'Capítulo libro' : 'Libro'} actualizado: ${response.title}`
				);
			}
			fillForm();
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	fillForm();

	let formDataDetailedFieldID = writable($formData.detailed_field_id);
	formDataDetailedFieldID.subscribe((value) => ($formData.detailed_field_id = value));

	let placeholderStart = today(getLocalTimeZone()).set({ day: 1, month: 1 });

	const df = new DateFormatter('ec-EC', {
		dateStyle: 'long'
	});

	$: publicationDate = $formData.publication_date
		? parseDate($formData.publication_date)
		: undefined;

	function fillForm() {
		$formData.id = updateEntity.ID;
		$formData.doi = updateEntity.DOI;
		$formData.is_chapter = updateEntity.is_chapter;
		$formData.title = updateEntity.title;
		$formData.publication_date = toISO8601(updateEntity.publication_date);
		$formData.peer_reviewed = updateEntity.peer_reviewed;
		$formData.detailed_field_id = updateEntity.detailed_field_id;
	}

	function toISO8601(dateString: string): string {
		const date = new Date(dateString);
		const year = date.getUTCFullYear();
		const month = String(date.getUTCMonth() + 1).padStart(2, '0'); // getUTCMonth() is zero-based
		const day = String(date.getUTCDate()).padStart(2, '0');
		return `${year}-${month}-${day}`;
	}

	function manageDateChanged(event: any) {
		const data: { date: string } = event.detail;
		const date = data.date;
		$formData.publication_date = date;
	}
</script>

<form action="?/updateBooksOrChaptersProductionLists" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="id">
			<Form.Control let:attrs>
				<input hidden value={$formData.id} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="academic_period_id">
			<Form.Control let:attrs>
				<input hidden value={$formData.academic_period_id} name={attrs.name} />
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
		<Form.Field {form} name="title">
			<Form.Control let:attrs>
				<Form.Label>{$formData.is_chapter ? 'Título del capítulo' : 'Título del libro'}</Form.Label>
				<Textarea
					{...attrs}
					placeholder="Saberes emergentes de las artes urbanas..."
					class="resize-none"
					bind:value={$formData.title}
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field
				{form}
				name="is_chapter"
				class="flex flex-row items-center justify-between rounded-lg border p-4"
			>
				<Form.Control let:attrs>
					<div class="space-y-0.5">
						<Form.Label>Tipo de publicación</Form.Label>
						<Form.Description
							>{!$formData.is_chapter ? 'Libro' : 'Capítulo de libro'}</Form.Description
						>
					</div>
					<Switch {...attrs} aria-readonly includeInput bind:checked={$formData.is_chapter} />
				</Form.Control>
			</Form.Field>
			<Form.Field
				{form}
				name="peer_reviewed"
				class="flex flex-row items-center justify-between rounded-lg border p-4"
			>
				<Form.Control let:attrs>
					<div class="space-y-0.5">
						<Form.Label>Revisado por pares</Form.Label>
						<Form.Description>{$formData.peer_reviewed ? 'Sí' : 'No'}</Form.Description>
					</div>
					<Switch {...attrs} aria-readonly includeInput bind:checked={$formData.peer_reviewed} />
				</Form.Control>
			</Form.Field>
		</div>
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
		<Form.Field {form} name="detailed_field_id">
			<DetailedFieldsServer {formDataDetailedFieldID}></DetailedFieldsServer>
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	{#if browser}
		<SuperDebug data={$formData} />
	{/if}
</form>
