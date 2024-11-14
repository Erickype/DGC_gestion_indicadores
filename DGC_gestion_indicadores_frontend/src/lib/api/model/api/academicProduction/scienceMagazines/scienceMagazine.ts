export interface ScienceMagazine {
    created_at?: string
    updated_at?: string
    ID?: number
    academic_database_id: number
    name: string
    abbreviation: string
    description: string
}

export interface FilterScienceMagazinesRequest {
    academic_database?: string
    science_magazine?: string
    magazine_abbreviation?: string
    magazine_description?: string
    page_size: number
    page: number
}

export interface ScienceMagazineJoined {
    academic_database_id: number
    academic_database: string
    science_magazine_id: number
    science_magazine: string
    science_magazine_abbreviation: string
    science_magazine_description: string
}

export interface FilterScienceMagazinesResponse {
    count: number,
    page_size: number,
    page: number,
    science_magazines: ScienceMagazineJoined[]
}