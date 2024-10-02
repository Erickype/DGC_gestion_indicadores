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

export function generateInitialFilterValue(request: FilterBooksOrChaptersProductionListsByAcademicPeriodRequest): string | undefined {
    let values = [
        request.DOI,
        request.title,
        request.publication_date,
        request.detailed_field
    ];
    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchOnFilterChanged(filter: string, filterBooksOrChaptersProductionListsByAcademicPeriodRequest: FilterBooksOrChaptersProductionListsByAcademicPeriodRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterBooksOrChaptersProductionListsByAcademicPeriodResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterBooksOrChaptersProductionListsByAcademicPeriodRequest as any)[key] = filter;
    });

    return fetchFilterBooksOrChaptersProductionLists(filterBooksOrChaptersProductionListsByAcademicPeriodRequest).then(
        (response: FilterBooksOrChaptersProductionListsByAcademicPeriodResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterBooksOrChaptersProductionListsByAcademicPeriodRequest as any)[key] = '';
                });
                return fetchFilterBooksOrChaptersProductionLists(filterBooksOrChaptersProductionListsByAcademicPeriodRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterBooksOrChaptersProductionListsByAcademicPeriodRequest: FilterBooksOrChaptersProductionListsByAcademicPeriodRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterBooksOrChaptersProductionListsByAcademicPeriodRequest, response: FilterBooksOrChaptersProductionListsByAcademicPeriodResponse }> {

    let request: FilterBooksOrChaptersProductionListsByAcademicPeriodRequest = {
        academic_period_id: filterBooksOrChaptersProductionListsByAcademicPeriodRequest.academic_period_id,
        page: filterBooksOrChaptersProductionListsByAcademicPeriodRequest.page,
        page_size: filterBooksOrChaptersProductionListsByAcademicPeriodRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterBooksOrChaptersProductionListsByAcademicPeriodRequest = request;

    const response = await fetchFilterBooksOrChaptersProductionLists(filterBooksOrChaptersProductionListsByAcademicPeriodRequest);

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
            (filterBooksOrChaptersProductionListsByAcademicPeriodRequest as any)[key] = '';
        });

        const newResponse = await fetchFilterBooksOrChaptersProductionLists(filterBooksOrChaptersProductionListsByAcademicPeriodRequest);
        return { request: filterBooksOrChaptersProductionListsByAcademicPeriodRequest, response: newResponse };
    }

    return { request: filterBooksOrChaptersProductionListsByAcademicPeriodRequest, response };
}