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

export async function fetchFiterTeachersLists(filterTeachersListsByAcademicPeriodRequest: FilterTeachersListsByAcademicPeriodRequest):
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