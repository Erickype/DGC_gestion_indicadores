import type { FilterBooksOrChaptersProductionListsByAcademicPeriodRequest, FilterBooksOrChaptersProductionListsByAcademicPeriodResponse } from "$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionLists";
import type { PopoverFilterDataMap } from "$lib/components/table/types";
import type { CommonError } from "$lib/api/model/errors";

import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterBooksOrChaptersProductionListsByAcademiPeriodRequest(page_size: number, page: number, academicPeriodID: number): FilterBooksOrChaptersProductionListsByAcademicPeriodRequest {
    let filterAcademicProductionRequest: FilterBooksOrChaptersProductionListsByAcademicPeriodRequest = {
        DOI: '',
        title: '',
        detailed_field: '',
        academic_period_id: academicPeriodID,
        page_size: page_size,
        page: page
    };
    return filterAcademicProductionRequest
}

export async function fetchFilterBooksOrChaptersProductionLists(filterAcademicProductionRequest: FilterBooksOrChaptersProductionListsByAcademicPeriodRequest):
    Promise<FilterBooksOrChaptersProductionListsByAcademicPeriodResponse> {
    const url = `/api/indicatorsInformation/booksOrChaptersProductionLists/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterAcademicProductionRequest)
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

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('DOI', { label: 'DOI', value: '' });
    /* popoverFilterDataMap.set('is_chapter', { label: 'Fecha', value: '' }); */
    popoverFilterDataMap.set('title', { label: 'Título', value: '' });
    popoverFilterDataMap.set('publication_date', { label: 'Fecha', value: '' });
    /* popoverFilterDataMap.set('peer_reviewed', { label: 'Revisado pares', value: '' }); */
    popoverFilterDataMap.set('detailed_field', { label: 'Campo detallado', value: '' });
    return popoverFilterDataMap
}