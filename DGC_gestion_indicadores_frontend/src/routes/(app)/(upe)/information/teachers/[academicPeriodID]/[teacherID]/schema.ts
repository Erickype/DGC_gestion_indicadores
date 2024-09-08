import { z } from "zod";

export const addDegreeAndTeachersListsDegreeSchema = z.object({
    academicPeriodID: z.number().gt(0),

    teacherID: z.number().gt(0),

    degreeLevelID: z.number({
        required_error: "Nivel requerido"
    }).gt(0),

    name: z.string({
        required_error: "Nombre requerida"
    }).min(10, {
        message: "Nombre mínimo 10 caracteres."
    }).max(200, {
        message: "Nombre máximo 200 caracteres."
    }),
});

export type AddDegreeAndTeachersListsDegreeSchema = typeof addDegreeAndTeachersListsDegreeSchema;
