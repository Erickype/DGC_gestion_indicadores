import { error } from "@sveltejs/kit";

export interface CommonError {
    status_code: number
    status: string
    message: string
    detail: string
}

export function generateErrorFromCommonError(err: CommonError) {
    const body: App.Error = {
        message: `${err.detail}: ${err.message}`
    }
    return error(err.status_code, body);
}