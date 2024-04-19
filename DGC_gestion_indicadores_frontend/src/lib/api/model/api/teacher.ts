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
}