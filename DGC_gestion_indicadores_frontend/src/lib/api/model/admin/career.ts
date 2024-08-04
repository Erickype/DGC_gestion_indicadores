export interface PostCareerRequest {
    faculty_id: number,
    name: string,
    abbreviation: string,
    description: string
}

export interface PutCareerRequest {
    ID: number,
    faculty_id: number,
    name: string,
    abbreviation: string,
    description: string
}