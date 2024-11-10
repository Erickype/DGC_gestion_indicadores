import { mainDashboarRoute } from "$lib/api/util/paths";
import type { PageServerLoad, Actions } from "./$types";
import { redirect } from "@sveltejs/kit"

import { addPostgraduateCohortListSchema, updatePostgraduateCohortListSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { filterPostgraduateProgramAuxSchema, generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { PostPostgraduateCohortList, UpdatePostgraduateCohortList } from "$lib/api/controller/api/indicatorsInformation/postgraduate";
import type { PostgraduateCohortList } from "$lib/api/model/api/indicatorsInformation/postgraduate";

export const load: PageServerLoad = async ({ cookies, locals }) => {
    const token = cookies.get("AuthorizationToken")
    if (!locals.user) {
        throw redirect(302, "/login")
    }
    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    return {
        filterPostgraduateProgramAuxForm: await superValidate(zod(filterPostgraduateProgramAuxSchema)),
        addPostgraduateCohortListForm: await superValidate(zod(addPostgraduateCohortListSchema)),
        updatePostgraduateCohortListForm: await superValidate(zod(updatePostgraduateCohortListSchema)),
    }
};

export const actions: Actions = {
    postPostgraduateCohortList: async (event) => {
        const form = await superValidate(event, zod(addPostgraduateCohortListSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const postgraduateCohortList: PostgraduateCohortList = {
            postgraduate_program_id: data.postgraduate_program_id,
            year: data.year,
            name: data.name,
            graduated_students: data.graduated_students,
            total_students: data.total_students,
        }

        const response = await PostPostgraduateCohortList(token!, postgraduateCohortList)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updatePostgraduateCohortList: async (event) => {
        const form = await superValidate(event, zod(updatePostgraduateCohortListSchema))
        const token = event.cookies.get("AuthorizationToken")

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const data = form.data
        const postgraduateCohortList: PostgraduateCohortList = {
            postgraduate_program_id: data.postgraduate_program_id,
            year: data.year,
            name: data.name,
            graduated_students: data.graduated_students,
            total_students: data.total_students,
        }

        const response = await UpdatePostgraduateCohortList(token!, postgraduateCohortList)

        return generateFormMessageFromHttpResponse(form, response)
    }
};