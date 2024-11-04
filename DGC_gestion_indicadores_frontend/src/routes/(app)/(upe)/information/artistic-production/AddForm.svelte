<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addArtisticProductionListSchema, type AddArtisticProductionListSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { Input } from '$lib/components/ui/input/index.js';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import AcademicPeriodsServer from '$lib/components/filters/academicPeriods/academicPeriodsServer.svelte';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: SuperValidated<Infer<AddArtisticProductionListSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	comboMessages = [];

	const dispatch = createEventDispatcher();

	function ArtisticProductionCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addArtisticProductionListSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				ArtisticProductionCreated();
				return toast.success(`Producción artística creada`);
			}
			formDataAcademicPeriodID.subscribe((value) => ($formData.academic_period_id = value));
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	let formDataAcademicPeriodID = writable($formData.academic_period_id);
	formDataAcademicPeriodID.subscribe((value) => ($formData.academic_period_id = value));
</script>

<form action="?/postResearchInnovationProjectList" use:enhance>
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
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
