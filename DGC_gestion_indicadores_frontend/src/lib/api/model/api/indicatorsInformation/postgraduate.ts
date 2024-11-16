export interface PostgraduateProgram {
    created_at?: string
    updated_at?: string
    ID?: number
    name: string
    start_year: number
    end_year: number
    is_active: boolean
}

export interface FilterPostgraduateProgramsRequest {
    name?: string
    start_year?: string
    end_year?: string
    is_active?: boolean
    page_size: number,
    page: number
}

export interface GetPostgraduateProgramMissingCohortYearsByProgramIDResponse {
    years: number[]
}

export interface FilterPostGraduateProgramsResponse {
    count: number,
    page_size: number,
    page: number,
    postgraduate_programs: PostgraduateProgram[]
}

export interface PostgraduateCohortList {
    created_at?: string
    updated_at?: string
    postgraduate_program_id: number
    year: number
    name: string
    graduated_students: number
    total_students: number
}

export interface CohortList {
    created_at?: string
    updated_at?: string
    year: number
}

export interface FilterCohortListsRequest {
    start_year: number
    end_year: number
    page_size: number
    page: number
}

export interface FilterCohortListsResponse {
    count: number,
    page_size: number,
    page: number,
    cohort_lists: CohortList[]
}