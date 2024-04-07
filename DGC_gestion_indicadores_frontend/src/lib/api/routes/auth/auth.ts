import { domain } from "../base"

// AUTH ROUTES
const authRoute = "/auth/user"
const authBaseURL = domain + authRoute

export const authRouteRegister = authBaseURL + "/register"
export const authRouteLogin = authBaseURL + "/login"