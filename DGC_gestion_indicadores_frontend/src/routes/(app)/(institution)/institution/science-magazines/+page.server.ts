import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { redirect } from "@sveltejs/kit";

import { addScienceMagazineSchema, updateScienceMagazineSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { PostScienceMagazine, UpdateScienceMagazine } from "$lib/api/controller/api/academicProduction/scienceMagazines/scienceMagazine";
import { LoadAcademicDatabasesWithComboMessages } from "$lib/api/controller/api/academicProduction/academicDatabases/academicDatabases";
import type { ScienceMagazine } from "$lib/api/model/api/academicProduction/scienceMagazines/scienceMagazine";

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
        addScienceMagazineForm: await superValidate(zod(addScienceMagazineSchema)),
        updateScienceMagazineForm: await superValidate(zod(updateScienceMagazineSchema)),
        academicDatabasesData: await LoadAcademicDatabasesWithComboMessages(token!)
    }
};

export const actions: Actions = {
    postScienceMagazine: async (event) => {
        const form = await superValidate(event, zod(addScienceMagazineSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const scienceMagazine: ScienceMagazine = {
            name: data.name,
            abbreviation: data.abbreviation,
            description: data.description,
            academic_database_id: data.academic_database_id
        }

        const response = await PostScienceMagazine(token!, scienceMagazine)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updateScienceMagazine: async (event) => {
        const form = await superValidate(event, zod(updateScienceMagazineSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const scienceMagazine: ScienceMagazine = {
            ID: data.id,
            name: data.name,
            abbreviation: data.abbreviation,
            description: data.description,
            academic_database_id: data.academic_database_id
        }

        const response = await UpdateScienceMagazine(token!, scienceMagazine)

        return generateFormMessageFromHttpResponse(form, response)
    },
};