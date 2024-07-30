import { redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { LoadPeopleWithComboMessages } from "$lib/api/controller/api/person";
import { addTeacherSchema, updateTeacherSchema } from "./scheme";
import { message, superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { LoadCareersWithComboMessages } from "$lib/api/controller/api/career";
import { LoadDedicationsWithComboMessages } from "$lib/api/controller/api/dedication";
import { LoadScaledGradesWithComboMessages } from "$lib/api/controller/api/scaledGrade";
import type { LoginError } from '$lib/api/model/auth/login'

import { CreateTeacher, GetTeachersByAcademicPeriodID, UpdateTeacher } from "$lib/api/controller/api/teacher";
import type { CreateTeacherRequest, GetTeachersByAcademicPeriodResponse, UpdateTeacherRequest } from "$lib/api/model/api/teacher";
import { toast } from "svelte-sonner";

export const load: PageServerLoad = async ({ locals, cookies }) => {
    const token = cookies.get("AuthorizationToken")

    if (!locals.user) {
        throw redirect(302, "/login")
    }

    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    const academicPeriodsData = await LoadAcademicPeriodsWithComboMessages()
    const academicPeriods = academicPeriodsData.periods
    const selectedAcademicPeriod = academicPeriods.at(
        academicPeriods.length - 1
    )!.ID;
    const academicPeriodID = selectedAcademicPeriod.toString()
    const teachersByAcademicPeriod = await GetTeachersByAcademicPeriodID(token!, academicPeriodID)
    return {
        academicPeriodsData: academicPeriodsData,
        teachersByAcademicPeriod: teachersByAcademicPeriod.json() as Promise<GetTeachersByAcademicPeriodResponse[]>,
        peopleData: await LoadPeopleWithComboMessages(token!),
        careersData: await LoadCareersWithComboMessages(token!),
        dedicationsData: await LoadDedicationsWithComboMessages(token!),
        scaledGradesData: await LoadScaledGradesWithComboMessages(token!),
        addTeacherForm: await superValidate(zod(addTeacherSchema)),
        updateTeacherForm: await superValidate(zod(updateTeacherSchema)),
    }
};

export const actions: Actions = {
    addTeacher: async (event) => {
        const form = await superValidate(event, zod(addTeacherSchema))
        const token = event.cookies.get("AuthorizationToken")

        if (!form.valid) {
            return message(form, { success: false, error: "Invalid form" }, {
                status: 400
            })
        }

        const data = form.data
        const createTeacherRequest: CreateTeacherRequest = {
            academic_period_id: data.academicPeriod,
            career_id: data.career,
            dedication_id: data.dedication,
            person_id: data.person,
            scaled_grade_id: data.scaledGrade
        }
        const res = await CreateTeacher(token!, createTeacherRequest)
        if (!res.ok) {
            const err: LoginError = await res.json()
            return message(form, { success: false, error: err.error })
        }

        return message(form, { success: true, error: "" })
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
            if(res.status === 401){
                toast.warning("No está autenticado.")
                return redirect(302, "/login")
            }
            return message(form, { success: false, error: "Error actualizando docente" })
        }

        return message(form, { success: true, error: "" })
    }
};