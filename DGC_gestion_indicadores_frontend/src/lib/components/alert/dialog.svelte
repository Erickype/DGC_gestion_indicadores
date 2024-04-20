<script lang="ts">
	import * as AlertDialog from '$lib/components/ui/alert-dialog/index.js';

	export let dialogOpen = false;

	export let title = 'Esta totalmente seguro?';
	export let description =
		'Esta es una acción irreversible. Esto borrará permanentemente el registro sin opción a recuperarlo.';
</script>

<AlertDialog.Root bind:open={dialogOpen}>
	<AlertDialog.Content>
		<AlertDialog.Header>
			<AlertDialog.Title>{title}</AlertDialog.Title>
			<AlertDialog.Description>
				{description}
			</AlertDialog.Description>
		</AlertDialog.Header>
		<AlertDialog.Footer>
			<AlertDialog.Cancel>Cancelar</AlertDialog.Cancel>
			<AlertDialog.Action>Continuar</AlertDialog.Action>
		</AlertDialog.Footer>
	</AlertDialog.Content>
</AlertDialog.Root>
