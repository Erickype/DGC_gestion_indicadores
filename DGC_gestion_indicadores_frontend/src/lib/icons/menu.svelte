<script lang="ts">
	export let fill = 'fill-accent';
	export let stroke = 'stroke-secondary-content';
	export let width = 'w-8';
	export let height = 'h-8';
	export let opacity = 'opacity-80';
</script>

<svg
	class="{fill} {stroke} {width} {height} {opacity}"
	viewBox="0 0 24 24"
	fill="none"
	xmlns="http://www.w3.org/2000/svg"
>
	<path
		d="M4 6H20M4 12H20M4 18H20"
		{stroke}
		stroke-width="2"
		stroke-linecap="round"
		stroke-linejoin="round"
	/>
</svg>
