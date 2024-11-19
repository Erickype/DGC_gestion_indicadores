import { z } from "zod";

export const addCarrerSchema = z.object({
  facultyID: z.number({
    required_error: "Facultad requerida"
  }).gt(0, {
    message: "Ingrese un valor válido"
  }),
  name: z.string({
    required_error: "Nombre requerido",
  }).min(5, {
    message: "Nombre mínimo 5 caracteres."
  }).max(200),

  description: z.string({
    required_error: "Descripción requerida"
  }).min(5, {
    message: "Descripción mínimo 5 caracteres."
  }).max(50),

  abbreviation: z.string({
    required_error: "Abreviación requerida"
  }).min(2, {
    message: "Abreviación mínimo 2 caracteres."
  }).max(50),
});

export const updateCarrerSchema = z.object({
  ID: z.number().gt(0),

  facultyID: z.number({
    required_error: "Facultad requerida"
  }).gt(0, {
    message: "Ingrese un valor válido"
  }),

  name: z.string({
    required_error: "Nombre requerido",
  }).min(5, {
    message: "Nombre mínimo 5 caracteres."
  }).max(200),

  description: z.string({
    required_error: "Descripción requerida"
  }).min(5, {
    message: "Descripción mínimo 5 caracteres."
  }).max(50),

  abbreviation: z.string({
    required_error: "Abreviación requerida"
  }).min(2, {
    message: "Abreviación mínimo 2 caracteres."
  }).max(50),
});

export type AddCarrerSchema = typeof addCarrerSchema;
export type UpdateCarrerSchema = typeof updateCarrerSchema;