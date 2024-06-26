import type { PageServerLoad } from "./$types";
import { redirect, type Actions } from "@sveltejs/kit";
import { addAcademicPeriodSchema, updateAcademicPeriodSchema } from "./schema";
import { message, superValidate, type ErrorStatus } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import type { PostAcademicPeriodRequest, UpdateAcademicPeriodRequest } from "$lib/api/model/view/academicPeriod";
import { PostAcademicPeriod, UpdateAcademicPeriod } from "$lib/api/controller/admin/academicPeriod";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/")
    }
    return {
        addAcademiPeriodForm: await superValidate(zod(addAcademicPeriodSchema)),
        updateAcademicPeriodForm: await superValidate(zod(updateAcademicPeriodSchema))
    }
};

export const actions: Actions = {
    addAcademicPeriod: async (event) => {
        const form = await superValidate(event, zod(addAcademicPeriodSchema))

        if (!form.valid) {
            return message(form,
                { success: false, error: "Invalid form" },
                { status: 400 })
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const academicPeriod: PostAcademicPeriodRequest = {
            name: data.name,
            abbreviation: data.abbreviation,
            description: data.description,
            start_date: new Date(data.startDate).toISOString(),
            end_date: new Date(data.endDate).toISOString()
        }

        const res = await PostAcademicPeriod(academicPeriod, token!)

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

    updateAcademicPeriod: async (event) => {
        const form = await superValidate(event, zod(updateAcademicPeriodSchema))
        if (!form.valid) {
            return message(form,
                { success: false, error: "Invalid form" },
                { status: 400 })
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const academicPeriod: UpdateAcademicPeriodRequest = {
            ID: data.ID,
            name: data.name,
            abbreviation: data.abbreviation,
            description: data.description,
            start_date: new Date(data.startDate).toISOString(),
            end_date: new Date(data.endDate).toISOString()
        }

        const res = await UpdateAcademicPeriod(academicPeriod, token!)
        if(res.status === 401){
            throw redirect(302,"/")
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