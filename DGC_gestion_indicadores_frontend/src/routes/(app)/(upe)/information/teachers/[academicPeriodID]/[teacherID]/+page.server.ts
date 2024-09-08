import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { addTeachersListsDegreeSchema } from "./schema";

export const load: PageServerLoad = async ({ params }) => {
    const academicPeriodID = parseInt(params.teacherID)
    const teacherID = parseInt(params.teacherID)
    if (academicPeriodID > 0 && teacherID > 0) {
        return {
            academicPeriodID: academicPeriodID,
            teacherID: teacherID,
            addTeachersListsDegreeForm: await superValidate(zod(addTeachersListsDegreeSchema))
        }
    }

    error(404, 'Not found');
};