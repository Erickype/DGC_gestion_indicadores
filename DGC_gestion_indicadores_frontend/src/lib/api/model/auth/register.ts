import type { User } from "./user"

export interface RegisterRequest {
    username: string,
    email: string,
    password: string
}

export interface RegisterResponse {
    user: User
}

export interface RegisterError{
    error: string
}