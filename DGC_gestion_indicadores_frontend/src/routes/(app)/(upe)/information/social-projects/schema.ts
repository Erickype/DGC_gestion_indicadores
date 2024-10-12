import { z } from "zod";

export const addSocialProjectListSchema = z.object({
    academic_period_id: z.number().gt(0),

    career_id: z.number({
        required_error: "Carrera requerida"
    }).gt(0, {
        message: "Ingrese una carrera válida"
    }),

    name: z.string({
        required_error: "Título requerido"
    }).max(1000).min(5, {
        message: "Ingrese un nombre válido"
    })
});

export const updateSocialProjectListSchema = z.object({
    academic_period_id: z.number().gt(0),

    career_id: z.number({
        required_error: "Carrera requerida"
    }).gt(0, {
        message: "Ingrese una carrera válida"
    }),

    name: z.string({
        required_error: "Título requerido"
    }).max(1000).min(5, {
        message: "Ingrese un nombre válido"
    })
});

export type AddSocialProjectListSchema = typeof addSocialProjectListSchema;
export type UpdateSocialProjectListSchema = typeof updateSocialProjectListSchema;
