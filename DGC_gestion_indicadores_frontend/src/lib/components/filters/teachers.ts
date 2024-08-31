import type { FilterTeachersRequest, FilterTeachersResponse } from "$lib/api/model/api/teacher";
import type { PopoverFilterDataMap } from "../table/types";
import type { CommonError } from "$lib/api/model/errors";

import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterTeachersRequest(page_size: number, page: number): FilterTeachersRequest {
    let filterTeachersRequest: FilterTeachersRequest = {
        identity: '',
        name: '',
        lastname: '',
        email: '',
        page_size: page_size,
        page: page
    };
    return filterTeachersRequest
}

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('identity', { label: 'Identificación', value: '' });
    popoverFilterDataMap.set('name', { label: 'Nombre', value: '' });
    popoverFilterDataMap.set('lastname', { label: 'Apellido', value: '' });
    popoverFilterDataMap.set('email', { label: 'Email', value: '' });
    return popoverFilterDataMap
}

export function generateInitialFilterValue(filterTeachersRequest: FilterTeachersRequest): string | undefined {
    let values = [
        filterTeachersRequest.identity,
        filterTeachersRequest.name,
        filterTeachersRequest.lastname,
        filterTeachersRequest.email
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchFilterTeachers(filterTeachersRequest: FilterTeachersRequest): Promise<FilterTeachersResponse> {
    const url = `/api/teachers/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterTeachersRequest)
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

export async function fetchOnFilterChanged(filter: string, filterTeachersRequest: FilterTeachersRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterTeachersResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterTeachersRequest as any)[key] = filter;
    });

    return fetchFilterTeachers(filterTeachersRequest).then(
        (response: FilterTeachersResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterTeachersRequest as any)[key] = '';
                });
                return fetchFilterTeachers(filterTeachersRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterTeachersRequest: FilterTeachersRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterTeachersRequest, response: FilterTeachersResponse }> {

    let request: FilterTeachersRequest = {
        page: filterTeachersRequest.page,
        page_size: filterTeachersRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterTeachersRequest = request;

    const response = await fetchFilterTeachers(filterTeachersRequest);

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
            (filterTeachersRequest as any)[key] = '';
        });

        const newResponse = await fetchFilterTeachers(filterTeachersRequest);
        return { request: filterTeachersRequest, response: newResponse };
    }

    return { request: filterTeachersRequest, response };
}
