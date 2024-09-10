import { error, redirect, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

import { addDegreeAndTeachersListsDegreeSchema, addDegreeNotAssignedSchema, updateDegreeAndTeachersListsDegreeSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { AddDegreeAndTeachersListsDegree, LoadDegreesNotAssignedWithComboMessages } from "$lib/api/controller/api/indicatorsInformation/teachersListsDegree";
import type { AddDegreeAndTeachersListsDegreeRequest } from "$lib/api/model/api/indicatorsInformation/teachersListsDegree";
import type { PatchTeachersDegreeByTeachersDegreeIDRequest } from "$lib/api/model/api/teachersDegree";
import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { PatchTeachersDegreeByTeachersDegreeID } from "$lib/api/controller/api/teachersDegree";
import { LoadDegreeLevelsWithComboMessages } from "$lib/api/controller/api/degreeLevels";
import { mainDashboarRoute } from "$lib/api/util/paths";

export const load: PageServerLoad = async ({ params, locals, cookies }) => {
    const token = cookies.get("AuthorizationToken")

    if (!locals.user) {
        throw redirect(302, "/login")
    }

    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    const academicPeriodID = parseInt(params.academicPeriodID)
    const teacherID = parseInt(params.teacherID)
    if (academicPeriodID > 0 && teacherID > 0) {
        return {
            academicPeriodID: academicPeriodID,
            teacherID: teacherID,
            addDegreeAndTeachersListsDegreeForm: await superValidate(zod(addDegreeAndTeachersListsDegreeSchema)),
            updateDegreeAndTeachersListsDegreeForm: await superValidate(zod(updateDegreeAndTeachersListsDegreeSchema)),
            addDegreeNotAssignedForm: await superValidate(zod(addDegreeNotAssignedSchema)),
            degreeLevelsData: await LoadDegreeLevelsWithComboMessages(token!),
            notAssigedDegrees: await LoadDegreesNotAssignedWithComboMessages(academicPeriodID.toString(), teacherID.toString(), token!)
        }
    }

    error(404, 'Not found');
};

export const actions: Actions = {
    addDegreeAndTeachersListsDegree: async (event) => {
        const form = await superValidate(event, zod(addDegreeAndTeachersListsDegreeSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const addDegreeAndTeachersListsDegreeRequest: AddDegreeAndTeachersListsDegreeRequest = {
            academic_period_id: data.academicPeriodID,
            teacher_id: data.teacherID,
            degree_level_id: data.degreeLevelID,
            name: data.name
        }

        const response = await AddDegreeAndTeachersListsDegree(token!, addDegreeAndTeachersListsDegreeRequest)

        return generateFormMessageFromHttpResponse(form, response)
    },

    patchTeachersDegree: async (event) => {
        const form = await superValidate(event, zod(updateDegreeAndTeachersListsDegreeSchema))
        const token = event.cookies.get("AuthorizationToken")

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const data = form.data
        const updateTeacherRequest: PatchTeachersDegreeByTeachersDegreeIDRequest = {
            degree_level_id: data.degreeLevelID,
            name: data.name
        }
        const response = await PatchTeachersDegreeByTeachersDegreeID(token!, updateTeacherRequest, data.teachersDegreeID.toString())

        return generateFormMessageFromHttpResponse(form, response)
    }
};