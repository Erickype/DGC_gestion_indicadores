import type { RegisterRequest } from "$lib/api/model/auth/register";
import { authRouteLogin, authRouteRegister } from "../../constants";
import type { LoginRequest } from "../../model/auth/login";

export async function Login(login: LoginRequest) {
    return await fetch(authRouteLogin, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(login)
    });
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
