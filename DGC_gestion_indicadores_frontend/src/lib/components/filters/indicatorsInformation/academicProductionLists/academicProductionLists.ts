import type { FilterAcademicProductionListsByAcademicPeriodRequest, FilterAcademicProductionListsByAcademicPeriodResponse } from "$lib/api/model/api/indicatorsInformation/academicProductionLists";
import type { PopoverFilterDataMap } from "$lib/components/table/types";
import type { CommonError } from "$lib/api/model/errors";

import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterAcademicProductionListsByAcademiPeriodRequest(page_size: number, page: number, academicPeriodID: number): FilterAcademicProductionListsByAcademicPeriodRequest {
    let filterAcademicProductionRequest: FilterAcademicProductionListsByAcademicPeriodRequest = {
        academic_period_id: academicPeriodID,
        doi: '',
        publication_date: '',
        publication_name: '',
        detailed_field: '',
        science_magazine: '',
        impact_coefficient: '',
        career: '',
        page_size: page_size,
        page: page
    };
    return filterAcademicProductionRequest
}

export async function fetchFilterAcademicProductionLists(filterAcademicProductionRequest: FilterAcademicProductionListsByAcademicPeriodRequest):
    Promise<FilterAcademicProductionListsByAcademicPeriodResponse> {
    const url = `/api/indicatorsInformation/academicProduction/filter`;
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
    popoverFilterDataMap.set('doi', { label: 'DOI', value: '' });
/*     popoverFilterDataMap.set('publication_date', { label: 'Fecha', value: '' });
 */    popoverFilterDataMap.set('publication_name', { label: 'Nombre', value: '' });
    popoverFilterDataMap.set('detailed_field', { label: 'Campo detallado', value: '' });
    popoverFilterDataMap.set('science_magazine', { label: 'Revista', value: '' });
    popoverFilterDataMap.set('impact_coefficient', { label: 'Coeficiente', value: '' });
    popoverFilterDataMap.set('career', { label: 'Carrera', value: '' });
/*     popoverFilterDataMap.set('intercultural_component', { label: 'Componente', value: '' });
 */    return popoverFilterDataMap
}

export function generateInitialFilterValue(filterAcademicProductionListsRequest: FilterAcademicProductionListsByAcademicPeriodRequest): string | undefined {
    let values = [
        filterAcademicProductionListsRequest.doi,
        filterAcademicProductionListsRequest.publication_date,
        filterAcademicProductionListsRequest.publication_name,
        filterAcademicProductionListsRequest.detailed_field,
        filterAcademicProductionListsRequest.science_magazine,
        filterAcademicProductionListsRequest.impact_coefficient,
        filterAcademicProductionListsRequest.career
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchOnFilterChanged(filter: string, filterAcademicProductionListsRequest: FilterAcademicProductionListsByAcademicPeriodRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterAcademicProductionListsByAcademicPeriodResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterAcademicProductionListsRequest as any)[key] = filter;
    });

    return fetchFilterAcademicProductionLists(filterAcademicProductionListsRequest).then(
        (response: FilterAcademicProductionListsByAcademicPeriodResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterAcademicProductionListsRequest as any)[key] = '';
                });
                return fetchFilterAcademicProductionLists(filterAcademicProductionListsRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterAcademicProductionListsRequest: FilterAcademicProductionListsByAcademicPeriodRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterAcademicProductionListsByAcademicPeriodRequest, response: FilterAcademicProductionListsByAcademicPeriodResponse }> {

    let request: FilterAcademicProductionListsByAcademicPeriodRequest = {
        academic_period_id: filterAcademicProductionListsRequest.academic_period_id,
        page: filterAcademicProductionListsRequest.page,
        page_size: filterAcademicProductionListsRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterAcademicProductionListsRequest = request;

    const response = await fetchFilterAcademicProductionLists(filterAcademicProductionListsRequest);

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
            (filterAcademicProductionListsRequest as any)[key] = '';
        });

        const newResponse = await fetchFilterAcademicProductionLists(filterAcademicProductionListsRequest);
        return { request: filterAcademicProductionListsRequest, response: newResponse };
    }

    return { request: filterAcademicProductionListsRequest, response };
}