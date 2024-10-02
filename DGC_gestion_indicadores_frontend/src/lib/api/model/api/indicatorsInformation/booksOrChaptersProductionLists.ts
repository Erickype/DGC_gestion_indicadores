export interface BooksOrChaptersProductionList {
    created_at?: string
    updated_at?: string
    ID?: number
    DOI: string
    is_chapter: boolean
    title: string
    publication_date: string
    peer_reviewed: boolean
    academic_period_id: number
    detailed_field_id: number
}

export interface FilterBooksOrChaptersProductionListsByAcademicPeriodRequest {
    DOI?: string
    is_chapter?: boolean
    title?: string
    publication_date?: string
    peer_reviewed?: boolean
    detailed_field?: string
    academic_period_id: number
    page_size: number,
    page: number
}

export interface BooksOrChaptersProductionListsByAcademicPeriodJoined {
    ID: number
    DOI: string
    is_chapter: boolean
    title: string
    publication_date: string
    peer_reviewed: boolean
    detailed_field_id: number
    detailed_field: string
    academic_period_id: number
}

export interface FilterBooksOrChaptersProductionListsByAcademicPeriodResponse {
    count: number,
    page_size: number,
    page: number,
    books_or_chapters_production_lists: BooksOrChaptersProductionListsByAcademicPeriodJoined[]
}