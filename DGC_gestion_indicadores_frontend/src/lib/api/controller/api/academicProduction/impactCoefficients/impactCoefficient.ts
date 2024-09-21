import type { ImpactCoefficientJoined } from "$lib/api/model/api/academicProduction/impactCoefficients/impactCoefficient";
import { getImpactCoefficientsRoute } from "$lib/api/routes/api/academicProduction/impactCoefficients/impactCoefficient";

import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import type { Message } from "$lib/components/combobox/combobox";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function GetImpactCoefficients(token: string) {
    try {
        const response = await fetch(getImpactCoefficientsRoute, {
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
        const coefficients: ImpactCoefficientJoined[] = await response.json()
        return coefficients
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}


export async function LoadGetImpactCoefficientssWithComboMessages(token: string) {
    const response = await GetImpactCoefficients(token);
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const coefficients: ImpactCoefficientJoined[] = response as ImpactCoefficientJoined[];
    let messages: Message[] = []
    messages = messages.concat(
        coefficients.map((coefficient) => ({
            value: coefficient.ID,
            label: `${coefficient.academic_database} ${coefficient.compensation_factor}`,
        }))
    );

    return {
        coefficients,
        messages
    }
}