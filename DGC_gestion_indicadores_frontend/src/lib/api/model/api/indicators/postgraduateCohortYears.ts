export interface IndicatorsPostgraduate {
    created_at: string,
    updated_at: string,
    indicator_type_id: number,
    cohort_list_year: number,
    actual_value: number,
    target_value: number
}

export interface IndicatorsPostgraduateJoined {
    indicator_type_id: number,
    indicator_type: string,
    cohort_list_year: number,
    actual_value: number,
    target_value: number
}