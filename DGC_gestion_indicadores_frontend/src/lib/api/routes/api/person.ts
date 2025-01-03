import { apiBaseURL } from "./base";

export const getPeopleRoute = apiBaseURL + "/people"
export const postFilterPeopleRoute = apiBaseURL + "/people/filter"

export const postPersonRoute = apiBaseURL + "/person"
export const putPersonRoute = apiBaseURL + "/person/"
export const deletePersonRoute = apiBaseURL + "/person/"

export const postPersonWithRolesRoute = apiBaseURL + "/person/withRoles"
export const updatePersonWithRolesRoute = apiBaseURL + "/person/withRoles"