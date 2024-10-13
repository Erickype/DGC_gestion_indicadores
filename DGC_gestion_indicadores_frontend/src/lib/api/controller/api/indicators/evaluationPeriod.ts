import { getCalculateIndicatorByTypeIDAndEvaluationPeriodRoute, getIndicatorByTypeIDAndEvaluationPeriod } from "$lib/api/routes/api/indicators/evaluationPeriod";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";
import type { IndicatorEvaluationPeriodJoined, IndicatorsEvaluationPeriod } from "$lib/api/model/api/indicators/evaluationPeriod";

export async function GetIndicatorByTypeIDAndEvaluationPeriod(evaluation_period_id: string, indicator_type_id: string, token: string) {
    try {
        const response = await fetch(getIndicatorByTypeIDAndEvaluationPeriod + evaluation_period_id + "/" + indicator_type_id, {
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
        const indicator: IndicatorEvaluationPeriodJoined = await response.json()
        return indicator
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetCalculateIndicatorByTypeIDAndEvaluationPeriod(evaluation_period_id: string, indicator_type_id: string, token: string) {
    try {
        const response = await fetch(getCalculateIndicatorByTypeIDAndEvaluationPeriodRoute + evaluation_period_id + "/" + indicator_type_id, {
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
        const indicator: IndicatorsEvaluationPeriod = await response.json()
        return indicator
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}