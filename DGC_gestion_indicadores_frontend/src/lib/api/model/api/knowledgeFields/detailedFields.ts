export interface FilterDetailedFieldRequest {
    wide_field?: string,
    specific_field?: string,
    detailed_field?: string,
    page_size: number,
    page: number
}

export interface FilterDetailedFieldResponse {
    count: number,
    page_size: number,
    page: number,
    detailed_fields: DetailedFilterJoined[]
}

export interface DetailedFilterJoined {
    wide_field_id: number,
    wide_field: string,
    specific_field_id: number,
    specific_field: string,
    detailed_field_id: number,
    detailed_field: string,
}