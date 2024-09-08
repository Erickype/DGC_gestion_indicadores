export interface AddDegreeAndTeachersListsDegreeRequest {
    academic_period_id: number;
    teacher_id: number;
    degree_level_id: number;
    name: string;
}

export interface GetTeachersListsDegreesJoinedResponse {
    academic_period_id: number;
    teacher_id: number;
    teachers_degree_id: number;
    degree_level_id: number;
    degree_level_name: string;
    name: string;
}