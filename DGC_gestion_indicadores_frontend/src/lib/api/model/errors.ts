import { error } from "@sveltejs/kit";

export interface CommonError {
    StatusCode: number
    Status: string
    Message: string
    Detail: string
}

export function generateErrorFromCommonError(err: CommonError) {
    const errorObject = new Error(err.Message);
    (errorObject as any).Message = err.Message;
    (errorObject as any).StatusCode = err.StatusCode;
    (errorObject as any).Detail = err.Detail;
    return error(err.StatusCode, errorObject);
}