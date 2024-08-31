export interface Teacher {
    ID: number;
    created_at: string;
    updated_at: string;
    person_id: number;
}

export interface CreateTeacherRequest {
    academic_period_id: number;
    person_id: number;
    career_id: number;
    dedication_id: number;
    scaled_grade_id: number;
    contract_type_id: number;
}

export interface GetTeachersByAcademicPeriodResponse {
    updated_at: string;
    academic_period_id: number;
    person_id: number;
    person: string;
    career_id: number;
    career: string;
    dedication_id: number;
    dedication: string;
    scaled_grade_id: number;
    scaled_grade: string;
    contract_type_id: number;
    contract_type: string;
}

export interface AddTeacherRequest {
    person_id: number
}

export interface UpdateTeacherRequest {
    ID: number;
    academic_period_id: number;
    person_id: number;
    career_id: number;
    dedication_id: number;
    scaled_grade_id: number;
}

export interface FilterTeachersRequest {
    identity?: string,
    name?: string,
    lastname?: string,
    email?: string,
    page_size: number,
    page: number
}

export interface FilterTeachersResponse {
    count: number,
    page_size: number,
    page: number,
    teachers: TeacherPerson[]
}

export interface TeacherPerson {
    ID: number,
    person_id: number,
    identity: number,
    name: string,
    lastname: string,
    email: string
}