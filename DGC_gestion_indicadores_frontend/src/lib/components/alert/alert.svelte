<script lang="ts">
	import CircleAlert from 'lucide-svelte/icons/circle-alert';
	import * as Alert from '$lib/components/ui/alert/index.js';

	export let title = 'Error';
	export let description = 'Algo salió mal, intentar nuevamente.';
	export let variant: 'default' | 'destructive' | undefined = 'default';
</script>

<Alert.Root {variant}>
	<CircleAlert class="h-4 w-4" />
	<Alert.Title>{title}</Alert.Title>
	<Alert.Description>{description}</Alert.Description>
</Alert.Root>
