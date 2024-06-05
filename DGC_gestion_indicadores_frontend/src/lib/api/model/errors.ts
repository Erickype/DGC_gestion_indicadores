import { error } from "@sveltejs/kit";

export interface CommonError {
    status_code: number
    status: string
    message: string
    detail: string
}

export function generateErrorFromCommonError(err: CommonError) {
    const errorObject = new Error(err.message);
    (errorObject as any).message = err.message;
    (errorObject as any).status_code = err.status_code;
    (errorObject as any).detail = err.detail;
    return error(err.status_code, errorObject);
}