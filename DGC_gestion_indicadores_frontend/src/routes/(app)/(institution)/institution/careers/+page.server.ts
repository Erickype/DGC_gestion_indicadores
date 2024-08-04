import type { Actions, PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { addCarrerSchema, updateCarrerSchema } from "./schema";

import { LoadFacultiesWithComboMessages } from "$lib/api/controller/api/faculty";
import type { PostCareerRequest } from "$lib/api/model/admin/career";
import { PostCareer } from "$lib/api/controller/admin/career";

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

export const actions: Actions = {
    addCareer: async (event) => {
        const form = await superValidate(event, zod(addCarrerSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const person: PostCareerRequest = {
            faculty_id: data.facultyID,
            name: data.name,
            description: data.description,
            abbreviation: data.abbreviation
        }

        const response = await PostCareer(person, token!)

        return generateFormMessageFromHttpResponse(form, response)
    },
};