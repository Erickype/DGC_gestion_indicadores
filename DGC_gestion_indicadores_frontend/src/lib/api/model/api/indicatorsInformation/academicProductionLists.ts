export interface AcademicProductionList {
    created_at?: string
    updated_at?: string
    ID?: number
    DOI: string
    publication_date: string
    publication_name: string
    academic_period_id: number
    detailed_field_id: number
    science_magazine_id: number
    impact_coefficient_id: number
    intercultural_component: boolean
}

export interface FilterAcademicProductionListsByAcademicPeriodRequest {
    academic_period_id?: number,
    doi?: string
    publication_date?: string
    publication_name?: string
    detailed_field?: string
    science_magazine?: string
    impact_coefficient?: string
    career?: string
    intercultural_component?: boolean
    page_size: number,
    page: number
}

export interface AcademicProductionListByAcademicPeriodJoined {
    doi: string
    publication_date: string
    publication_name: string
    publication_type_id: number
    detailed_field: string
    science_magazine_id: number
    science_magazine: string
    impact_coefficient_id: number
    impact_coefficient: string
    career_id: number,
    career: string,
    intercultural_component: boolean
}

export interface FilterAcademicProductionListsByAcademicPeriodResponse {
    count: number,
    page_size: number,
    page: number,
    academic_production_lists: AcademicProductionListByAcademicPeriodJoined[]
}