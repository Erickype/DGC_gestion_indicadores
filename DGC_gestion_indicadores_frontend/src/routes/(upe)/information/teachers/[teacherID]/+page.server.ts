import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params }) => {
    const param = parseInt(params.teacherID)
    if(param > 0){
        return{
            data: "hi "+param
        }
    }

    error(404, 'Not found');
};