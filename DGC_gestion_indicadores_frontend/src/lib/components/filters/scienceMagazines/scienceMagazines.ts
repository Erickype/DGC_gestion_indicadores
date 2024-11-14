import type { ScienceMagazineJoined, FilterScienceMagazinesRequest, FilterScienceMagazinesResponse } from "$lib/api/model/api/academicProduction/scienceMagazines/scienceMagazine";
import type { PopoverFilterDataMap } from "$lib/components/table/types";

import type { CommonError } from "$lib/api/model/errors";
import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterScienceMagazinesRequest(page_size: number, page: number): FilterScienceMagazinesRequest {
    let filterScienceMagazinesRequest: FilterScienceMagazinesRequest = {
        academic_database: '',
        science_magazine: '',
        magazine_abbreviation: '',
        magazine_description: '',
        page_size: page_size,
        page: page
    };
    return filterScienceMagazinesRequest
}

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('academic_database', { label: 'Base de datos', value: '' });
    popoverFilterDataMap.set('science_magazine', { label: 'Revista', value: '' });
    popoverFilterDataMap.set('magazine_abbreviation', { label: 'Abreviación', value: '' });
    popoverFilterDataMap.set('magazine_description', { label: 'Descripción', value: '' });
    return popoverFilterDataMap
}

export function generateInitialFilterValue(filterScienceMagazinesRequest: FilterScienceMagazinesRequest): string | undefined {
    let values = [
        filterScienceMagazinesRequest.academic_database,
        filterScienceMagazinesRequest.science_magazine,
        filterScienceMagazinesRequest.magazine_abbreviation,
        filterScienceMagazinesRequest.magazine_description,
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchFilteScienceMagazines(filterScienceMagazinesRequest: FilterScienceMagazinesRequest): Promise<FilterScienceMagazinesResponse> {
    const url = `/api/scienceMagazines/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterScienceMagazinesRequest)
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

export async function fetchGetScienceMagazineFilterJoinedByScienceMagazineID(detailedFieldID: string): Promise<ScienceMagazineJoined> {
    const url = `/api/scienceMagazine/filter/` + detailedFieldID;
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

export async function fetchOnFilterChanged(filter: string, filterScienceMagazinesRequest: FilterScienceMagazinesRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterScienceMagazinesResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterScienceMagazinesRequest as any)[key] = filter;
    });

    return fetchFilteScienceMagazines(filterScienceMagazinesRequest).then(
        (response: FilterScienceMagazinesResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterScienceMagazinesRequest as any)[key] = '';
                });
                return fetchFilteScienceMagazines(filterScienceMagazinesRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterScienceMagazinesRequest: FilterScienceMagazinesRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterScienceMagazinesRequest, response: FilterScienceMagazinesResponse }> {

    let request: FilterScienceMagazinesRequest = {
        page: filterScienceMagazinesRequest.page,
        page_size: filterScienceMagazinesRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterScienceMagazinesRequest = request;

    const response = await fetchFilteScienceMagazines(filterScienceMagazinesRequest);

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
            (filterScienceMagazinesRequest as any)[key] = '';
        });

        const newResponse = await fetchFilteScienceMagazines(filterScienceMagazinesRequest);
        return { request: filterScienceMagazinesRequest, response: newResponse };
    }

    return { request: filterScienceMagazinesRequest, response };
}
