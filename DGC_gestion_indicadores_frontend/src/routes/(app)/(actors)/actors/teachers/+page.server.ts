import { addTeacherSchema, updateTeacherSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { redirect, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import type { AddTeacherRequest } from "$lib/api/model/api/teacher";
import type { PutPersonRequest } from "$lib/api/model/api/person";
import { CreateTeacher } from "$lib/api/controller/api/teacher";
import { PutPerson } from "$lib/api/controller/api/person";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/")
    }

    return {
        addPersonForm: await superValidate(zod(addTeacherSchema)),
        updatePersonForm: await superValidate(zod(updateTeacherSchema))
    }
};

export const actions: Actions = {
    addTeacher: async (event) => {
        const form = await superValidate(event, zod(addTeacherSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const addTeacherRequest: AddTeacherRequest = {
            person_id: data.person_id
        }

        const response = await CreateTeacher(token!, addTeacherRequest)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updatePerson: async (event) => {
        const form = await superValidate(event, zod(updateTeacherSchema))
        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const person: PutPersonRequest = {
            ID: data.ID,
            identity: data.identity,
            name: data.name,
            lastname: data.lastname,
            email: data.email
        }

        const response = await PutPerson(person, token!)

        return generateFormMessageFromHttpResponse(form, response)
    }
};