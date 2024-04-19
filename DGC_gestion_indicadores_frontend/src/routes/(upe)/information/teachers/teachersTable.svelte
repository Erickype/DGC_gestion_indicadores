<script lang="ts">
	import type { Teacher } from '$lib/api/model/api/teacher';

	import * as Table from '$lib/components/ui/table';

	export let teachers: Teacher[];
</script>

<Table.Root>
	<Table.Caption
		>Lista de docentes en el periodo {teachers.at(0)?.academic_period_id}.</Table.Caption
	>
	<Table.Header>
		<Table.Row>
			<Table.Head class="w-[100px]">Fecha Modificación</Table.Head>
			<Table.Head>Persona</Table.Head>
			<Table.Head>Carrera</Table.Head>
			<Table.Head>Grado escalafonado</Table.Head>
			<Table.Head class="text-right">Dedicación</Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body>
		{#each teachers as teacher}
			<Table.Row>
				<Table.Cell class="font-medium">{teacher.UpdatedAt}</Table.Cell>
				<Table.Cell>{teacher.person_id}</Table.Cell>
				<Table.Cell>{teacher.career_id}</Table.Cell>
				<Table.Cell>{teacher.scaled_grade_id}</Table.Cell>
				<Table.Cell class="text-right">{teacher.dedication_id}</Table.Cell>
			</Table.Row>
		{/each}
	</Table.Body>
</Table.Root>
