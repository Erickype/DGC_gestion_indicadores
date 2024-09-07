export interface TeachersList {
    created_at: string;
    updated_at: string;
    academic_period_id: number;
    teacher_id: number;
    career_id: number;
    dedication_id: number;
    scaled_grade_id: number;
    contract_type_id: number;
}

export interface CreateTeachersListsRequest {
    academic_period_id: number;
    teacher_id: number;
    career_id: number;
    dedication_id: number;
    scaled_grade_id: number;
    contract_type_id: number;
}

export interface UpdateTeachersListRequest {
    academic_period_id: number;
    teacher_id: number;
    career_id: number;
    dedication_id: number;
    scaled_grade_id: number;
    contract_type_id: number;
}

export interface FilterTeachersListsByAcademicPeriodRequest {
    academic_period_id?: number,
    teacher_identity?: string,
    teacher_name?: string,
    teacher_lastname?: string,
    career?: string,
    dedication?: string,
    scaled_grade?: string,
    contract_type?: string,
    page_size: number,
    page: number
}

export interface TeachersListByAcademicPeriodJoined {
    teacher_id: number,
    teacher_identity: string,
    teacher: string,
    career_id: number,
    career: string,
    dedication_id: number,
    dedication: string,
    scaled_grade_id: number,
    scaled_grade: string,
    contract_type_id: number,
    contract_type: string
}

export interface FilterTeachersListsByAcademicPeriodResponse {
    count: number,
    page_size: number,
    page: number,
    teachers_lists: TeachersListByAcademicPeriodJoined[]
}