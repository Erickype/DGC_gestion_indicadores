import { z } from "zod";

export const addTeacherSchema = z.object({
    person: z.string({
        required_error: "Persona requerida"
    }).min(10, {
        message: "Ingresa una cédula válida."
    }),
    career: z.string({
        required_error: "Carrera requerida"
    }).min(2, {
        message: "Ingrese una carrera válida"
    })
});

export type AddTeacherSchema = typeof addTeacherSchema;