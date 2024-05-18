import type { EvaluationPeriod } from "$lib/api/model/view/evaluationPeriod";
import { getEvaluationPeriodsRoute } from "$lib/api/routes/view/evaluationPeriod";
import type { Message } from "$lib/components/combobox/combobox";

export async function GetEvaluationPeriods() {
    return await fetch(getEvaluationPeriodsRoute, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
}


export async function LoadEvaluationPeriodsWithComboMessages() {
    const res = await GetEvaluationPeriods();
    if (!res.ok) {
        throw new Error('Fallo cargando periodos');
    }
    const periods: EvaluationPeriod[] = await res.json();
    let messages: Message[] = []
    messages = messages.concat(
        periods.map((period) => ({
            value: period.ID,
            label: period.abbreviation,
        }))
    );

    return {
        periods,
        messages
    }
}