import { getUsersRoute, updateUserRoute } from "$lib/api/routes/admin/user";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";
import type { UpdateUsersRequest, User } from "$lib/api/model/admin/user";

export async function GetUsers(token: string): Promise<User[] | CommonError> {
    try {
        const response = await fetch(getUsersRoute, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const users: User[] = await response.json()
        return users;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function UpdateUser(token: string, request: UpdateUsersRequest): Promise<User | CommonError> {
    try {
        const response = await fetch(updateUserRoute + request.ID, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(request)
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const user: User = await response.json()
        return user;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}