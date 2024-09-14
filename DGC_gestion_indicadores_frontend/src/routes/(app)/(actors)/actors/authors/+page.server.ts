import { message, superValidate, type ErrorStatus } from "sveltekit-superforms";
import { redirect, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

import { addAuthorFromPersonSchema, addAuthorSchema, updateAuthorPersonSchema } from "./schema";
import { zod } from "sveltekit-superforms/adapters";

import type { Author, PostAuthorFromPersonRequest } from "$lib/api/model/api/academicProduction/authors/author";
import { PostAuthor, PostAuthorFromPerson } from "$lib/api/controller/api/academicProduction/authors/author";
import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import type { PutPersonRequest } from "$lib/api/model/api/person";
import { PutPerson } from "$lib/api/controller/api/person";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/")
    }

    return {
        addAuthorForm: await superValidate(zod(addAuthorSchema)),
        addAuthorFromPersonForm: await superValidate(zod(addAuthorFromPersonSchema)),
        updateAuthorPersonForm: await superValidate(zod(updateAuthorPersonSchema))
    }
};

export const actions: Actions = {
    addAuthor: async (event) => {
        const form = await superValidate(event, zod(addAuthorSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const addAuthorRequest: Author = {
            person_id: data.person_id
        }

        const response = await PostAuthor(token!, addAuthorRequest)

        return generateFormMessageFromHttpResponse(form, response)
    },

    postAuthorFromPerson: async (event) => {
        const form = await superValidate(event, zod(addAuthorFromPersonSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const postAuthorFormPersonRequest: PostAuthorFromPersonRequest = {
            identity: data.identity,
            name: data.name,
            lastname: data.lastname,
            email: data.email
        }

        const response = await PostAuthorFromPerson(token!, postAuthorFormPersonRequest)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updatePerson: async (event) => {
        const form = await superValidate(event, zod(updateAuthorPersonSchema))
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