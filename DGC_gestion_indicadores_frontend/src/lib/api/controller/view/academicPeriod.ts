import type { AcademicPeriod } from "$lib/api/model/view/academicPeriod";
import { getAcademicPeriodsRoute } from "$lib/api/routes/view/academicPeriod";
import type { Message } from "$lib/components/combobox/combobox";

export async function GetAcademicPeriods() {
    return await fetch(getAcademicPeriodsRoute, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
}


export async function LoadAcademicPeriodsWithComboMessages() {
    const res = await GetAcademicPeriods();
    if (!res.ok) {
        throw new Error('Fallo cargando periodos');
    }
    const periods: AcademicPeriod[] = await res.json();
    let messages: Message[] = []
    messages = messages.concat(
        periods.map((period) => ({
            value: period.name,
            label: period.name,
        }))
    );

    return {
        periods,
        messages
    }
}