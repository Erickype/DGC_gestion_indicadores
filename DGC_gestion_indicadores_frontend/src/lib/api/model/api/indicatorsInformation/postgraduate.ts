export interface PostgraduateProgram {
    created_at: string
    updated_at: string
    ID: number
    name: string
    start_year: number
    end_year: number
    is_active: boolean
}

export interface FilterPostgraduateProgramsRequest {
    name?: string
    start_year?: number
    end_year?: number
    is_active?: boolean
    page_size: number,
    page: number
}

export interface FilterPostGraduateProgramsResponse {
    count: number,
    page_size: number,
    page: number,
    postgraduate_programs: PostgraduateProgram[]
}