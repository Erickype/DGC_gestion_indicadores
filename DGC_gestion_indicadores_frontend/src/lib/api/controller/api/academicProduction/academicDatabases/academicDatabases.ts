import { getAcademicDatabasesRoute } from "$lib/api/routes/api/academicProduction/academicDatabases/academicDatabases";
import type { AcademicDatabase } from "$lib/api/model/api/academicProduction/academicDatabases/academicDatabases";

import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { Message } from "$lib/components/combobox/combobox";

export async function GetAcademicDatabases(token: string) {
    try {
        const response = await fetch(getAcademicDatabasesRoute, {
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
        const databases: AcademicDatabase[] = await response.json()
        return databases
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function LoadAcademicDatabasesWithComboMessages(token: string) {
    const response = await GetAcademicDatabases(token);
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const careers = response as AcademicDatabase[];
    let messages: Message[] = []
    messages = messages.concat(
        careers.map((career) => ({
            value: career.ID!,
            label: career.name,
        }))
    );

    return {
        careers,
        messages
    }
}