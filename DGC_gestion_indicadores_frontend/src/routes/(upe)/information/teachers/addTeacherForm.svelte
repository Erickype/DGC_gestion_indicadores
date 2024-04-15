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
	import { tick } from 'svelte';
	import { toast } from 'svelte-sonner';
	import { browser } from '$app/environment';

	export let data: SuperValidated<Infer<AddTeacherSchema>>;
	export let people: Message[];

	const form = superForm(data, {
		validators: zodClient(addTeacherSchema),
		onUpdated: ({ form: f }) => {
			if (f.valid) {
				toast.success(`You submitted ${JSON.stringify(f.data, null, 2)}`);
			} else {
				toast.error('Please fix the errors in the form.');
			}
		}
	});

	const { form: formData, message, enhance } = form;

	let open = false;

	function closeAndFocusTrigger(triggerId: string) {
		open = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
</script>

<form action="?/addTeacher" method="post" use:enhance>
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
					{people.find((f) => f.value === $formData.person)?.label ?? 'Select language'}
					<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
				</Popover.Trigger>
				<input hidden value={$formData.person} name={attrs.name} />
			</Form.Control>
			<Popover.Content class="w-[200px] p-0">
				<Command.Root>
					<Command.Input autofocus placeholder="Search language..." class="h-9" />
					<Command.Empty>No language found.</Command.Empty>
					<Command.Group>
						{#each people as person}
							<Command.Item
								value={person.value}
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
		<Form.Description>This is the language that will be used in the dashboard.</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button>Submit</Form.Button>
	{#if browser}
		<SuperDebug data={$formData} />
	{/if}
</form>
