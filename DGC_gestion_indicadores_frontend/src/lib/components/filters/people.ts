import { goto } from "$app/navigation";
import type { FilterPeopleRequest, FilterPeopleResponse } from "$lib/api/model/api/person";
import type { CommonError } from "$lib/api/model/errors";
import { toast } from "svelte-sonner";
import type { PopoverFilterDataMap } from "../table/types";

export function newFilterPeopleRequest(page_size: number, page: number): FilterPeopleRequest {
    let filterPeopleRequest: FilterPeopleRequest = {
        identity: '',
        name: '',
        lastname: '',
        email: '',
        page_size: page_size,
        page: page
    };
    return filterPeopleRequest
}

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('identity', { label: 'Identificación', value: '' });
    popoverFilterDataMap.set('name', { label: 'Nombre', value: '' });
    popoverFilterDataMap.set('lastname', { label: 'Apellido', value: '' });
    popoverFilterDataMap.set('email', { label: 'Email', value: '' });
    return popoverFilterDataMap
}

export function generateInitialFilterValue(filterPeopleRequest: FilterPeopleRequest): string | undefined {
    let values = [
        filterPeopleRequest.identity,
        filterPeopleRequest.name,
        filterPeopleRequest.lastname,
        filterPeopleRequest.email
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchFilterPeople(filterPeopleRequest: FilterPeopleRequest): Promise<FilterPeopleResponse> {
    const url = `/api/people/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterPeopleRequest)
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

export async function fetchOnFilterChanged(filter: string, filterPeopleRequest: FilterPeopleRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterPeopleResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterPeopleRequest as any)[key] = filter;
    });

    return fetchFilterPeople(filterPeopleRequest).then(
        (response: FilterPeopleResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterPeopleRequest as any)[key] = '';
                });
                return fetchFilterPeople(filterPeopleRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterPeopleRequest: FilterPeopleRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterPeopleRequest, response: FilterPeopleResponse }> {

    let request: FilterPeopleRequest = {
        page: filterPeopleRequest.page,
        page_size: filterPeopleRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterPeopleRequest = request;

    const response = await fetchFilterPeople(filterPeopleRequest);

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
            (filterPeopleRequest as any)[key] = '';
        });

        const newResponse = await fetchFilterPeople(filterPeopleRequest);
        return { request: filterPeopleRequest, response: newResponse };
    }

    return { request: filterPeopleRequest, response };
}
