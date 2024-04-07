// Base domain
const domain = "http://localhost:8000"

// Auth routes
const authRoute = "/auth/user"
const authBaseURL = domain + authRoute

export const authRouteRegister = authBaseURL + "/register"
export const authRouteLogin = authBaseURL + "/login"

