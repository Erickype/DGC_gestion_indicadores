import { indicatorsInformationBaseRoute } from "./base";

const postgraduateRoute = indicatorsInformationBaseRoute + "/postgraduate"

//Programs
export const postFilterPostgraduateProgramsRoute = postgraduateRoute + "/programs/filter"
export const getPostgraduateProgramByProgramIDRoute = postgraduateRoute + "/program/"
export const postPostgraduateProgramRoute = postgraduateRoute + "/program"
export const updatePostgraduateProgramRoute = postgraduateRoute + "/program/"
