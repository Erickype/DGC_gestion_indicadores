import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { addCarrerSchema, updateCarrerSchema } from "./schema";
import { LoadFacultiesWithComboMessages } from "$lib/api/controller/api/faculty";

export const load: PageServerLoad = async ({ locals, cookies }) => {
    if (!locals.user) {
        throw redirect(302, "/")
    }

    const token = cookies.get("AuthorizationToken")

    return {
        facultiesData: await LoadFacultiesWithComboMessages(token!),
        addCarrerForm: await superValidate(zod(addCarrerSchema)),
        updateCarrerForm: await superValidate(zod(updateCarrerSchema))
    }
};