import { z } from "zod";

export const loginSchema = z.object({
  username: z.string({
    required_error: "Usuario requerido",
  }).min(5, {
    message: "Usuario mínimo 5 caracteres."
  }).max(50),

  password: z.string({
    required_error: "Contraseña requerida"
  }).min(8, {
    message: "Contraseña mínimo 8 caracteres."
  }).max(50)
});

export type LoginSchema = typeof loginSchema;