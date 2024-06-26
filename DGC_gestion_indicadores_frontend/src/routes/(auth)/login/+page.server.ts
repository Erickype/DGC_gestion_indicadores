import { fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types.js'
import type { LoginError, LoginRequest, LoginResponse } from '$lib/api/model/auth/login'
import { Login } from '$lib/api/controller/auth/auth'
import { loginSchema } from './schema'
import { message, superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import type { Token } from '$lib/api/model/auth/token.js'
import { jwtDecode } from 'jwt-decode'

export const load: PageServerLoad = async ({ locals }) => {
    if (locals.user) {
        throw redirect(302, "/")
    }

    return {
        loginForm: await superValidate(zod(loginSchema)),
    }
}

export const actions: Actions = {
    login: async (event) => {
        const form = await superValidate(event, zod(loginSchema));

        if (!form.valid) {
            return fail(400, {
                form,
            });
        };

        const data = form.data
        const username = data.username
        const password = data.password

        const LoginRequest: LoginRequest = {
            username: username,
            password: password
        }

        const res = await Login(LoginRequest)

        if (!res.ok) {
            const err: LoginError = await res.json()
            return message(form, { success: false, error: err.error }, {
                status: 400
            })
        }

        const authenticatedUser: LoginResponse = await res.json()
        const tokenDecoded: Token = jwtDecode(authenticatedUser.token)

        event.cookies.set('AuthorizationToken', `Bearer ${authenticatedUser.token}`, {
            httpOnly: true,
            path: '/',
            secure: true,
            sameSite: 'strict',
            maxAge: tokenDecoded.eat - tokenDecoded.iat
        });

        redirect(302, "/")
    }
};
