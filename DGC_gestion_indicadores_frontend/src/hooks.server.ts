import type { Token } from "$lib/api/model/auth/token";
import type { Handle } from "@sveltejs/kit";
import { jwtDecode } from 'jwt-decode';

export const handle: Handle = async ({event, resolve})=>{
    const session = event.cookies.get("AuthorizationToken")
    
    if(!session){
        return await resolve(event)
    }

    const decodedToken: Token = jwtDecode(session);
    
    event.locals.user = {
        id: decodedToken.id,
        role: decodedToken.role,
    }
    
    return await resolve(event)
}