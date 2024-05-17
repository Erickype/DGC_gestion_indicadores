import { z } from "zod";

export const addEvaluationPeriodSchema = z.object({
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

export type AddEvaluationPeriodSchema = typeof addEvaluationPeriodSchema;