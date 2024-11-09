import { z } from "zod";

export const addPostgraduateProgramSchema = z.object({
    name: z.string({
        required_error: "Nombre requerido"
    }).min(10, {
        message: "Ingrese al menos 10 caracteres"
    }).max(1000),

    start_year: z.coerce.number({
        required_error: "Año inicio requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    end_year: z.coerce.number({
        required_error: "Año fin requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    is_active: z.boolean({
        required_error: "Programa activo requerido"
    }).default(true),
});

export const updatePostgraduateProgramSchema = z.object({
    id: z.number().gt(0),

    name: z.string({
        required_error: "Nombre requerido"
    }).min(10, {
        message: "Ingrese al menos 10 caracteres"
    }).max(1000),

    start_year: z.coerce.number({
        required_error: "Año inicio requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    end_year: z.coerce.number({
        required_error: "Año fin requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    is_active: z.boolean({
        required_error: "Programa activo requerido"
    }).default(true),
});

export type AddPostgraduateProgramSchema = typeof addPostgraduateProgramSchema;
export type UpdatePostgraduateProgramSchema = typeof updatePostgraduateProgramSchema;
