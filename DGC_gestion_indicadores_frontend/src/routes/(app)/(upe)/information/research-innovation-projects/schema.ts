import { z } from "zod";

export const addResearchInnovationProjectListSchema = z.object({
    academic_period_id: z.number({
        required_error: "Periodo académico requerido"
    }).gt(0, {
        message: "Ingrese una periodo académico válido"
    }),

    total_projects_uep: z.coerce.number({
        required_error: "Total de UEP requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    projects_external_funding: z.coerce.number({
        required_error: "Financiamiento externo requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    international_cooperation_projects: z.coerce.number({
        required_error: "Internacional requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    national_cooperation_projects: z.coerce.number({
        required_error: "Nacional requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),
});

export const updateResearchInnovationProjectListSchema = z.object({
    academic_period_id: z.number({
        required_error: "Periodo académico requerido"
    }).gt(0, {
        message: "Ingrese una periodo académico válido"
    }),

    total_projects_uep: z.coerce.number({
        required_error: "Total de UEP requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    projects_external_funding: z.coerce.number({
        required_error: "Financiamiento externo requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    international_cooperation_projects: z.coerce.number({
        required_error: "Internacional requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    national_cooperation_projects: z.coerce.number({
        required_error: "Nacional requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),
});

export type AddResearchInnovationProjectListSchema = typeof addResearchInnovationProjectListSchema;
export type UpdateResearchInnovationProjectListSchema = typeof updateResearchInnovationProjectListSchema;
