import { message, superValidate, type ErrorStatus } from "sveltekit-superforms";
import type { PageServerLoad } from "./$types";
import { redirect, type Actions } from "@sveltejs/kit";
import { zod } from "sveltekit-superforms/adapters";
import { addPersonSchema, updatePersonSchema } from "./schema";
import type { PostPersonRequest } from "$lib/api/model/api/person";
import { PostPerson } from "$lib/api/controller/admin/person";


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
    addPerson: async (event) => {
        const form = await superValidate(event, zod(addPersonSchema))

        if (!form.valid) {
            return message(form,
                { success: false, error: "Invalid form" },
                { status: 400 })
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const person: PostPersonRequest = {
            identity: data.identity,
            name: data.name,
            lastname: data.lastname,
            email: data.email
        }

        const res = await PostPerson(person, token!)

        if (!res.ok) {
            const status = res.status as unknown as ErrorStatus
            const data = await res.json()
            return message(form,
                { success: false, error: data },
                { status: status })
        }
        return message(form, {
            success: true
        })
    },
};