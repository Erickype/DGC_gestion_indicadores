<script lang="ts">
	import DashboardIcon from '$lib/icons/dashboard.svelte';
	import FileIcon from '$lib/icons/file.svelte';
	import { page } from '$app/stores';

	let filteredMenus: Menu[] = [];

	$: menus = [
		{
			name: 'Dashboards',
			icon: DashboardIcon,
			links: [
				{ name: 'Periodos Evaluación', route: '/dashboards/evaluation-periods' },
				{ name: 'Periodos Académicos', route: '/dashboards/academic-periods' },
				{ name: 'Tasas académicas', route: '/dashboards/academic-fees' }
			],
			roles: [1, 2, 3]
		},
		{
			name: 'Información',
			icon: FileIcon,
			links: [
				{ name: 'Docentes', route: '/information/teachers' },
				{ name: 'Publicaciones Académicas', route: '/information/academic-publications' },
				{ name: 'Libros', route: '/information/books' }
			],
			roles: [1, 2]
		}
	];

	$: if ($page.data.user) {
		filteredMenus = menus.filter(
			(menu) => $page.data.user.role && menu.roles.includes($page.data.user.role)
		);
	}
</script>

{#if filteredMenus.length > 0}
	{#each filteredMenus as menu}
		<li>
			<details open>
				<summary>
					<svelte:component this={menu.icon}></svelte:component>
					{menu.name}
				</summary>
				<ul>
					{#each menu.links as link}
						<li>
							<a href={link.route}>{link.name}</a>
						</li>
					{/each}
				</ul>
			</details>
		</li>
	{/each}
{/if}
