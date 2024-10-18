import type { IndicatorAcademicPeriodJoined, IndicatorsAcademicPeriod } from "$lib/api/model/api/indicators/academicPeriods";
import type { CommonError } from "$lib/api/model/errors";
import { getCalculateIndicatorByTypeIDAndAcademicPeriodRoute, getCalculateIndicatorsByAndAcademicPeriodRoute, getIndicatorsByAcademicPeriodIDRoute } from "$lib/api/routes/api/indicators/academicPeriods";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function GetIndicatorsByAcademicPeriod(academic_period_id: string, token: string) {
    try {
        const response = await fetch(getIndicatorsByAcademicPeriodIDRoute + academic_period_id, {
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
        const indicators: IndicatorAcademicPeriodJoined[] = await response.json()
        return indicators
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetCalculateIndicatorByTypeIDAndAcademicPeriod(academic_period_id: string, indicator_type_id: string, token: string) {
    try {
        const response = await fetch(getCalculateIndicatorByTypeIDAndAcademicPeriodRoute + academic_period_id + "/" + indicator_type_id, {
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
        const indicators: IndicatorsAcademicPeriod[] = await response.json()
        return indicators
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetCalculateIndicatorsByAcademicPeriod(academic_period_id: string, token: string) {
    try {
        const response = await fetch(getCalculateIndicatorsByAndAcademicPeriodRoute + academic_period_id, {
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
        const indicators: IndicatorsAcademicPeriod[] = await response.json()
        return indicators
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}