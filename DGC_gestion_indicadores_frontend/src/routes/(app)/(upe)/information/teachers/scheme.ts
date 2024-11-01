import { z } from "zod";

export const addTeacherSchema = z.object({
    academicPeriod: z.number({
        required_error: "Periodo académico requerido"
    }).gt(0, {
        message: "Ingrese un periodo válido"
    }),
    teacher: z.number({
        required_error: "Profesor requerido"
    }).gt(0, {
        message: "No existe el profesor."
    }),
    career: z.number({
        required_error: "Carrera requerida"
    }).gt(0, {
        message: "Ingrese una carrera válida"
    }),
    dedication: z.number({
        required_error: "Dedicación requerida"
    }).gt(0, {
        message: "Ingrese una dedicación válida"
    }),
    contractType: z.number({
        required_error: "Tipo contrato requerido"
    }).gt(0, {
        message: "Ingrese un tipo contrato válido"
    }),
    scaledGrade: z.number().nullable().refine(value => value === null || value > 0, {
        message: "Ingrese un grado escalafonado válido"
    })
}).refine(input => {
    if (input.contractType === 1 && input.scaledGrade === null) {
        return false;
    }
    return true;
}, {
    message: "Grado escalafonario requerido",
    path: ["scaledGrade"]
});


export type AddTeacherSchema = typeof addTeacherSchema;

export const updateTeacherSchema = z.object({
    academicPeriod: z.number().gt(0),
    teacher: z.number({
        required_error: "Profesor requerido"
    }).gt(0, {
        message: "No existe el profesor."
    }),
    career: z.number({
        required_error: "Carrera requerida"
    }).gt(0, {
        message: "Ingrese una carrera válida"
    }),
    dedication: z.number({
        required_error: "Dedicación requerida"
    }).gt(0, {
        message: "Ingrese una dedicación válida"
    }),
    scaledGrade: z.number({
        required_error: "Grado escalafonado requerido"
    }).gt(0, {
        message: "Ingrese un grado escalafonado válido"
    }),
    contractType: z.number({
        required_error: "Tipo contrato requerido"
    }).gt(0, {
        message: "Ingrese un tipo contrato válido"
    })
});

export type UpdateTeacherSchema = typeof updateTeacherSchema;
