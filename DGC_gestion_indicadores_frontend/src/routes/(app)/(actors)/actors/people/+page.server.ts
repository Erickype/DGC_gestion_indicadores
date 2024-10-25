import { addPersonSchema, updatePersonSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { redirect, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

import type { Person, PostPersonWithRolesRequest, UpdatePersonWithRolesRequest } from "$lib/api/model/api/person";
import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { PostPersonWithRoles, UpdatePersonWithRoles } from "$lib/api/controller/api/person";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/")
    }

    return {
        addPersonForm: await superValidate(zod(addPersonSchema)),
        updatePersonForm: await superValidate(zod(updatePersonSchema))
    }
};

export const actions: Actions = {
    addPersonWithRoles: async (event) => {
        const form = await superValidate(event, zod(addPersonSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const person: Person = {
            identity: data.identity,
            name: data.name,
            lastname: data.lastname,
            email: data.email,
        }
        const request: PostPersonWithRolesRequest = {
            person: person,
            roles: data.roles
        }

        const response = await PostPersonWithRoles(request, token!)

        return generateFormMessageFromHttpResponse(form, response)

    },

    updatePersonWithRoles: async (event) => {
        const form = await superValidate(event, zod(updatePersonSchema))
        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const person: Person = {
            ID: data.ID,
            identity: data.identity,
            name: data.name,
            lastname: data.lastname,
            email: data.email,
        }
        const request: UpdatePersonWithRolesRequest = {
            person: person,
            roles: data.roles
        }

        const response = await UpdatePersonWithRoles(request, token!)

        return generateFormMessageFromHttpResponse(form, response)
    }
};