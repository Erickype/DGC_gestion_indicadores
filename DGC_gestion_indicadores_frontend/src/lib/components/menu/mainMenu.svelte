<script lang="ts">
	import MonitrorCheck from 'lucide-svelte/icons/monitor-check';

	import { page } from '$app/stores';

	import * as Accordion from '$lib/components/ui/accordion';
	import { menus, type Menu } from './menuStruct';

	let filteredMenus: Menu[] = [];

	$: if ($page.data.user) {
		filteredMenus = menus.filter(
			(menu) => $page.data.user.role && menu.roles.includes($page.data.user.role)
		);
	}
</script>

<div class="bg-muted/40 hidden border-r md:block">
	<div class="flex h-full max-h-screen flex-col gap-2">
		<div class="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6">
			<a href="/" class="flex items-center gap-2 font-semibold">
				<MonitrorCheck class="h-6 w-6" />
				<span class="">DGCA</span>
			</a>
		</div>
		<div class="flex-1">
			<nav class="grid items-start px-2 text-sm font-medium lg:px-4">
				<Accordion.Root>
					{#each filteredMenus as menu, i}
						<Accordion.Item value="item-{i}">
							<Accordion.Trigger>
								<div class="flex items-center justify-start gap-2 px-3">
									<svelte:component this={menu.icon} class="h-4 w-4" />
									{menu.name}
								</div>
							</Accordion.Trigger>
							<Accordion.Content>
								{#each menu.links as link}
									<a
										href={link.route}
										class="text-muted-foreground hover:text-primary mx-4 my-2 flex items-center gap-3 rounded-lg py-2 transition-all"
									>
										{link.name}
									</a>
								{/each}
							</Accordion.Content>
						</Accordion.Item>
					{/each}
				</Accordion.Root>
			</nav>
		</div>
	</div>
</div>
