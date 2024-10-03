import { z } from "zod";

export const addBooksOrChaptersProductionListsAuthorCareersSchema = z.object({
    academicProductionID: z.number().gt(0),

    authorID: z.number({
        required_error: "Autor requerido"
    }).gt(0, {
        message: "Seleccione un autor válido"
    }),

    career: z.number(),

    careers: z.array(z.number()).refine((value) => value.some((item) => item), {
        message: "Seleccionar al menos una carrera"
    }),
});

export type AddBooksOrChaptersProductionListsAuthorCareersSchema = typeof addBooksOrChaptersProductionListsAuthorCareersSchema