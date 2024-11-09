import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { redirect } from "@sveltejs/kit";

import { addPostgraduateProgramSchema, updatePostgraduateProgramSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { PostPostgraduateProgram, UpdatePostgraduateProgram } from "$lib/api/controller/api/indicatorsInformation/postgraduate";
import type { PostgraduateProgram } from "$lib/api/model/api/indicatorsInformation/postgraduate";

import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";

export const load: PageServerLoad = async ({ locals, cookies }) => {
    const token = cookies.get("AuthorizationToken")

    if (!locals.user) {
        throw redirect(302, "/login")
    }

    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }
    return {
        addPostGraduateProgramForm: await superValidate(zod(addPostgraduateProgramSchema)),
        updatePostGraduateProgramForm: await superValidate(zod(updatePostgraduateProgramSchema))
    }
};

export const actions: Actions = {
    postPostgraduateProgram: async (event) => {
        const form = await superValidate(event, zod(addPostgraduateProgramSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const postGraduateProgram: PostgraduateProgram = {
            name: data.name,
            start_year: data.start_year,
            end_year: data.end_year,
            is_active: data.is_active
        }

        const response = await PostPostgraduateProgram(token!, postGraduateProgram)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updatePostgraduateProgram: async (event) => {
        const form = await superValidate(event, zod(updatePostgraduateProgramSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const postGraduateProgram: PostgraduateProgram = {
            ID: data.id,
            name: data.name,
            start_year: data.start_year,
            end_year: data.end_year,
            is_active: data.is_active
        }

        const response = await UpdatePostgraduateProgram(token!, postGraduateProgram)

        return generateFormMessageFromHttpResponse(form, response)
    },
};