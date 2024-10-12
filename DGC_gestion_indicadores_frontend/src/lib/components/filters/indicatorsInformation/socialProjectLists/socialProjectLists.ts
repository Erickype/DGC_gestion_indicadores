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