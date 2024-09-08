import { z } from "zod";

export const addTeachersListsDegreeSchema = z.object({
    academicPeriodID: z.number().gt(0),

    teacherID: z.number().gt(0),

    degreeLevelID: z.number({
        required_error: "Nivel requerido"
    }).gt(0),

    name: z.string({
        required_error: "Nombre requerida"
    }).min(10, {
        message: "Descripción mínimo 10 caracteres."
    }).max(50),
});

export type AddTeachersListsDegreeSchema = typeof addTeachersListsDegreeSchema;
