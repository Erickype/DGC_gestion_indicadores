export interface TeachersList{
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