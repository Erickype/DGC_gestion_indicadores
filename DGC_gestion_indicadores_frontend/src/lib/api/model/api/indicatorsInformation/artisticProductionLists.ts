export interface ArtisticProductionList {
    created_at?: string
    updated_at?: string
    academic_period_id: number
    international_artistic_work: number
    national_artistic_work: number
    intellectual_property: number
}

export interface ArtisticProductionListJoined {
    academic_period_id: number
    academic_period: string
    international_artistic_work: number
    national_artistic_work: number
    intellectual_property: number
}