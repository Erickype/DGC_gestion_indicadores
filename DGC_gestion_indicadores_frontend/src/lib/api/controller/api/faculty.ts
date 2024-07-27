import { getFacultiesRoute } from "$lib/api/routes/api/faculty";
import type { Message } from "$lib/components/combobox/combobox";
import type { Faculty } from "$lib/api/model/api/faculty";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";

export async function GetFaculties(token: string) {
    try {
        const response = await fetch(getFacultiesRoute, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            }
        });

        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const faculties: Faculty[] = await response.json()
        return faculties
    } catch (error) {
        const errorMessage: CommonError = {
            status: "500",
            status_code: 500,
            detail: "Error al solicitar datos",
            message: (error as Error).message
        }
        return errorMessage
    }
}


export async function LoadFacultiesWithComboMessages(token: string) {
    const response = await GetFaculties(token);
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const faculties = response as Faculty[];
    let messages: Message[] = []
    messages = messages.concat(
        faculties.map((faculty) => ({
            value: faculty.ID,
            label: faculty.abbreviation,
        }))
    );

    return {
        faculties,
        messages
    }
}