export interface FilterAuthorsRequest {
    identity?: string,
    name?: string,
    lastname?: string,
    email?: string,
    page_size: number,
    page: number
}

export interface FilterAuthorsResponse {
    count: number,
    page_size: number,
    page: number,
    authors: AuthorPerson[]
}

export interface AuthorPerson {
    ID: number,
    person_id: number,
    identity: number,
    name: string,
    lastname: string,
    email: string
}