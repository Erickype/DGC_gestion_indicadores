export interface TeachersDegree {
    id:number
    created_at:string
    updated_at:string
    degree_level_id: number;
    teacher_id: number;
    name: string;
}

export interface PatchTeachersDegreeByTeachersDegreeIDRequest{
    degree_level_id: number;
    name: string;
}