<script lang="ts">
	import { DateFormatter, getLocalTimeZone, today, type DateValue } from '@internationalized/date';
	import { createEventDispatcher, type EventDispatcher } from 'svelte';
	import * as Calendar from '$lib/components/ui/calendar/index.js';
	import * as Select from '$lib/components/ui/select/index.js';
	import { Calendar as CalendarPrimitive } from 'bits-ui';
	import { cn } from '$lib/utils.js';

	type $$Events = CalendarPrimitive.Events;
	type $$Props = CalendarPrimitive.Props;

	export let placeholder: $$Props['placeholder'] = today(getLocalTimeZone());
	export let weekdayFormat: $$Props['weekdayFormat'] = 'short';
	export let value: $$Props['value'] = undefined;

	export let monthLabels = [
		'enero',
		'febrero',
		'marzo',
		'abril',
		'mayo',
		'junio',
		'julio',
		'agosto',
		'septiembre',
		'octubre',
		'noviembre',
		'diciembre'
	];
	const monthOptions = monthLabels.map((month, i) => ({ value: i + 1, label: month }));

	const monthFmt = new DateFormatter('ec-EC', {
		month: 'long'
	});

	export let startYear = today(getLocalTimeZone()).year - 5;
	export let endYear = today(getLocalTimeZone()).year + 5;

	const yearOptions = Array.from({ length: endYear - startYear + 1 }, (_, i) => ({
		label: String(endYear - i),
		value: endYear - i
	}));

	$: defaultYear = placeholder
		? {
				value: placeholder.year,
				label: String(placeholder.year)
			}
		: undefined;

	$: defaultMonth = placeholder
		? {
				value: placeholder.month,
				label: monthFmt.format(placeholder.toDate(getLocalTimeZone()))
			}
		: undefined;

	const dispatcher: EventDispatcher<any> = createEventDispatcher();

	function sendDateSelected(v: DateValue | DateValue[] | undefined) {
		let date: string = '';
		if (v) {
			date = v.toString();
		} else {
			date = '';
		}
		dispatcher('date-selected', {
			date: date
		});
	}

	let className: $$Props['class'] = undefined;
	let open: false;

	export { className as class };
</script>

<CalendarPrimitive.Root
	bind:hidden={open}
	{weekdayFormat}
	class={cn('rounded-md border p-3', className)}
	{...$$restProps}
	on:keydown
	onValueChange={(v) => {
		open = false;
		sendDateSelected(v);
	}}
	let:months
	let:weekdays
	bind:value
	bind:placeholder
>
	<Calendar.Header>
		<Calendar.Heading class="flex w-full items-center justify-between gap-2">
			<Select.Root
				selected={defaultMonth}
				items={monthOptions}
				onSelectedChange={(v) => {
					if (!v || !placeholder) return;
					if (v.value === placeholder?.month) return;
					placeholder = placeholder.set({ month: v.value });
				}}
			>
				<Select.Trigger aria-label="Elegir mes" class="w-[60%]">
					<Select.Value placeholder="Elegir mes" />
				</Select.Trigger>
				<Select.Content class="max-h-[200px] overflow-y-auto">
					{#each monthOptions as { value, label }}
						<Select.Item {value} {label}>
							{label}
						</Select.Item>
					{/each}
				</Select.Content>
			</Select.Root>
			<Select.Root
				selected={defaultYear}
				items={yearOptions}
				onSelectedChange={(v) => {
					if (!v || !placeholder) return;
					if (v.value === placeholder?.year) return;
					placeholder = placeholder.set({ year: v.value });
				}}
			>
				<Select.Trigger aria-label="Elegir año" class="w-[40%]">
					<Select.Value placeholder="Elegir año" />
				</Select.Trigger>
				<Select.Content class="max-h-[200px] overflow-y-auto">
					{#each yearOptions as { value, label }}
						<Select.Item {value} {label}>
							{label}
						</Select.Item>
					{/each}
				</Select.Content>
			</Select.Root>
		</Calendar.Heading>
	</Calendar.Header>
	<Calendar.Months>
		{#each months as month}
			<Calendar.Grid>
				<Calendar.GridHead>
					<Calendar.GridRow class="flex">
						{#each weekdays as weekday}
							<Calendar.HeadCell>
								{weekday.slice(0, 2)}
							</Calendar.HeadCell>
						{/each}
					</Calendar.GridRow>
				</Calendar.GridHead>
				<Calendar.GridBody>
					{#each month.weeks as weekDates}
						<Calendar.GridRow class="mt-2 w-full">
							{#each weekDates as date}
								<Calendar.Cell {date}>
									<Calendar.Day {date} month={month.value} />
								</Calendar.Cell>
							{/each}
						</Calendar.GridRow>
					{/each}
				</Calendar.GridBody>
			</Calendar.Grid>
		{/each}
	</Calendar.Months>
</CalendarPrimitive.Root>
