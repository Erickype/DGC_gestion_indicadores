export interface User {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    ID: number;
    role_id: number;
    username: string;
    email: string;
}
