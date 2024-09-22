import type { FilterDetailedFieldRequest, FilterDetailedFieldResponse } from "$lib/api/model/api/knowledgeFields/detailedFields";
import type { PopoverFilterDataMap } from "$lib/components/table/types";

import type { CommonError } from "$lib/api/model/errors";
import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";


export function newFilterDetailedFieldsRequest(page_size: number, page: number): FilterDetailedFieldRequest {
    let filterDetailedFieldRequest: FilterDetailedFieldRequest = {
        wide_field: '',
        specific_field: '',
        detailed_field: '',
        page_size: page_size,
        page: page
    };
    return filterDetailedFieldRequest
}

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('wide_field', { label: 'Campo amplio', value: '' });
    popoverFilterDataMap.set('specific_field', { label: 'Campo específico', value: '' });
    popoverFilterDataMap.set('detailed_field', { label: 'Campo detallado', value: '' });
    return popoverFilterDataMap
}

export function generateInitialFilterValue(filterDetailedFieldRequest: FilterDetailedFieldRequest): string | undefined {
    let values = [
        filterDetailedFieldRequest.wide_field,
        filterDetailedFieldRequest.specific_field,
        filterDetailedFieldRequest.detailed_field,
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchFilterDetailedFields(filterDetailedFieldRequest: FilterDetailedFieldRequest): Promise<FilterDetailedFieldResponse> {
    const url = `/api/knowledgeFields/detailedFields/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterDetailedFieldRequest)
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

export async function fetchOnFilterChanged(filter: string, filterDetailedFieldRequest: FilterDetailedFieldRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterDetailedFieldResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterDetailedFieldRequest as any)[key] = filter;
    });

    return fetchFilterDetailedFields(filterDetailedFieldRequest).then(
        (response: FilterDetailedFieldResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterDetailedFieldRequest as any)[key] = '';
                });
                return fetchFilterDetailedFields(filterDetailedFieldRequest);
            }
            return response;
        }
    );
}