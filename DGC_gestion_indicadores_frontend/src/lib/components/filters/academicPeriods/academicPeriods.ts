import type { FilterAcademicPeriodsRequest, FilterAcademicPeriodsResponse } from "$lib/api/model/view/academicPeriod";
import type { PopoverFilterDataMap } from "$lib/components/table/types";

import type { CommonError } from "$lib/api/model/errors";
import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterAcademicPeriodsRequest(page_size: number, page: number): FilterAcademicPeriodsRequest {
    let filterAcademicPeriodsRequest: FilterAcademicPeriodsRequest = {
        name: '',
        abbreviation: '',
        description: '',
        start_date: '',
        end_date: '',
        page_size: page_size,
        page: page
    };
    return filterAcademicPeriodsRequest
}

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('name', { label: 'Nombre', value: '' });
    popoverFilterDataMap.set('abbreviation', { label: 'Abreviación', value: '' });
    popoverFilterDataMap.set('description', { label: 'Descripción', value: '' });
    popoverFilterDataMap.set('start_date', { label: 'Fecha inicio', value: '' });
    popoverFilterDataMap.set('end_date', { label: 'Fecha fin', value: '' });
    return popoverFilterDataMap
}

export function generateInitialFilterValue(filterAcademicPeriodsRequest: FilterAcademicPeriodsRequest): string | undefined {
    let values = [
        filterAcademicPeriodsRequest.name,
        filterAcademicPeriodsRequest.description,
        filterAcademicPeriodsRequest.abbreviation,
        filterAcademicPeriodsRequest.start_date,
        filterAcademicPeriodsRequest.end_date
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchFilterAcademicPeriods(filterAcademicPeriodsRequest: FilterAcademicPeriodsRequest): Promise<FilterAcademicPeriodsResponse> {
    const url = `/api/academicPeriods/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterAcademicPeriodsRequest)
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

export async function fetchOnFilterChanged(filter: string, filterAcademicPeriodsRequest: FilterAcademicPeriodsRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterAcademicPeriodsResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterAcademicPeriodsRequest as any)[key] = filter;
    });

    return fetchFilterAcademicPeriods(filterAcademicPeriodsRequest).then(
        (response: FilterAcademicPeriodsResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterAcademicPeriodsRequest as any)[key] = '';
                });
                return fetchFilterAcademicPeriods(filterAcademicPeriodsRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterAcademicPeriodsRequest: FilterAcademicPeriodsRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterAcademicPeriodsRequest, response: FilterAcademicPeriodsResponse }> {

    let request: FilterAcademicPeriodsRequest = {
        page: filterAcademicPeriodsRequest.page,
        page_size: filterAcademicPeriodsRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterAcademicPeriodsRequest = request;

    const response = await fetchFilterAcademicPeriods(filterAcademicPeriodsRequest);

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
            (filterAcademicPeriodsRequest as any)[key] = '';
        });

        const newResponse = await fetchFilterAcademicPeriods(filterAcademicPeriodsRequest);
        return { request: filterAcademicPeriodsRequest, response: newResponse };
    }

    return { request: filterAcademicPeriodsRequest, response };
}