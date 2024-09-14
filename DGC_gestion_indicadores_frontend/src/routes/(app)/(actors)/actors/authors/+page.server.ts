import { message, superValidate, type ErrorStatus } from "sveltekit-superforms";
import { redirect, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

import { addAuthorSchema, updateAuthorPersonSchema } from "./schema";
import { zod } from "sveltekit-superforms/adapters";

import { generateFormMessageFromHttpResponse } from "$lib/utils";

import { PostAuthor } from "$lib/api/controller/api/academicProduction/authors/author";
import type { Author } from "$lib/api/model/api/academicProduction/authors/author";
import type { PutPersonRequest } from "$lib/api/model/api/person";
import { PutPerson } from "$lib/api/controller/api/person";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/")
    }

    return {
        addAuthorForm: await superValidate(zod(addAuthorSchema)),
        updateAuthorPersonForm: await superValidate(zod(updateAuthorPersonSchema))
    }
};

export const actions: Actions = {
    addAuthor: async (event) => {
        const form = await superValidate(event, zod(addAuthorSchema))

        if (!form.valid) {
            return message(form,
                { success: false, error: "Invalid form" },
                { status: 400 })
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const addAuthorRequest: Author = {
            person_id: data.person_id
        }

        const response = await PostAuthor(token!, addAuthorRequest)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updatePerson: async (event) => {
        const form = await superValidate(event, zod(updateAuthorPersonSchema))
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