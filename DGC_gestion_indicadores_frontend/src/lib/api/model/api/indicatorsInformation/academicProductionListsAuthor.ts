import type { Career } from "../career"

export interface AcademicProductionListsAuthorsCareersJoined {
    author_id: number
    author: string
    careers: Career[]
}

export interface PostAcademicProductionListsAuthorCareersRequest {
    academic_production_list_id: number
    author_id: number
    careers: number[]
}

export interface UpdateAcademicProductionListsAuthorCareersRequest {
    academic_production_list_id: number
    author_id: number
    careers: number[]
}
