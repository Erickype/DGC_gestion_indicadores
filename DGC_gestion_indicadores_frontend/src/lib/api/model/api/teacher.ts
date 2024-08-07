export interface Teacher {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    ID: number;
    academic_period_id: number;
    person_id: number;
    career_id: number;
    dedication_id: number;
    scaled_grade_id: number;
}

export interface CreateTeacherRequest {
    academic_period_id: number;
    person_id: number;
    career_id: number;
    dedication_id: number;
    scaled_grade_id: number;
    contract_type_id: number;
}

export interface GetTeachersByAcademicPeriodResponse{
    updated_at: string;
    ID: number;
    academic_period_id: number;
    person_id: number;
    person: string;
    career_id: number;
    career: string;
    dedication_id: number;
    dedication: string;
    scaled_grade_id: number;
    scaled_grade: string;
}

export interface UpdateTeacherRequest {
    ID: number;
    academic_period_id: number;
    person_id: number;
    career_id: number;
    dedication_id: number;
    scaled_grade_id: number;
}
