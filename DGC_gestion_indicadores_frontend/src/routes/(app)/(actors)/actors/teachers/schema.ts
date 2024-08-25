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

export const updatePersonSchema = z.object({
  ID: z.number().gt(0),

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

export type AddPersonSchema = typeof addPersonSchema;
export type UpdatePersonSchema = typeof updatePersonSchema;