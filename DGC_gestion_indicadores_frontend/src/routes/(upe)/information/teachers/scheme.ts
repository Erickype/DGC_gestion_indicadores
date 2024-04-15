import { z } from "zod";

export const addTeacherSchema = z.object({
    person: z.string({
        required_error: "Persona requerida"
    }).min(10,{
        message: "Ingresa una cédula válida."
    })
});

export type AddTeacherSchema = typeof addTeacherSchema;