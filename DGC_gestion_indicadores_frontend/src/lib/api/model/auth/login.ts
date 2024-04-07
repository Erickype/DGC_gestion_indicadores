export interface LoginRequest {
    username: string;
    password: string;
}

export interface LoginResponse {
    message: string;
    token: string;
    username: string;
}

export interface LoginError{
    error: string
}
