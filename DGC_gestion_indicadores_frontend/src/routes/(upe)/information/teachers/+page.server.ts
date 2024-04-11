import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/login")
    }

    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }
};