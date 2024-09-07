import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { redirect } from "@sveltejs/kit";

import { addTeacherSchema, updateTeacherSchema } from "./scheme";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { LoadContractTypesWithComboMessages } from "$lib/api/controller/api/contractTypes";
import { LoadScaledGradesWithComboMessages } from "$lib/api/controller/api/scaledGrade";
import { LoadDedicationsWithComboMessages } from "$lib/api/controller/api/dedication";
import { LoadCareersWithComboMessages } from "$lib/api/controller/api/career";

import type { CreateTeachersListsRequest, UpdateTeachersListRequest } from "$lib/api/model/api/indicatorsInformation/teachersLists";
import { CreateTeachersList, PatchTeachersLists } from "$lib/api/controller/api/indicatorsInformation/teachersList";
import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";

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
        careersData: await LoadCareersWithComboMessages(token!),
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
        const createTeachersListRequest: CreateTeachersListsRequest = {
            academic_period_id: data.academicPeriod,
            career_id: data.career,
            dedication_id: data.dedication,
            teacher_id: data.teacher,
            scaled_grade_id: data.scaledGrade,
            contract_type_id: data.contractType
        }

        const response = await CreateTeachersList(token!, createTeachersListRequest)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updateTeacher: async (event) => {
        const form = await superValidate(event, zod(updateTeacherSchema))
        const token = event.cookies.get("AuthorizationToken")

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const data = form.data
        const updateTeacherRequest: UpdateTeachersListRequest = {
            academic_period_id: data.academicPeriod,
            teacher_id: data.teacher,
            career_id: data.career,
            dedication_id: data.dedication,
            scaled_grade_id: data.scaledGrade,
            contract_type_id: data.contractType
        }
        const response = await PatchTeachersLists(token!, updateTeacherRequest)

        return generateFormMessageFromHttpResponse(form, response)
    }
};