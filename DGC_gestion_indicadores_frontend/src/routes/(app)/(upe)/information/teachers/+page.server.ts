import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { redirect } from "@sveltejs/kit";

import { addTeacherSchema, updateTeacherSchema } from "./scheme";
import { message, superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { LoadContractTypesWithComboMessages } from "$lib/api/controller/api/contractTypes";
import { LoadScaledGradesWithComboMessages } from "$lib/api/controller/api/scaledGrade";
import { LoadDedicationsWithComboMessages } from "$lib/api/controller/api/dedication";

import type { CreateTeacherRequest, UpdateTeacherRequest } from "$lib/api/model/api/teacher";
import { CreateTeacher, UpdateTeacher } from "$lib/api/controller/api/teacher";

import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { toast } from "svelte-sonner";

export const load: PageServerLoad = async ({ locals, cookies }) => {
    const token = cookies.get("AuthorizationToken")

    if (!locals.user) {
        throw redirect(302, "/login")
    }

    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    return {
        academicPeriodsData: await LoadAcademicPeriodsWithComboMessages(),
        dedicationsData: await LoadDedicationsWithComboMessages(token!),
        contractTypesData: await LoadContractTypesWithComboMessages(token!),
        scaledGradesData: await LoadScaledGradesWithComboMessages(token!),
        addTeacherForm: await superValidate(zod(addTeacherSchema)),
        updateTeacherForm: await superValidate(zod(updateTeacherSchema)),
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
        const createTeacherRequest: CreateTeacherRequest = {
            academic_period_id: data.academicPeriod,
            career_id: data.career,
            dedication_id: data.dedication,
            person_id: data.person,
            scaled_grade_id: data.scaledGrade,
            contract_type_id: data.contractType
        }

        const response = await CreateTeacher(token!, createTeacherRequest)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updateTeacher: async (event) => {
        const form = await superValidate(event, zod(updateTeacherSchema))
        const token = event.cookies.get("AuthorizationToken")

        if (!form.valid) {
            return message(form, { success: false, error: "Invalid form" }, {
                status: 400
            })
        }

        const data = form.data
        const updateTeacherRequest: UpdateTeacherRequest = {
            ID: data.ID,
            academic_period_id: data.academicPeriod,
            career_id: data.career,
            dedication_id: data.dedication,
            person_id: data.person,
            scaled_grade_id: data.scaledGrade
        }
        const teacherID = data.ID.toString()
        const res = await UpdateTeacher(token!, updateTeacherRequest, teacherID)

        if (!res.ok) {
            if (res.status === 401) {
                toast.warning("No está autenticado.")
                return redirect(302, "/login")
            }
            return message(form, { success: false, error: "Error actualizando docente" })
        }

        return message(form, { success: true, error: "" })
    }
};