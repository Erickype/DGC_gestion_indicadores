import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { redirect } from "@sveltejs/kit";

import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { PostBooksOrChaptersProductionLists } from "$lib/api/controller/api/indicatorsInformation/booksOrChaptersProductionLists";
import type { BooksOrChaptersProductionList } from "$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionLists";
import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { addBookOrChaptersProductionListSchema, updateBookOrChaptersProductionListSchema } from "./schema";

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
        addBookOrChaptersProductionForm: await superValidate(zod(addBookOrChaptersProductionListSchema)),
        updateBookOrChaptersProductionForm: await superValidate(zod(updateBookOrChaptersProductionListSchema))
    }
};

export const actions: Actions = {
    postBooksOrChaptersProductionLists: async (event) => {
        const form = await superValidate(event, zod(addBookOrChaptersProductionListSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const booksOrChaptersProductionList: BooksOrChaptersProductionList = {
            DOI: data.doi,
            is_chapter: data.is_chapter,
            title: data.title,
            publication_date: new Date(data.publication_date).toISOString(),
            peer_reviewed: data.peer_reviewed,
            detailed_field_id: data.detailed_field_id,
            academic_period_id: data.academic_period_id,
        }

        const response = await PostBooksOrChaptersProductionLists(token!, booksOrChaptersProductionList)

        return generateFormMessageFromHttpResponse(form, response)
    },
};