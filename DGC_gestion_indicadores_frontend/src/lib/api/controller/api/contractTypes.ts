import type { ContractType } from "$lib/api/model/api/contractType";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { getContractTypesRoute } from "$lib/api/routes/api/contractType";
import type { Message } from "$lib/components/combobox/combobox";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function GetContractTypes(token: string) {
    try {
        const response = await fetch(getContractTypesRoute, {
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
        const contractTypes: ContractType[] = await response.json()
        return contractTypes
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}


export async function LoadContractTypesWithComboMessages(token: string) {
    const response = await GetContractTypes(token);
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const contractTypes = response as ContractType[]
    let messages: Message[] = []
    messages = messages.concat(
        contractTypes.map((contractType) => ({
            value: contractType.ID,
            label: contractType.abbreviation,
        }))
    );

    return {
        contractTypes,
        messages
    }
}