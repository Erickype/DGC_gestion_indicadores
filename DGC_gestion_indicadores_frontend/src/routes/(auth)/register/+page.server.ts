import { fail, redirect } from '@sveltejs/kit'
import type { Action, Actions, PageServerLoad } from './$types'
import type { RegisterError, RegisterRequest } from '$lib/api/model/auth/register'
import { Register } from '$lib/api/controller/auth/auth'
import { registerSchema } from './schema'
import { message, superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'

export const load: PageServerLoad = async ({ locals }) => {
    if (locals.user) {
        throw redirect(302, "/dashboards/evaluation-periods")
    }
    return {
        registerForm: await superValidate(zod(registerSchema)),
    }
}

const register: Action = async (event) => {
    const form = await superValidate(event, zod(registerSchema));

    if (!form.valid) {
        return fail(400, {
            form,
        });
    };

    const data = form.data
    const username = data.username
    const email = data.email
    const password = data.password

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
        return message(form, "Register error: " + err.error, { status: 400 })
    }

    redirect(303, "/login")
}

export const actions: Actions = { register }
