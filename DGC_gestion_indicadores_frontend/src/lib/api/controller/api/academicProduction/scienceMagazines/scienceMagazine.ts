import type { ScienceMagazine } from "$lib/api/model/api/academicProduction/scienceMagazines/scienceMagazine";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { getScienceMagazinesRoute } from "$lib/api/routes/api/academicProduction/scienceMagazines/scienceMagazine";
import type { Message } from "$lib/components/combobox/combobox";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function GetScienceMagazines(token: string) {
    try {
        const response = await fetch(getScienceMagazinesRoute, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            }
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const magazines: ScienceMagazine[] = await response.json()
        return magazines
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}


export async function LoadGetScienceMagazinesWithComboMessages(token: string) {
    const response = await GetScienceMagazines(token);
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const magazines: ScienceMagazine[] = response as ScienceMagazine[];
    let messages: Message[] = []
    messages = messages.concat(
        magazines.map((magazine) => ({
            value: magazine.ID,
            label: magazine.name,
        }))
    );

    return {
        magazines,
        messages
    }
}