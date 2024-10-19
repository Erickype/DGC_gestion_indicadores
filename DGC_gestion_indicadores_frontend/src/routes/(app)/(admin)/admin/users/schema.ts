import { z } from "zod";

export const updateUserSchema = z.object({
  ID: z.number().gt(0),

  username: z.string({
    required_error: "Nombre requerido",
  }).min(5, {
    message: "Nombre mínimo 5 caracteres."
  }).max(50),

  email: z.string({
    required_error: "Descripción requerida"
  }).email({
    message: "Ingrese un correo válido"
  }),

  role_id: z.number({
    required_error: "Abreviación requerida"
  }).gt(0, {
    message: "Ingrese un rol válido"
  }),
});

export type UpdateUserSchema = typeof updateUserSchema;