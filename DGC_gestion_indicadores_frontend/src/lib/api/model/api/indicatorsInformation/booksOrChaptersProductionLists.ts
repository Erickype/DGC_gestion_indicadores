export interface BooksOrChaptersProductionList {
    created_at?: string
    updated_at?: string
    ID?: number
    DOI: string
    is_chapter:boolean
    title:string
    publication_date: string
    peer_reviewed: boolean
    academic_period_id: number
    detailed_field_id: number
}