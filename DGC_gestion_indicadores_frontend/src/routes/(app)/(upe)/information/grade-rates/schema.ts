import { z } from "zod";

export const addGradeRateListSchema = z.object({
    academic_period_id: z.number().gt(0),

    career_id: z.number({
        required_error: "Carrera requerida"
    }).gt(0, {
        message: "Ingrese una carrera válida"
    }),

    count_graduated_students: z.number({
        required_error: "Estudiantes graduados requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    count_court_students: z.number({
        required_error: "Estudiantes en cohorte requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    count_admitted_matriculated_students: z.number({
        required_error: "Estudiantes admitidos matriculados requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    count_admitted_students: z.number({
        required_error: "Estudiantes admitidos requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),
});

export const updateGradeRateListSchema = z.object({
    academic_period_id: z.number().gt(0),

    career_id: z.number({
        required_error: "Carrera requerida"
    }).gt(0, {
        message: "Ingrese una carrera válida"
    }),

    count_graduated_students: z.number({
        required_error: "Estudiantes graduados requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    count_court_students: z.number({
        required_error: "Estudiantes en cohorte requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    count_admitted_matriculated_students: z.number({
        required_error: "Estudiantes admitidos matriculados requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),

    count_admitted_students: z.number({
        required_error: "Estudiantes admitidos requerido"
    }).gt(0, {
        message: "Ingrese una valor válido"
    }),
});

export type AddGradeRateListSchema = typeof addGradeRateListSchema;
export type UpdateGradeRateListSchema = typeof updateGradeRateListSchema;
