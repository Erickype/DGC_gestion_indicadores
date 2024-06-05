<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addAcademicPeriodSchema, type AddAcademicPeriodSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { browser } from '$app/environment';
	import { cn } from '$lib/utils';
	import { toast } from 'svelte-sonner';
	import { createEventDispatcher } from 'svelte';

	import CalendarIcon from 'lucide-svelte/icons/calendar';

	import CalendarMY from '$lib/components/calendar/month-year.svelte';

	import {
		DateFormatter,
		type DateValue,
		getLocalTimeZone,
		parseDate,
		today
	} from '@internationalized/date';

	import { buttonVariants } from '$lib/components/ui/button/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';
	import type { CommonError } from '$lib/api/model/errors';

	export let data: SuperValidated<Infer<AddAcademicPeriodSchema>, App.Superforms.Message>;

	const dispatch = createEventDispatcher();

	function EvaluationPeriodCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addAcademicPeriodSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message!;
			if (message.success) {
				EvaluationPeriodCreated();
				return toast.success(`Periodo académico creado.`);
			}
			if (message.error as CommonError) {
				const commonError = message!.error as CommonError;
				return toast.error(`${commonError.message}: ${commonError.detail}`);
			} else {
				const error = message.error as string;
				return toast.error(error);
			}
		}
	});

	const { form: formData, message, enhance } = form;

	const df = new DateFormatter('ec-EC', {
		dateStyle: 'long'
	});

	let startDateValue: DateValue | undefined;
	let endDateValue: DateValue | undefined;

	$: startDateValue = $formData.startDate ? parseDate($formData.startDate) : undefined;
	$: endDateValue = $formData.endDate ? parseDate($formData.endDate) : undefined;

	type yearType = 'start' | 'end';
	function manageDateChanged(event: any, yearType: yearType) {
		const data: { date: string } = event.detail;
		const date = data.date;
		switch (yearType) {
			case 'start':
				$formData.startDate = date;
				break;
			case 'end':
				$formData.endDate = date;
				break;
		}
	}
</script>

<form action="?/addAcademicPeriod" use:enhance>
	<div class="flex flex-col gap-2">
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="name">
				<Form.Control let:attrs>
					<Form.Label>Nombre</Form.Label>
					<Input
						type="text"
						{...attrs}
						bind:value={$formData.name}
						placeholder="septiembre 2023 - febrero 2024"
					/>
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
						placeholder="sep 2023 - feb 2024"
					/>
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</div>
		<Form.Field {form} name="description">
			<Form.Control let:attrs>
				<Form.Label>Descripción</Form.Label>
				<Input
					type="text"
					{...attrs}
					bind:value={$formData.description}
					placeholder="Este periodo es muy importante"
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="startDate" class="flex flex-col">
				<Form.Control let:attrs>
					<Form.Label>Fecha inicio</Form.Label>
					<Popover.Root>
						<Popover.Trigger
							{...attrs}
							class={cn(
								buttonVariants({ variant: 'outline' }),
								'justify-start pl-4 text-left font-normal',
								!startDateValue && 'text-muted-foreground'
							)}
						>
							{startDateValue
								? df.format(startDateValue.toDate(getLocalTimeZone()))
								: 'Selecciona una fecha'}
							<CalendarIcon class="ml-auto h-4 w-4 opacity-50" />
						</Popover.Trigger>
						<Popover.Content class="w-auto p-0" side="top">
							<CalendarMY
								on:keydown={(v) => {
									manageDateChanged(v, 'start');
								}}
							/>
						</Popover.Content>
					</Popover.Root>
					<Form.FieldErrors />
					<input hidden value={$formData.startDate} name={attrs.name} />
				</Form.Control>
			</Form.Field>
			<Form.Field {form} name="endDate" class="flex flex-col">
				<Form.Control let:attrs>
					<Form.Label>Fecha fin</Form.Label>
					<Popover.Root>
						<Popover.Trigger
							{...attrs}
							class={cn(
								buttonVariants({ variant: 'outline' }),
								'justify-start pl-4 text-left font-normal',
								!endDateValue && 'text-muted-foreground'
							)}
						>
							{endDateValue
								? df.format(endDateValue.toDate(getLocalTimeZone()))
								: 'Selecciona una fecha'}
							<CalendarIcon class="ml-auto h-4 w-4 opacity-50" />
						</Popover.Trigger>
						<Popover.Content class="w-auto p-0" side="top">
							<CalendarMY
								on:keydown={(v) => {
									manageDateChanged(v, 'end');
								}}
							/>
						</Popover.Content>
					</Popover.Root>
					<Form.FieldErrors />
					<input hidden value={$formData.endDate} name={attrs.name} />
				</Form.Control>
			</Form.Field>
		</div>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
<!-- 	{#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
