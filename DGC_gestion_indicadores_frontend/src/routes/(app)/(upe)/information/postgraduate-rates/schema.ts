import { z } from "zod";

export const addPostgraduateCohortListSchema = z.object({
    postgraduate_program_id: z.number().gt(0),

    year: z.coerce.number({
        required_error: "Año cohorte requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    name: z.string({
        required_error: "Nombre requerido"
    }).min(10, {
        message: "Ingrese al menos 10 caracteres"
    }).max(1000),

    graduated_students: z.coerce.number({
        required_error: "Estudiantes graduados requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    total_students: z.coerce.number({
        required_error: "Total estudiantes cohorte requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),
});

export const updatePostgraduateCohortListSchema = z.object({
    postgraduate_program_id: z.number().gt(0),

    year: z.coerce.number({
        required_error: "Año cohorte requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    name: z.string({
        required_error: "Nombre requerido"
    }).min(10, {
        message: "Ingrese al menos 10 caracteres"
    }).max(1000),

    graduated_students: z.coerce.number({
        required_error: "Estudiantes graduados requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    total_students: z.coerce.number({
        required_error: "Total estudiantes cohorte requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),
});

export type AddPostgraduateCohortListSchema = typeof addPostgraduateCohortListSchema;
export type UpdatePostgraduateCohortListSchema = typeof updatePostgraduateCohortListSchema;
