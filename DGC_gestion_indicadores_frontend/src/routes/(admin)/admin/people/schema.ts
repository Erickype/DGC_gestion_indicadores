import { z } from "zod";

export const addPersonSchema = z.object({
  identity: z.string({
    required_error: "Cédula requerida"
  }).length(10, {
    message: "Ingrese 10 dígitos"
  }).regex(/^\d+$/, {
    message: "Solo se permiten dígitos"
  }),

  name: z.string({
    required_error: "Nombre requerido",
  }).min(5, {
    message: "Nombre mínimo 5 caracteres."
  }).max(50),

  lastname: z.string({
    required_error: "Apellido requerida"
  }).min(5, {
    message: "Apellido mínimo 5 caracteres."
  }).max(50),

  email: z.string({
    required_error: "Abreviación requerida"
  }).email({ message: "Formato: user@mail.com" }),
});

export const updateAcademicPeriodSchema = z.object({
  ID: z.number().gt(0),

  name: z.string({
    required_error: "Nombre requerido",
  }).min(5, {
    message: "Nombre mínimo 5 caracteres."
  }).max(50),

  description: z.string({
    required_error: "Descripción requerida"
  }).min(10, {
    message: "Descripción mínimo 10 caracteres."
  }).max(255),

  abbreviation: z.string({
    required_error: "Abreviación requerida"
  }).min(5, {
    message: "Abreviación mínimo 5 caracteres."
  }).max(255),

  startDate: z.string().refine((v) => v, { message: "Fecha de inicio requerida" }),

  endDate: z.string().refine((v) => v, { message: 'Fecha de fin requerida' }),
});

export type AddPersonSchema = typeof addPersonSchema;
export type UpdateAcademicPeriodSchema = typeof updateAcademicPeriodSchema;