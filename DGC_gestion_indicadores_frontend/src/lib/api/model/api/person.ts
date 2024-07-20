export interface Person {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    ID: number;
    identity: string;
    name: string;
    lastname: string;
    email: string;
}

export interface PostPersonRequest {
    identity: string;
    name: string;
    lastname: string;
    email: string;
}

export interface PutPersonRequest {
    ID: number;
    identity: string;
    name: string;
    lastname: string;
    email: string;
}