// See https://kit.svelte.dev/docs/types#app

import type { CommonError } from "$lib/api/model/errors";

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: {
				id: number,
				name?: string,
				role: number,
				iat: number,
				eat: number
			}
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
		namespace Superforms {
			export type Message = {
				success: boolean,
				data?: unknown
				error?: CommonError | string
			}
		}
	}
}

export { };
