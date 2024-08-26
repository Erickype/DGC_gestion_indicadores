import { message, superValidate, type ErrorStatus } from "sveltekit-superforms";
import type { PageServerLoad } from "./$types";
import { redirect, type Actions } from "@sveltejs/kit";
import { zod } from "sveltekit-superforms/adapters";
import { addTeacherSchema, updateTeacherSchema } from "./schema";
import type { PostPersonRequest, PutPersonRequest } from "$lib/api/model/api/person";
import { PostPerson, PutPerson } from "$lib/api/controller/api/person";
import type { AddTeacherRequest } from "$lib/api/model/api/teacher";
import { CreateTeacher } from "$lib/api/controller/api/teacher";
import { generateFormMessageFromHttpResponse } from "$lib/utils";

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
            return message(form,
                { success: false, error: "Invalid form" },
                { status: 400 })
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
            return message(form,
                { success: false, error: "Invalid form" },
                { status: 400 })
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

        const res = await PutPerson(person, token!)
        if (res.status === 401) {
            throw redirect(302, "/")
        }
        if (!res.ok) {
            const status = res.status as unknown as ErrorStatus
            const err = await res.json()
            return message(form,
                { success: false, error: err },
                { status: status })
        }

        return message(form, { success: true, error: "" })
    }
};