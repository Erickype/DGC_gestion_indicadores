import { addBooksOrChaptersProductionListsAuthorCareersSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { error, redirect, type Actions } from "@sveltejs/kit";
import { mainDashboarRoute } from "$lib/api/util/paths";
import type { PageServerLoad } from "./$types";

import type { PostAcademicProductionListsAuthorCareersRequest, UpdateAcademicProductionListsAuthorCareersRequest } from "$lib/api/model/api/indicatorsInformation/academicProductionListsAuthor";
import { PostAcademicProductionListsAuthorCareers, UpdateAcademicProductionListsAuthorCareers } from "$lib/api/controller/api/indicatorsInformation/academicProductionListsAuthor";
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
            careersData: await LoadCareersWithComboMessages(token!),
        }
    }

    error(404, 'Not found');
};

export const actions: Actions = {
};