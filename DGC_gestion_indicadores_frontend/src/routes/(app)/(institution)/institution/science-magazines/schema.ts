import { z } from "zod";

export const addScienceMagazineSchema = z.object({
    name: z.string({
        required_error: "Nombre requerido"
    }).min(10, {
        message: "Ingrese al menos 10 caracteres"
    }).max(200),

    abbreviation: z.string({
        required_error: "Abbreviación requerida"
    }).min(5, {
        message: "Ingrese al menos 5 caracteres"
    }).max(200),

    description: z.string({
        required_error: "Descripción requerida"
    }).min(10, {
        message: "Ingrese al menos 10 caracteres"
    }).max(200),

    academic_database_id: z.number({
        required_error: "Base de datos científica requerida"
    }).gt(0, {
        message: "Ingrese un valor válido"
    }),
});

export const updateScienceMagazineSchema = z.object({
    id: z.number().gt(0),

    name: z.string({
        required_error: "Nombre requerido"
    }).min(10, {
        message: "Ingrese al menos 10 caracteres"
    }).max(200),

    abbreviation: z.string({
        required_error: "Abbreviación requerida"
    }).min(5, {
        message: "Ingrese al menos 5 caracteres"
    }).max(200),

    description: z.string({
        required_error: "Descripción requerida"
    }).min(10, {
        message: "Ingrese al menos 10 caracteres"
    }).max(200),

    academic_database_id: z.number({
        required_error: "Base de datos científica requerida"
    }).gt(0, {
        message: "Ingrese un valor válido"
    }),
});

export type AddScienceMagazineSchema = typeof addScienceMagazineSchema;
export type UpdateScienceMagazineSchema = typeof updateScienceMagazineSchema;
