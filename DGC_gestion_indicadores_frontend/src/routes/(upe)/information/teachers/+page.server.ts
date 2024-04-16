import { redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { LoadPeopleWithComboMessages } from "$lib/api/controller/api/person";
import { addTeacherSchema } from "./scheme";
import { message, superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { LoadCareersWithComboMessages } from "$lib/api/controller/api/career";
import { LoadDedicationsWithComboMessages } from "$lib/api/controller/api/dedication";
import { LoadScaledGradesWithComboMessages } from "$lib/api/controller/api/scaledGrade";
import type { LoginError, LoginRequest, LoginResponse } from '$lib/api/model/auth/login'

import { Login } from "$lib/api/controller/auth/auth";

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
        peopleData: await LoadPeopleWithComboMessages(token!),
        careersData: await LoadCareersWithComboMessages(token!),
        dedicationsData: await LoadDedicationsWithComboMessages(token!),
        scaledGradesData: await LoadScaledGradesWithComboMessages(token!),
        addTeacherForm: await superValidate(zod(addTeacherSchema)),
    }
};

export const actions: Actions = {
    addTeacher: async (event) => {
        const form = await superValidate(event, zod(addTeacherSchema))

        if (!form.valid) {
            return message(form, { success: false, error: "Invalid form" }, {
                status: 400
            })
        }

        const LoginRequest: LoginRequest = {
            username: "Erickype",
            password: "admin_password"
        }

        const res = await Login(LoginRequest)
        if (!res.ok) {
            const err: LoginError = await res.json()
            return message(form, { success: false, error: err.error })
        }

        return message(form, { success: true, error: "" })
    }
};