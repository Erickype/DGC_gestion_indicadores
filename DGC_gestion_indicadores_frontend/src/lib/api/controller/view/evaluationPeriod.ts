import { getEvaluationPeriodsRoute } from "$lib/api/routes/view/evaluationPeriod";
import type { EvaluationPeriod } from "$lib/api/model/view/evaluationPeriod";

import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import type { Message } from "$lib/components/combobox/combobox";

export async function GetEvaluationPeriods() {
    try {
        const response = await fetch(getEvaluationPeriodsRoute, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const periods: EvaluationPeriod[] = await response.json()
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


export async function LoadEvaluationPeriodsWithComboMessages() {
    const response = await GetEvaluationPeriods();
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const periods = response as EvaluationPeriod[]
    let messages: Message[] = []
    messages = messages.concat(
        periods.map((period) => ({
            value: period.ID,
            label: period.abbreviation,
        }))
    );

    return {
        periods,
        messages
    }
}