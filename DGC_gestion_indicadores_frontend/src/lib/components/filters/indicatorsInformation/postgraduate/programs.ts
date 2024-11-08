import type { FilterPostgraduateProgramsRequest, FilterPostGraduateProgramsResponse, PostgraduateProgram } from "$lib/api/model/api/indicatorsInformation/postgraduate";
import type { PopoverFilterDataMap } from "$lib/components/table/types";

import type { CommonError } from "$lib/api/model/errors";
import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterDetailedFieldsRequest(page_size: number, page: number): FilterPostgraduateProgramsRequest {
    let filterPostgraduateProgramsRequest: FilterPostgraduateProgramsRequest = {
        name: '',
        start_year: -1,
        end_year: -1,
        page_size: page_size,
        page: page
    };
    return filterPostgraduateProgramsRequest
}

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('name', { label: 'Nombre', value: '' });
    popoverFilterDataMap.set('start_year', { label: 'Año inicio', value: '' });
    popoverFilterDataMap.set('end_year', { label: 'Año fin', value: '' });
    return popoverFilterDataMap
}

export function generateInitialFilterValue(filterPostgraduateProgramsRequest: FilterPostgraduateProgramsRequest): string | undefined {
    let values = [
        filterPostgraduateProgramsRequest.name,
        filterPostgraduateProgramsRequest.start_year?.toString(),
        filterPostgraduateProgramsRequest.end_year?.toString(),
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchFilterPostgraduatePrograms(filterPostgraduateProgramsRequest: FilterPostgraduateProgramsRequest): Promise<FilterPostGraduateProgramsResponse> {
    const url = `/api/indicatorsInformation/postgraduate/programs/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterPostgraduateProgramsRequest)
    });
    if (!response.ok) {
        const errorData = (await response.json()) as CommonError;
        if (response.status === 401) {
            throw goto('/');
        }
        throw errorData;
    }
    return response.json();
}

export async function fetchGetGetPostgraduateProgramByID(programID: string): Promise<PostgraduateProgram> {
    const url = `/api/indicatorsInformation/postgraduate/program/` + programID;
    const response = await fetch(url, {
        method: 'GET',
    });
    if (!response.ok) {
        const errorData = (await response.json()) as CommonError;
        if (response.status === 401) {
            throw goto('/');
        }
        throw errorData;
    }
    return response.json();
}

export async function fetchOnFilterChanged(filter: string, filterPostgraduateProgramsRequest: FilterPostgraduateProgramsRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterPostGraduateProgramsResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterPostgraduateProgramsRequest as any)[key] = filter;
    });

    return fetchFilterPostgraduatePrograms(filterPostgraduateProgramsRequest).then(
        (response: FilterPostGraduateProgramsResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterPostgraduateProgramsRequest as any)[key] = '';
                });
                return fetchFilterPostgraduatePrograms(filterPostgraduateProgramsRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterPostgraduateProgramsRequest: FilterPostgraduateProgramsRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterPostgraduateProgramsRequest, response: FilterPostGraduateProgramsResponse }> {

    let request: FilterPostgraduateProgramsRequest = {
        page: filterPostgraduateProgramsRequest.page,
        page_size: filterPostgraduateProgramsRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterPostgraduateProgramsRequest = request;

    const response = await fetchFilterPostgraduatePrograms(filterPostgraduateProgramsRequest);

    if (response.count === 0) {
        let message = 'No hay datos para el filtro\n';
        popoverFilterDataMap.forEach((item, _) => {
            if (item.value !== '') {
                message += `${item.label}: ${item.value}; `;
            }
        });
        message = message.slice(0, message.length - 2);
        toast.warning(message);

        popoverFilterDataMap.forEach((_, key) => {
            (filterPostgraduateProgramsRequest as any)[key] = '';
        });

        const newResponse = await fetchFilterPostgraduatePrograms(filterPostgraduateProgramsRequest);
        return { request: filterPostgraduateProgramsRequest, response: newResponse };
    }

    return { request: filterPostgraduateProgramsRequest, response };
}