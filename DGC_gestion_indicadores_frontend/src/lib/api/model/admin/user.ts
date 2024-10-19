export interface User {
    created_at: string
    updated_at: string
    deleted_at: string
    ID: number
    role_id: number
    username: string
    email: string
}

export interface UpdateUsersRequest {
    ID: number
    username: string;
    email: string;
    role_id: number;
}