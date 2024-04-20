<script lang="ts">
	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import { updateTeacherAction } from '../../../../stores';
	import type { AddTeacherSchema } from './scheme';
	import type { Message } from '$lib/components/combobox/combobox';
	import type { Teacher } from '$lib/api/model/api/teacher';

	let updateTeacherActionValue: { status: boolean; teacherID: number };

	const unsubscribe = updateTeacherAction.subscribe((value) => {
		updateTeacherActionValue = value;
	});

	/*  export let data: SuperValidated<Infer<AddTeacherSchema>>;
	export let academicPeriod: number;
	export let people: Message[];
	export let careers: Message[];
	export let dedications: Message[];
	export let scaledGrades: Message[]; */

	export let selectedTeacherToUpdate: Teacher;

</script>

<pre>
    {JSON.stringify(selectedTeacherToUpdate)}
</pre>
