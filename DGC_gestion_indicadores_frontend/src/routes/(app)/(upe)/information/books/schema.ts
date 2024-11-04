import { z } from "zod";

export const addBookOrChaptersProductionListSchema = z.object({
    doi: z.string({
        required_error: "DOI requerido"
    }).url({
        message: "URL inválida"
    }).max(1000),

    is_chapter: z.boolean({
        required_error: "Tipo de publicación requerido"
    }).default(false),

    title: z.string({
        required_error: "Título requerido"
    }).max(1000).min(5, {
        message: "Ingrese un nombre válido"
    }),

    publication_date: z.string({
        required_error: "Fecha publicación requerida"
    }).refine((v) => v, { message: "Fecha publicación requerida" }),

    peer_reviewed: z.boolean({
        required_error: "Tipo de publicación requerido"
    }).default(false),

    academic_period_id: z.number().gt(0),

    detailed_field_id: z.number({
        required_error: "Campo detallado requerido"
    }).gt(0, {
        message: "Ingrese un campo detallado válido"
    }),

    startDate: z.string(),
    endDate: z.string(),
});

export const updateBookOrChaptersProductionListSchema = z.object({
    id: z.number().gt(0),

    doi: z.string({
        required_error: "DOI requerido"
    }).url({
        message: "URL inválida"
    }).max(1000),

    is_chapter: z.boolean({
        required_error: "Tipo de publicación requerido"
    }).default(false),

    title: z.string({
        required_error: "Título requerido"
    }).max(1000).min(5, {
        message: "Ingrese un nombre válido"
    }),

    publication_date: z.string({
        required_error: "Fecha publicación requerida"
    }).refine((v) => v, { message: "Fecha publicación requerida" }),

    peer_reviewed: z.boolean({
        required_error: "Tipo de publicación requerido"
    }).default(false),

    academic_period_id: z.number().gt(0),

    detailed_field_id: z.number({
        required_error: "Campo detallado requerido"
    }).gt(0, {
        message: "Ingrese un campo detallado válido"
    }),

    startDate: z.string(),
    endDate: z.string(),
});

export type AddBookOrChaptersProductionListSchema = typeof addBookOrChaptersProductionListSchema;
export type UpdateBookOrChaptersProductionListSchema = typeof updateBookOrChaptersProductionListSchema;
