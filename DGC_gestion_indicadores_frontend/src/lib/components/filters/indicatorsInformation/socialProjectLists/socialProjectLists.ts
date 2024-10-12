import type { FilterSocialProjectListsByAcademicPeriodRequest, FilterSocialProjectListsByAcademicPeriodResponse } from "$lib/api/model/api/indicatorsInformation/socialProjectLists";
import type { PopoverFilterDataMap } from "$lib/components/table/types";
import type { CommonError } from "$lib/api/model/errors";

import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterSocialProjectListsByAcademiPeriodRequest(page_size: number, page: number, academicPeriodID: number): FilterSocialProjectListsByAcademicPeriodRequest {
    let filterSocialProjectListsByAcademicPeriodRequest: FilterSocialProjectListsByAcademicPeriodRequest = {
        academic_period_id: academicPeriodID,
        career: '',
        name: '',
        page_size: page_size,
        page: page
    };
    return filterSocialProjectListsByAcademicPeriodRequest
}

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('career', { label: 'Carrera', value: '' });
    popoverFilterDataMap.set('name', { label: 'Nombre', value: '' });
    return popoverFilterDataMap
}

export async function fetchFilterSocialProjectLists(filterSocialProjectListsByAcademicPeriodRequest: FilterSocialProjectListsByAcademicPeriodRequest):
    Promise<FilterSocialProjectListsByAcademicPeriodResponse> {
    const url = `/api/indicatorsInformation/socialProjectLists/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterSocialProjectListsByAcademicPeriodRequest)
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

export function generateInitialFilterValue(filterTeachersListsRequest: FilterSocialProjectListsByAcademicPeriodRequest): string | undefined {
    let values = [
        filterTeachersListsRequest.career,
        filterTeachersListsRequest.name,
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchOnFilterChanged(filter: string, filterSocialProjectListsByAcademicPeriodRequest: FilterSocialProjectListsByAcademicPeriodRequest, popoverFilterDataMap: PopoverFilterDataMap) {
    popoverFilterDataMap.forEach((_, key) => {
        (filterSocialProjectListsByAcademicPeriodRequest as any)[key] = filter;
    });

    return fetchFilterSocialProjectLists(filterSocialProjectListsByAcademicPeriodRequest).then(
        (response: FilterSocialProjectListsByAcademicPeriodResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterSocialProjectListsByAcademicPeriodRequest as any)[key] = '';
                });
                return fetchFilterSocialProjectLists(filterSocialProjectListsByAcademicPeriodRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterSocialProjectListsByAcademicPeriodRequest: FilterSocialProjectListsByAcademicPeriodRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterSocialProjectListsByAcademicPeriodRequest, response: FilterSocialProjectListsByAcademicPeriodResponse }> {

    let request: FilterSocialProjectListsByAcademicPeriodRequest = {
        academic_period_id: filterSocialProjectListsByAcademicPeriodRequest.academic_period_id,
        page: filterSocialProjectListsByAcademicPeriodRequest.page,
        page_size: filterSocialProjectListsByAcademicPeriodRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterSocialProjectListsByAcademicPeriodRequest = request;

    const response = await fetchFilterSocialProjectLists(filterSocialProjectListsByAcademicPeriodRequest);

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
            (filterSocialProjectListsByAcademicPeriodRequest as any)[key] = '';
        });

        const newResponse = await fetchFilterSocialProjectLists(filterSocialProjectListsByAcademicPeriodRequest);
        return { request: filterSocialProjectListsByAcademicPeriodRequest, response: newResponse };
    }

    return { request: filterSocialProjectListsByAcademicPeriodRequest, response };
}