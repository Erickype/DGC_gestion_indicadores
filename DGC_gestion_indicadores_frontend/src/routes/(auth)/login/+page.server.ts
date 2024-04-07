import { fail, redirect } from '@sveltejs/kit'
import type { Action, Actions, PageServerLoad } from './$types'
import type { LoginError, LoginRequest, LoginResponse } from '$lib/api/model/auth/login'
import { Login } from '$lib/api/controller/auth/auth'

export const load: PageServerLoad = async ({ locals }) => {
    if (locals.user) {
        throw redirect(302, "/")
    }
}

const login: Action = async ({ cookies, request }) => {
    const data = await request.formData()
    const username = data.get('username')
    const password = data.get('password')

    if (
        typeof username !== 'string' ||
        typeof password !== 'string' ||
        !username ||
        !password
    ) {
        return fail(400, { invalid: true })
    }

    const LoginRequest: LoginRequest = {
        username: username,
        password: password
    }

    const res = await Login(LoginRequest)

    if (!res.ok) {
        const err: LoginError = await res.json()
        return fail(res.status, { error: err.error })
    }

    const authenticatedUser: LoginResponse = await res.json()

    cookies.set('AuthorizationToken', `Bearer ${authenticatedUser.token}`, {
        httpOnly: true,
        path: '/',
        secure: true,
        sameSite: 'strict',
        maxAge: 60 * 60 * 30 // 30 minutes
    });

    redirect(303, "/")
}

export const actions: Actions = { login }
