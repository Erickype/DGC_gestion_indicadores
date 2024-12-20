import { getCalculateIndicatorByTypeIDAndCohortYearRoute, getCalculateIndicatorsByPostgraduateCohortYearRoute, getIndicatorsByCohortYearRoute } from "$lib/api/routes/api/indicators/postgraduateCohortYears";
import type { IndicatorsPostgraduate, IndicatorsPostgraduateJoined } from "$lib/api/model/api/indicators/postgraduateCohortYears";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function GetCalculateIndicatorsByPostgraduateCohortYear(postgraduate_cohort_year: string, token: string) {
    try {
        const response = await fetch(getCalculateIndicatorsByPostgraduateCohortYearRoute + postgraduate_cohort_year, {
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
        const indicator: IndicatorsPostgraduateJoined = await response.json()
        return indicator
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetIndicatorsByCohortYear(evaluationPeriodID: string, token: string) {
    try {
        const response = await fetch(getIndicatorsByCohortYearRoute + evaluationPeriodID, {
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
        const indicators: IndicatorsPostgraduateJoined[] = await response.json()
        return indicators
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetCalculateIndicatorByTypeIDAndCohortYear(postgraduate_cohort_year: string, indicator_type_id: string, token: string) {
    try {
        const response = await fetch(getCalculateIndicatorByTypeIDAndCohortYearRoute + postgraduate_cohort_year + "/" + indicator_type_id, {
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
        const indicator: IndicatorsPostgraduate = await response.json()
        return indicator
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}