export interface ImpactCoefficient {
    id: number
    academic_database_id: number
    created_at: string
    updated_at: string
    name: string
    abbreviation: string
    description: string
}

export interface ImpactCoefficientJoined {
    ID: number
    academic_database_id: number
    academic_database: string
    compensation_factor_id: number
    compensation_factor: string
}
