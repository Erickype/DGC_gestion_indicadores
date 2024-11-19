<script lang="ts">
	import { scaleLinear, scaleBand } from 'd3-scale';

	import { Chart, Svg, Axis, Spline, Highlight, Tooltip, LinearGradient, Bars } from 'layerchart';
	import { format, PeriodType } from '@layerstack/utils';

	let chartData = [
		{
			date: new Date('2024-10-21T05:00:00.000Z'),
			value: 73,
			baseline: 68
		},
		{
			date: new Date('2024-10-22T05:00:00.000Z'),
			value: 70,
			baseline: 85
		},
		{
			date: new Date('2024-10-23T05:00:00.000Z'),
			value: 66,
			baseline: 42
		},
		{
			date: new Date('2024-10-24T05:00:00.000Z'),
			value: 68,
			baseline: 74
		},
		{
			date: new Date('2024-10-25T05:00:00.000Z'),
			value: 86,
			baseline: 39
		},
		{
			date: new Date('2024-10-26T05:00:00.000Z'),
			value: 64,
			baseline: 28
		},
		{
			date: new Date('2024-10-27T05:00:00.000Z'),
			value: 94,
			baseline: 48
		}
	];
	const colorKeys = [...new Set(chartData.map((x) => x.date))];

	const keyColors = [
		'hsl(212 100% 40.6%)',
		'hsl(175 100% 50%)',
		'hsl(289 100% 68%)',
		'hsl(60 100% 61%)'
	];
</script>

<svelte:head>
	<title>Home</title>
</svelte:head>

<div class="h-[300px] rounded border p-4">
	<Chart
		data={chartData}
		x="date"
		xScale={scaleBand().padding(0.4)}
		y="value"
		c="date"
		yDomain={[0, null]}
		cDomain={colorKeys}
		cRange={keyColors}
		yNice={4}
		padding={{ left: 16, bottom: 24 }}
	>
		<Svg>
			<Axis placement="left" grid rule />
			<Axis
				placement="bottom"
				format={(d) => format(d, PeriodType.Day, { variant: 'short' })}
				rule
			/>
			<Bars strokeWidth={0} />
			<Highlight area />
		</Svg>
		<Tooltip.Root let:data>
			<Tooltip.Header>{data.date}</Tooltip.Header>
			<Tooltip.List>
				<Tooltip.Item label={data.date} value={data.value} format="integer" valueAlign="right" />
			</Tooltip.List>
		</Tooltip.Root>
	</Chart>
</div>