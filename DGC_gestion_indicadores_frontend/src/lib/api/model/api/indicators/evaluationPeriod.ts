export interface IndicatorsEvaluationPeriod {
    created_at: string,
    updated_at: string,
    indicator_type_id: number,
    academic_period_id: number,
    actual_value: number,
    target_value: number
}

export interface IndicatorEvaluationPeriodJoined {
    indicator_type_id: number,
    indicator_type: string,
    academic_period_id: number,
    academic_period: string,
    actual_value: number,
    target_value: number
}