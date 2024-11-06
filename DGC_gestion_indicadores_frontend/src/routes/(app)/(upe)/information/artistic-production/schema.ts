import { z } from "zod";

export const addArtisticProductionListSchema = z.object({
    academic_period_id: z.number({
        required_error: "Periodo académico requerido"
    }).gt(0, {
        message: "Ingrese una periodo académico válido"
    }),

    international_artistic_work: z.coerce.number({
        required_error: "Obra internacional requerida"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    national_artistic_work: z.coerce.number({
        required_error: "Obra nacional requerida"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    intellectual_property: z.coerce.number({
        required_error: "Propiedad intelectual requerida"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),
});

export const updateArtisticProductionListSchema = z.object({
    academic_period_id: z.number({
        required_error: "Periodo académico requerido"
    }).gt(0, {
        message: "Ingrese una periodo académico válido"
    }),

    international_artistic_work: z.coerce.number({
        required_error: "Obra internacional requerida"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),
    
    national_artistic_work: z.coerce.number({
        required_error: "Obra nacional requerida"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),
    
    intellectual_property: z.coerce.number({
        required_error: "Propiedad intelectual requerida"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),
});

export type AddArtisticProductionListSchema = typeof addArtisticProductionListSchema;
export type UpdateArtisticProductionListSchema = typeof updateArtisticProductionListSchema;
