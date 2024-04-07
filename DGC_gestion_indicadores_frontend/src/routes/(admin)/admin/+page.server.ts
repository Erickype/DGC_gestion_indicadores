import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/login")
    }

    const role = locals.user.role
    switch (role) {
        case 2:
            throw redirect(302, "/information")
        case 3:
            throw redirect(302, "/")
    }
};