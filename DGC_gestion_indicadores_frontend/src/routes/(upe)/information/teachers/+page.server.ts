import { redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { LoadPeopleWithComboMessages } from "$lib/api/controller/api/person";
import { addTeacherSchema } from "./scheme";
import { fail, superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

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
        addTeacherForm: await superValidate(zod(addTeacherSchema)),
    }
};

export const actions: Actions = {
    addTeacher: async (event) => {
        const form = await superValidate(event, zod(addTeacherSchema))

        if (!form.valid) {
            return fail(400, {
                form,
            })
        }

        return{
            form
        }
    }
};