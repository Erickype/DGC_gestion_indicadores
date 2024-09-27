import type { AuthorPerson, FilterAuthorsRequest, FilterAuthorsResponse } from "$lib/api/model/api/academicProduction/authors/authorsFilter";
import type { PopoverFilterDataMap } from "$lib/components/table/types";
import type { CommonError } from "$lib/api/model/errors";

import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterAuthorsRequest(page_size: number, page: number): FilterAuthorsRequest {
    let filterAuthorsRequest: FilterAuthorsRequest = {
        identity: '',
        name: '',
        lastname: '',
        email: '',
        page_size: page_size,
        page: page
    };
    return filterAuthorsRequest
}

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('identity', { label: 'Identificación', value: '' });
    popoverFilterDataMap.set('name', { label: 'Nombre', value: '' });
    popoverFilterDataMap.set('lastname', { label: 'Apellido', value: '' });
    popoverFilterDataMap.set('email', { label: 'Email', value: '' });
    return popoverFilterDataMap
}

export function generateInitialFilterValue(filterAuthorsRequest: FilterAuthorsRequest): string | undefined {
    let values = [
        filterAuthorsRequest.identity,
        filterAuthorsRequest.name,
        filterAuthorsRequest.lastname,
        filterAuthorsRequest.email
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchFilterAuthors(filterAuthorsRequest: FilterAuthorsRequest): Promise<FilterAuthorsResponse> {
    const url = `/api/authors/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterAuthorsRequest)
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

export async function fetchAuthorPersonJoinedByAuthorID(authorID: string): Promise<AuthorPerson> {
    const url = `/api/author/joined/` + authorID;
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

export async function fetchOnFilterChanged(filter: string, filterAuthorsRequest: FilterAuthorsRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterAuthorsResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterAuthorsRequest as any)[key] = filter;
    });

    return fetchFilterAuthors(filterAuthorsRequest).then(
        (response: FilterAuthorsResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterAuthorsRequest as any)[key] = '';
                });
                return fetchFilterAuthors(filterAuthorsRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterAuthorsRequest: FilterAuthorsRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterAuthorsRequest, response: FilterAuthorsResponse }> {

    let request: FilterAuthorsRequest = {
        page: filterAuthorsRequest.page,
        page_size: filterAuthorsRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterAuthorsRequest = request;

    const response = await fetchFilterAuthors(filterAuthorsRequest);

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
            (filterAuthorsRequest as any)[key] = '';
        });

        const newResponse = await fetchFilterAuthors(filterAuthorsRequest);
        return { request: filterAuthorsRequest, response: newResponse };
    }

    return { request: filterAuthorsRequest, response };
}