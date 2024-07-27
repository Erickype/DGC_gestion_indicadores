import { z } from "zod";

export const addFacultySchema = z.object({
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
});

export const updateFacultySchema = z.object({
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
});

export type AddFacultySchema = typeof addFacultySchema;
export type UpdateFacultySchema = typeof updateFacultySchema;