import { fail, redirect } from '@sveltejs/kit'
import type { Action, Actions, PageServerLoad } from './$types'
import type { RegisterError, RegisterRequest } from '$lib/api/model/auth/register'
import { Register } from '$lib/api/controller/auth/auth'

export const load: PageServerLoad = async ({ locals }) => {
    if (locals.user) {
        throw redirect(302, "/dashboards/evaluation-periods")
    }
}

const register: Action = async ({ request }) => {
    const data = await request.formData()
    const username = data.get('username')
    const email = data.get('email')
    const password = data.get('password')

    if (
        typeof username !== 'string' ||
        typeof password !== 'string' ||
        typeof email !== 'string' ||
        !username ||
        !email ||
        !password
    ) {
        return fail(400, { invalid: true })
    }

    const registerRequest: RegisterRequest = {
        username: username,
        email: email,
        password: password
    }

    const res = await Register(registerRequest)

    if (!res.ok) {
        const err: RegisterError = await res.json()
        return fail(res.status, { error: err.error })
    }

    redirect(303, "/login")
}

export const actions: Actions = { register }
