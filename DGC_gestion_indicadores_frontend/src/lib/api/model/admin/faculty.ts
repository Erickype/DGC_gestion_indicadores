export interface PostFacultyRequest {
    name: string,
    abbreviation: string,
    description: string
}

export interface PutFacultyRequest {
    ID: number,
    name: string,
    abbreviation: string,
    description: string
}