import type { PostEvaluationPeriodRequest, UpdateEvaluationPeriodRequest } from "$lib/api/model/view/evaluationPeriod";
import { PostEvaluationPeriod, UpdateEvaluationPeriod } from "$lib/api/controller/admin/evaluationPeriod";

import { message, superValidate, type ErrorStatus } from "sveltekit-superforms";
import { addEvaluationPeriodSchema, updateEvaluationPeriodSchema } from "./schema";
import { zod } from "sveltekit-superforms/adapters";

import { redirect, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/")
    }
    return {
        addEvaluationPeriodForm: await superValidate(zod(addEvaluationPeriodSchema)),
        updateEvaluationPeriodForm: await superValidate(zod(updateEvaluationPeriodSchema))
    }
};

export const actions: Actions = {
    addEvaluationPeriod: async (event) => {
        const form = await superValidate(event, zod(addEvaluationPeriodSchema))

        if (!form.valid) {
            return message(form,
                { success: false, error: "Invalid form" },
                { status: 400 })
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const evaluationPeriod: PostEvaluationPeriodRequest = {
            name: data.name,
            abbreviation: data.abbreviation,
            description: data.description,
            start_year: new Date(data.startDate).toISOString(),
            end_year: new Date(data.endDate).toISOString()
        }

        const res = await PostEvaluationPeriod(evaluationPeriod, token!)

        if (!res.ok) {
            const status = res.status as unknown as ErrorStatus
            const message = await res.json()
            return message(form,
                { success: false, error: message },
                { status: status })
        }
        return message(form, { success: true, error: "" })
    },
    
    updateEvaluationPeriod: async (event) => {
        const form = await superValidate(event, zod(updateEvaluationPeriodSchema))
        if (!form.valid) {
            return message(form,
                { success: false, error: "Invalid form" },
                { status: 400 })
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const evaluationPeriod: UpdateEvaluationPeriodRequest = {
            ID: data.ID,
            name: data.name,
            abbreviation: data.abbreviation,
            description: data.description,
            start_year: new Date(data.startDate).toISOString(),
            end_year: new Date(data.endDate).toISOString()
        }

        const res = await UpdateEvaluationPeriod(evaluationPeriod, token!)
        if (!res.ok) {
            const status = res.status as unknown as ErrorStatus
            const message = await res.json()
            return message(form,
                { success: false, error: message },
                { status: status })
        }

        return message(form, { success: true, error: "" })
    }
};