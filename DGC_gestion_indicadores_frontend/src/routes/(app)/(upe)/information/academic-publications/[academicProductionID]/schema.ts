import { z } from "zod";

export const addAcademicProductionListsAuthorSchema = z.object({
    academicProductionID: z.number({
        required_error: "DOI requerido"
    }).gt(0),

    authorID: z.number({
        required_error: "Autor requerido"
    }).gt(0, {
        message: "Seleccione un autor válido"
    }),

    careers: z.array(z.number()).refine((value) => value.some((item) => item), {
        message: "Seleccionar al menos una carrera"
    }),
});

export type AddAcademicProductionListsAuthorSchema = typeof addAcademicProductionListsAuthorSchema;
