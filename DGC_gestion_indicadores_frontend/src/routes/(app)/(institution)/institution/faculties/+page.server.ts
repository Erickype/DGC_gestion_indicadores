import { message, superValidate, type ErrorStatus } from "sveltekit-superforms";
import { redirect, type Actions } from "@sveltejs/kit";
import { zod } from "sveltekit-superforms/adapters";

import { addFacultySchema, updateFacultySchema } from "./schema";

import type { PageServerLoad } from "./$types";

import type { PostFacultyRequest, PutFacultyRequest } from "$lib/api/model/admin/faculty";
import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { PostFaculty, PutFaculty } from "$lib/api/controller/admin/faculty";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/")
    }

    return {
        addFacultyForm: await superValidate(zod(addFacultySchema)),
        updateFacultyForm: await superValidate(zod(updateFacultySchema))
    }
};

export const actions: Actions = {
    addFaculty: async (event) => {
        const form = await superValidate(event, zod(addFacultySchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const person: PostFacultyRequest = {
            name: data.name,
            description: data.description,
            abbreviation: data.abbreviation
        }

        const response = await PostFaculty(person, token!)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updateFaculty: async (event) => {
        const form = await superValidate(event, zod(updateFacultySchema))
        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const faculty: PutFacultyRequest = {
            ID: data.ID,
            name: data.name,
            abbreviation: data.abbreviation,
            description: data.description
        }

        const response = await PutFaculty(faculty, token!)

        return generateFormMessageFromHttpResponse(form, response)
    }
};