import type { FilterAcademicPeriodsRequest } from "$lib/api/model/view/academicPeriod";

import { goto } from "$app/navigation";
import { toast } from "svelte-sonner";

export function newFilterAcademicPeriodsRequest(page_size: number, page: number): FilterAcademicPeriodsRequest {
    let filterAcademicPeriodsRequest: FilterAcademicPeriodsRequest = {
        name: '',
        abbreviation: '',
        description: '',
        start_date: '',
        end_date: '',
        page_size: page_size,
        page: page
    };
    return filterAcademicPeriodsRequest
}