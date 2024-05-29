import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import type { AcademicPeriod } from "$lib/api/model/view/academicPeriod";
import { getAcademicPeriodsRoute } from "$lib/api/routes/view/academicPeriod";
import type { Message } from "$lib/components/combobox/combobox";

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
            Status: "500",
            StatusCode: 500,
            Detail: "Error al solicitar datos",
            Message: (error as Error).message
        }
        return errorMessage
    }
}


export async function LoadAcademicPeriodsWithComboMessages() {
    const response = await GetAcademicPeriods();
    if ((response as CommonError).Status) {
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