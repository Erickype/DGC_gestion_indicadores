import type { RegisterRequest } from "$lib/api/model/auth/register";
import type { CommonError } from "$lib/api/model/errors";
import { authRouteLogin, authRouteRegister } from "$lib/api/routes/auth/auth";
import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { LoginRequest, LoginResponse } from "../../model/auth/login";

export async function Login(login: LoginRequest): Promise<LoginResponse | CommonError> {
    try {
        const response = await fetch(authRouteLogin, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(login)
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const LoginResponse: LoginResponse = await response.json()
        return LoginResponse;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function Register(registerRequest: RegisterRequest) {
    return await fetch(authRouteRegister, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(registerRequest)
    });
}
