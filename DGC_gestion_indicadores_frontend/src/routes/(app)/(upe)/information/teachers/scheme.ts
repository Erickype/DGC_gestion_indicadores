import { z } from "zod";

export const addTeacherSchema = z.object({
    academicPeriod: z.number({
        required_error: "Periodo académico requirido"
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
