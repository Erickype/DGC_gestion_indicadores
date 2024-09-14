export interface Author {
    ID: number;
    created_at: string;
    updated_at: string;
    person_id: number;
}

export interface PostAuthorFromPersonRequest {
    identity: string;
    name: string;
    lastname: string;
    email: string;
}