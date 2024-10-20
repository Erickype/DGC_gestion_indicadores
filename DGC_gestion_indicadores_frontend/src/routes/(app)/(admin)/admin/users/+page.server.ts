import type { PageServerLoad } from "./$types";
import { redirect, type Actions } from "@sveltejs/kit";

import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { updateUserSchema } from "./schema";

import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import type { UpdateUsersRequest } from "$lib/api/model/admin/user";
import { UpdateUser } from "$lib/api/controller/admin/user";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/login")
    }

    if (locals.user.role !== 1) {
        throw redirect(302, "/")
    }

    return {
        updateUserForm: await superValidate(zod(updateUserSchema)),
    }
};

export const actions: Actions = {
    updateUser: async (event) => {
        const form = await superValidate(event, zod(updateUserSchema))
        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const user: UpdateUsersRequest = {
            ID: data.ID,
            username: data.username,
            email: data.email,
            role_id: data.role_id
        }

        const response = await UpdateUser(token!, user)

        return generateFormMessageFromHttpResponse(form, response)
    }
};