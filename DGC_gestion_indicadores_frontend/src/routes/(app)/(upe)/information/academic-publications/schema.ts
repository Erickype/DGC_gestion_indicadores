import { z } from "zod";

export const addAcademicProductionSchema = z.object({
    doi: z.string({
        required_error: "DOI requerido"
    }).max(1000),

    academicPeriod: z.number({
        required_error: "Periodo académico requirido"
    }).gt(0, {
        message: "Ingrese un periodo válido"
    }),

    publication_date: z.string().refine((v) => v, { message: "Fecha publicación requerida" }),

    publication_name: z.string({
        required_error: "Nombre publicación requerido"
    }).max(1000),

    detailed_field_id: z.number({
        required_error: "Campo detallado requirido"
    }).gt(0, {
        message: "Ingrese un campo detallado válida"
    }),

    science_magazine_id: z.number({
        required_error: "Revista requirida"
    }).gt(0, {
        message: "Ingrese una revista válida"
    }),

    impact_coefficient_id: z.number({
        required_error: "Coeficiente impacto requirido"
    }).gt(0, {
        message: "Ingrese un coeficiente impacto válido"
    }),

    intercultural_component: z.boolean({
        required_error: "Componente intercultural requerido"
    }).default(false),
});

export type AddAcademicProductionSchema = typeof addAcademicProductionSchema;
