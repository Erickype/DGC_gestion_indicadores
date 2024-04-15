import { z } from "zod";

export const addTeacherSchema = z.object({
    academicPeriod: z.string({
        required_error: "Periodo académico requirido"
    }).min(5, {
        message: "Ingrese un periodo válido"
    }),
    person: z.string({
        required_error: "Persona requerida"
    }).min(10, {
        message: "Ingrese una cédula válida."
    }),
    career: z.string({
        required_error: "Carrera requerida"
    }).min(2, {
        message: "Ingrese una carrera válida"
    })
});

export type AddTeacherSchema = typeof addTeacherSchema;