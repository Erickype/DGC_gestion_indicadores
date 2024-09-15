import type { FilterAcademicProductionListsByAcademicPeriodRequest, FilterAcademicProductionListsByAcademicPeriodResponse } from "$lib/api/model/api/indicatorsInformation/academicProductionLists";
import type { CommonError } from "$lib/api/model/errors";

import { goto } from "$app/navigation";
import type { PopoverFilterDataMap } from "$lib/components/table/types";

export function newFilterAcademicProductionListsByAcademiPeriodRequest(page_size: number, page: number, academicPeriodID: number): FilterAcademicProductionListsByAcademicPeriodRequest {
    let filterAcademicProductionRequest: FilterAcademicProductionListsByAcademicPeriodRequest = {
        academic_period_id: academicPeriodID,
        doi: '',
        publication_date: '',
        publication_name: '',
        publication_type: '',
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
    popoverFilterDataMap.set('publication_date', { label: 'Fecha', value: '' });
    popoverFilterDataMap.set('publication_name', { label: 'Nombre', value: '' });
    popoverFilterDataMap.set('publication_type', { label: 'Tipo', value: '' });
    popoverFilterDataMap.set('science_magazine', { label: 'Revista', value: '' });
    popoverFilterDataMap.set('impact_coefficient', { label: 'Coeficiente', value: '' });
    popoverFilterDataMap.set('career', { label: 'Carrera', value: '' });
    popoverFilterDataMap.set('intercultural_component', { label: 'Componente', value: '' });
    return popoverFilterDataMap
}