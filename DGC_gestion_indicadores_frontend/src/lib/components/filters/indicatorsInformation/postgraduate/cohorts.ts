import type { FilterCohortListsRequest, FilterCohortListsResponse, CohortList } from "$lib/api/model/api/indicatorsInformation/postgraduate";
import type { PopoverFilterDataMap } from "$lib/components/table/types";

import type { CommonError } from "$lib/api/model/errors";
import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterCohortListsRequest(page_size: number, page: number): FilterCohortListsRequest {
    let filterCohortListsRequest: FilterCohortListsRequest = {
        start_year: 1,
        end_year: 10000,
        page_size: page_size,
        page: page
    };
    return filterCohortListsRequest
}

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('start_year', { label: 'Año inicio', value: '' });
    popoverFilterDataMap.set('end_year', { label: 'Año fin', value: '' });
    return popoverFilterDataMap
}

export function generateInitialFilterValue(filterCohortListsRequest: FilterCohortListsRequest): string | undefined {
    let values = [
        filterCohortListsRequest.start_year?.toString(),
        filterCohortListsRequest.end_year?.toString(),
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchFilterCohorts(filterCohortListsRequest: FilterCohortListsRequest): Promise<FilterCohortListsResponse> {
    const url = `/api/indicatorsInformation/postgraduate/cohorts/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterCohortListsRequest)
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

export async function fetchGetCohortByYear(year: string) {
    const url = `/api/indicatorsInformation/postgraduate/cohort/` + year;
    const response = await fetch(url, {
        method: 'GET',
    });
    if (!response.ok) {
        const errorData = (await response.json()) as CommonError;
        if (response.status === 401) {
            throw goto('/');
        }
        return errorData;
    }
    return response.json();
}

export async function fetchOnFilterChanged(filter: string, filterCohortListsRequest: FilterCohortListsRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterCohortListsResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterCohortListsRequest as any)[key] = filter;
    });

    return fetchFilterCohorts(filterCohortListsRequest).then(
        (response: FilterCohortListsResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterCohortListsRequest as any)[key] = '';
                });
                return fetchFilterCohorts(filterCohortListsRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterCohortListsRequest: FilterCohortListsRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterCohortListsRequest, response: FilterCohortListsResponse }> {

    let request: FilterCohortListsRequest = {
        page: filterCohortListsRequest.page,
        page_size: filterCohortListsRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterCohortListsRequest = request;

    const response = await fetchFilterCohorts(filterCohortListsRequest);

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
            (filterCohortListsRequest as any)[key] = '';
        });

        const newResponse = await fetchFilterCohorts(filterCohortListsRequest);
        return { request: filterCohortListsRequest, response: newResponse };
    }

    return { request: filterCohortListsRequest, response };
}