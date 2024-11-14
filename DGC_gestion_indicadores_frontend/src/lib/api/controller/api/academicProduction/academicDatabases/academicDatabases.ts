import { getAcademicDatabasesRoute } from "$lib/api/routes/api/academicProduction/academicDatabases/academicDatabases";
import type { AcademicDatabase } from "$lib/api/model/api/academicProduction/academicDatabases/academicDatabases";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

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