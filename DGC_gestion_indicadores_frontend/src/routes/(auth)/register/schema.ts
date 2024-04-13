import { z } from "zod";

export const registerSchema = z.object({
  username: z.string({
    required_error: "Usuario requerido",
  }).min(5, {
    message: "Usuario mínimo 5 caracteres."
  }).max(50),

  email: z.string({
    required_error: "Email requerido"
  }).email({
    message: "Email inválido"
  }),

  password: z.string({
    required_error: "Contraseña requerida"
  }).min(8, {
    message: "Contraseña mínimo 8 caracteres."
  }).max(50)
});

export type RegisterSchema = typeof registerSchema;