import { addBooksOrChaptersProductionListsAuthorCareersSchema, updateBooksOrChaptersProductionListsAuthorCareersSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { error, redirect, type Actions } from "@sveltejs/kit";
import { mainDashboarRoute } from "$lib/api/util/paths";
import type { PageServerLoad } from "./$types";

import type { PostBooksOrChaptersProductionListsAuthorCareersRequest, UpdateBooksOrChaptersProductionListsAuthorCareersRequest } from "$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionListsAuthor";
import { PostBooksOrChaptersProductionListsAuthorCareers, UpdateBooksOrChaptersProductionListsAuthorCareers } from "$lib/api/controller/api/indicatorsInformation/booksOrChaptersProductionListsAuthor";
import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { LoadCareersWithComboMessages } from "$lib/api/controller/api/career";

export const load: PageServerLoad = async ({ locals, cookies, params }) => {
    const token = cookies.get("AuthorizationToken")

    if (!locals.user) {
        throw redirect(302, "/login")
    }

    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    const booksOrChapersProductionListID = parseInt(params.booksOrChapersProductionListID)
    if (booksOrChapersProductionListID > 0) {
        return {
            booksOrChapersProductionListID: booksOrChapersProductionListID,
            addBooksOrChaptersProductionListsAuthorCareersForm: await superValidate(zod(addBooksOrChaptersProductionListsAuthorCareersSchema)),
            updateBooksOrChaptersProductionListsAuthorCareersForm: await superValidate(zod(updateBooksOrChaptersProductionListsAuthorCareersSchema)),
            careersData: await LoadCareersWithComboMessages(token!),
        }
    }

    error(404, 'Not found');
};

export const actions: Actions = {
    postBooksOrChaptersProductionListsAuthorCareers: async (event) => {
        const form = await superValidate(event, zod(addBooksOrChaptersProductionListsAuthorCareersSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const academicProductionListsAuthorCareersRequest: PostBooksOrChaptersProductionListsAuthorCareersRequest = {
            books_or_chapters_production_list_id: data.booksOrChaptersProductionListID,
            author_id: data.authorID,
            careers: data.careers
        }

        const response = await PostBooksOrChaptersProductionListsAuthorCareers(token!, academicProductionListsAuthorCareersRequest)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updateBooksOrChaptersProductionListsAuthorCareers: async (event) => {
        const form = await superValidate(event, zod(updateBooksOrChaptersProductionListsAuthorCareersSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const updateBooksOrChaptersProductionListsAuthorCareersRequest: UpdateBooksOrChaptersProductionListsAuthorCareersRequest = {
            books_or_chapters_production_list_id: data.booksOrChaptersProductionListID,
            author_id: data.authorID,
            careers: data.careers
        }

        const response = await UpdateBooksOrChaptersProductionListsAuthorCareers(token!, updateBooksOrChaptersProductionListsAuthorCareersRequest)

        return generateFormMessageFromHttpResponse(form, response)
    },
};