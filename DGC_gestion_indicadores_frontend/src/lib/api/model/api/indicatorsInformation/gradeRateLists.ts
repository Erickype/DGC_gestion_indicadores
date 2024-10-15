export interface GradeRateList {
    created_at: string
    updated_at: string
    academic_period_id: number
    career_id: number
    count_graduated_students: number
    count_court_students: number
    count_admitted_matriculated_students: number
    count_admitted_students: number
}

export interface GradeRateListJoined {
    academic_period_id: number
    academic_period: string
    career_id: number
    career: string
    count_graduated_students: number
    count_court_students: number
    count_admitted_matriculated_students: number
    count_admitted_students: number
}