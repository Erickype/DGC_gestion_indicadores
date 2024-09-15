import type { FilterTeachersListsByAcademicPeriodRequest, FilterTeachersListsByAcademicPeriodResponse } from "$lib/api/model/api/indicatorsInformation/teachersLists";
import type { PopoverFilterDataMap } from "$lib/components/table/types";
import type { CommonError } from "$lib/api/model/errors";

import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterTeachersListsByAcademiPeriodRequest(page_size: number, page: number, academicPeriodID: number): FilterTeachersListsByAcademicPeriodRequest {
    let filterTeachersRequest: FilterTeachersListsByAcademicPeriodRequest = {
        academic_period_id: academicPeriodID,
        teacher_identity: '',
        teacher_name: '',
        teacher_lastname: '',
        career: '',
        dedication: '',
        scaled_grade: '',
        contract_type: '',
        page_size: page_size,
        page: page
    };
    return filterTeachersRequest
}

export function newPopoverFilterDataMap(): PopoverFilterDataMap {
    let popoverFilterDataMap: PopoverFilterDataMap = new Map();
    popoverFilterDataMap.set('teacher_identity', { label: 'Identificación', value: '' });
    popoverFilterDataMap.set('teacher_name', { label: 'Nombre', value: '' });
    popoverFilterDataMap.set('teacher_lastname', { label: 'Apellido', value: '' });
    popoverFilterDataMap.set('career', { label: 'Carrera', value: '' });
    popoverFilterDataMap.set('dedication', { label: 'Dedicación', value: '' });
    popoverFilterDataMap.set('scaled_grade', { label: 'Grado escalafonado', value: '' });
    popoverFilterDataMap.set('contract_type', { label: 'Tipo contrato', value: '' });
    return popoverFilterDataMap
}

export async function fetchFilterTeachersLists(filterTeachersListsByAcademicPeriodRequest: FilterTeachersListsByAcademicPeriodRequest):
    Promise<FilterTeachersListsByAcademicPeriodResponse> {
    const url = `/api/indicatorsInformation/teachersLists/filter`;
    const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(filterTeachersListsByAcademicPeriodRequest)
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

export function generateInitialFilterValue(filterTeachersListsRequest: FilterTeachersListsByAcademicPeriodRequest): string | undefined {
    let values = [
        filterTeachersListsRequest.teacher_identity,
        filterTeachersListsRequest.teacher_name,
        filterTeachersListsRequest.teacher_lastname,
        filterTeachersListsRequest.career,
        filterTeachersListsRequest.dedication,
        filterTeachersListsRequest.scaled_grade,
        filterTeachersListsRequest.contract_type,
    ];

    let uniqueValues = [...new Set(values.filter((value) => value !== ''))];

    let initialFilterValue: string | undefined =
        uniqueValues.length === 1 ? uniqueValues[0] : uniqueValues.join(' ');

    return initialFilterValue
}

export async function fetchOnFilterChanged(filter: string, filterTeachersListsByAcademicPeriodRequest: FilterTeachersListsByAcademicPeriodRequest, popoverFilterDataMap: PopoverFilterDataMap): Promise<FilterTeachersListsByAcademicPeriodResponse> {
    popoverFilterDataMap.forEach((_, key) => {
        (filterTeachersListsByAcademicPeriodRequest as any)[key] = filter;
    });

    return fetchFilterTeachersLists(filterTeachersListsByAcademicPeriodRequest).then(
        (response: FilterTeachersListsByAcademicPeriodResponse) => {
            if (response.count === 0) {
                toast.warning(`No hay datos para el filtro: ${filter}`);
                popoverFilterDataMap.forEach((_, key) => {
                    (filterTeachersListsByAcademicPeriodRequest as any)[key] = '';
                });
                return fetchFilterTeachersLists(filterTeachersListsByAcademicPeriodRequest);
            }
            return response;
        }
    );
}

export async function fetchOnDetailedFilter(
    filterTeachersRequest: FilterTeachersListsByAcademicPeriodRequest,
    popoverFilterDataMap: PopoverFilterDataMap
): Promise<{ request: FilterTeachersListsByAcademicPeriodRequest, response: FilterTeachersListsByAcademicPeriodResponse }> {

    let request: FilterTeachersListsByAcademicPeriodRequest = {
        academic_period_id: filterTeachersRequest.academic_period_id,
        page: filterTeachersRequest.page,
        page_size: filterTeachersRequest.page_size
    };

    popoverFilterDataMap.forEach((item, key) => {
        if (item.value !== '') {
            (request as any)[key] = item.value;
        }
    });

    filterTeachersRequest = request;

    const response = await fetchFilterTeachersLists(filterTeachersRequest);

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
            (filterTeachersRequest as any)[key] = '';
        });

        const newResponse = await fetchFilterTeachersLists(filterTeachersRequest);
        return { request: filterTeachersRequest, response: newResponse };
    }

    return { request: filterTeachersRequest, response };
}