export interface FilterSocialProjectListsByAcademicPeriodRequest {
    academic_period_id: number
    career?: string
    name?: string
    page_size: number,
    page: number
}

export interface SocialProjectListByAcademicPeriodJoined {
    academic_period_id: number
    career_id: number
    career: string
    name: string
}

export interface FilterSocialProjectListsByAcademicPeriodResponse {
    count: number,
    page_size: number,
    page: number,
    social_project_lists: SocialProjectListByAcademicPeriodJoined[]
}