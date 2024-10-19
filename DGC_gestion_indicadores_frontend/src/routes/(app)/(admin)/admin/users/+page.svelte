<script lang="ts">
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';

	import Icon from 'lucide-svelte/icons/user-round-search';
	import Alert from '$lib/components/alert/alert.svelte';

	import type { CommonError } from '$lib/api/model/errors';
	import { toast } from 'svelte-sonner';
	import Table from './Table.svelte';

	import type { User } from '$lib/api/model/admin/user';

	export let data: PageData;

	const updateUserForm = data.updateUserForm;

	let usersPromise: Promise<User[]> = fetchUsers();

	async function fetchUsers() {
		const url = `/api/admin/users`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			return toast.error(errorData.message);
		}
		return (usersPromise = response.json());
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			fetchUsers();
		}
	}
</script>

<svelte:head>
	<title>Usuarios</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Usuarios</h2>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center px-8">
	{#await usersPromise}
		cargando...
	{:then users}
		{#if users.length > 0}
			<Table
				{users}
				formData={updateUserForm}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
			></Table>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay usuarios'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
