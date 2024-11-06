<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import {
		updateArtisticProductionListSchema,
		type UpdateArtisticProductionListSchema
	} from './schema';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { Input } from '$lib/components/ui/input/index.js';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { type ArtisticProductionListJoined } from '$lib/api/model/api/indicatorsInformation/artisticProductionLists';
	import AcademicPeriodsServer from '$lib/components/filters/academicPeriods/academicPeriodsServer.svelte';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: SuperValidated<
		Infer<UpdateArtisticProductionListSchema>,
		App.Superforms.Message
	>;
	export let comboMessages: Message[][];
	export let updateEntity!: ArtisticProductionListJoined;
	comboMessages = [];

	const dispatch = createEventDispatcher();

	function ArtisticProductionUpdated() {
		dispatch('updated', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(updateArtisticProductionListSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				ArtisticProductionUpdated();
				return toast.success(`Producc´ón artística actualizada`);
			}
			fillForm();
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	fillForm();

	let formDataAcademicPeriodID = writable($formData.academic_period_id);
	formDataAcademicPeriodID.subscribe((value) => ($formData.academic_period_id = value));

	function fillForm() {
		$formData.academic_period_id = updateEntity.academic_period_id;
		$formData.international_artistic_work = updateEntity.international_artistic_work;
		$formData.national_artistic_work = updateEntity.national_artistic_work;
		$formData.intellectual_property = updateEntity.intellectual_property;
	}
</script>

<form action="?/updateArtisticProductionList" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academic_period_id" class="w-full">
			<AcademicPeriodsServer {formDataAcademicPeriodID} formSelectWidth="w-[90%]" />
		</Form.Field>

		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="international_artistic_work">
				<Form.Control let:attrs>
					<Form.Label>Obras internacionales</Form.Label>
					<Input
						{...attrs}
						bind:value={$formData.international_artistic_work}
						type="number"
						min="0"
					/>
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="national_artistic_work">
				<Form.Control let:attrs>
					<Form.Label>Obras nacionales</Form.Label>
					<Input {...attrs} bind:value={$formData.national_artistic_work} type="number" min="0" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</div>
		<Form.Field {form} name="intellectual_property">
			<Form.Control let:attrs>
				<Form.Label>Propiedad intelectual</Form.Label>
				<Input {...attrs} bind:value={$formData.intellectual_property} type="number" min="0" />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
