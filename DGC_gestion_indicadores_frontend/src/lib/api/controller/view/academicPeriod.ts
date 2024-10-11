import type { AcademicPeriod, FilterAcademicPeriodsRequest, FilterAcademicPeriodsResponse } from "$lib/api/model/view/academicPeriod";
import { getAcademicPeriodsRoute, postFilterAcademicPeriodsRoute } from "$lib/api/routes/view/academicPeriod";

import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import type { Message } from "$lib/components/combobox/combobox";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function GetAcademicPeriods(): Promise<AcademicPeriod[] | CommonError> {
    try {
        const response = await fetch(getAcademicPeriodsRoute, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const periods: AcademicPeriod[] = await response.json()
        return periods
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

export async function PostFilterAcademicPeriods(token: string, filterAcademicPeriodsRequest: FilterAcademicPeriodsRequest) {
    try {
        const response = await fetch(postFilterAcademicPeriodsRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(filterAcademicPeriodsRequest)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const filterAcademicPeriodsResponse: FilterAcademicPeriodsResponse = await response.json()
        return filterAcademicPeriodsResponse

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function LoadAcademicPeriodsWithComboMessages() {
    const response = await GetAcademicPeriods();
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const periods = response as AcademicPeriod[]
    let messages: Message[] = []
    messages = messages.concat(
        periods.map((period) => ({
            value: period.ID,
            label: period.name,
        }))
    );

    return {
        periods,
        messages
    }
}