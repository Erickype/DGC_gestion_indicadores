<script lang="ts">
	import { Skeleton } from '../ui/skeleton';

	export let tableHeightClass = 'h-[60vh]';

	const items = [1, 2, 3, 4, 5];
</script>

<div class="flex w-full items-center gap-2 rounded-sm py-4">
	<div class="w-3/4">
		<Skeleton class="h-10 w-full" />
	</div>
	<Skeleton class="h-10 w-12" />
	<Skeleton class="h-10 w-1/4" />
</div>
<div class="flex flex-col gap-1 rounded-md {tableHeightClass}">
	<Skeleton class="h-full w-full rounded-none rounded-t-md" />
	{#each items as _}
		<Skeleton class="h-full w-full rounded-none" />
	{/each}
	<Skeleton class="h-full w-full rounded-none rounded-b-md" />
	<div class="flex w-full justify-end gap-1">
		<Skeleton class="h-10 w-1/12" />
		<Skeleton class="h-10 w-1/12" />
		<Skeleton class="h-10 w-1/12" />
	</div>
</div>
