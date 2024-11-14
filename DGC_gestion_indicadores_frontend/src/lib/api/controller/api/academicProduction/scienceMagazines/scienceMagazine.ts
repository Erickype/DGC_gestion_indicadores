import type { FilterScienceMagazinesRequest, FilterScienceMagazinesResponse, ScienceMagazine, ScienceMagazineJoined } from "$lib/api/model/api/academicProduction/scienceMagazines/scienceMagazine";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { getScienceMagazineFilterJoinedByScienceMagazineIDRoute, getScienceMagazinesRoute, postFilterScienceMagazinesRoute, postScienceMagazineRoute, updateScienceMagazineRoute } from "$lib/api/routes/api/academicProduction/scienceMagazines/scienceMagazine";
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

export async function PostFilterScienceMagazines(token: string, filterScienceMagazinesRequest: FilterScienceMagazinesRequest): Promise<FilterScienceMagazinesResponse | CommonError> {
    try {
        const response = await fetch(postFilterScienceMagazinesRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(filterScienceMagazinesRequest)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const filterScienceMagazinesResponse: FilterScienceMagazinesResponse = await response.json()
        return filterScienceMagazinesResponse

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetScienceMagazineFilterJoinedByScienceMagazineID(token: string, detailedFieldID: string): Promise<ScienceMagazineJoined | CommonError> {
    try {
        const response = await fetch(getScienceMagazineFilterJoinedByScienceMagazineIDRoute + detailedFieldID, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const detailedField: ScienceMagazineJoined = await response.json()
        return detailedField

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PostScienceMagazine(token: string, request: ScienceMagazine) {
    try {
        const response = await fetch(postScienceMagazineRoute, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(request)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const scienceMagazine: ScienceMagazine = await response.json()
        return scienceMagazine;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function UpdateScienceMagazine(token: string, request: ScienceMagazine) {
    try {
        const response = await fetch(updateScienceMagazineRoute + request.ID, {
            method: "PUT",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(request)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const scienceMagazine: ScienceMagazine = await response.json()
        return scienceMagazine;
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