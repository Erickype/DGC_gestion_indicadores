export interface Person {
    CreatedAt?: string;
    UpdatedAt?: string;
    ID?: number;
    identity: string;
    name: string;
    lastname: string;
    email: string;
}

export interface PostPersonRequest {
    identity: string;
    name: string;
    lastname: string;
    email: string;
}

export interface PostPersonWithRolesRequest {
    roles: string[]
    person: Person
}

export interface UpdatePersonWithRolesRequest {
    roles: string[]
    person: Person
}

export interface PutPersonRequest {
    ID: number;
    identity: string;
    name: string;
    lastname: string;
    email: string;
}

export interface FilterPeopleRequest {
    identity?: string,
    name?: string,
    lastname?: string,
    email?: string,
    page_size: number,
    page: number
}

export interface PersonWithRoles {
    ID: number
    identity: string
    name: string
    lastname: string
    email: string
    roles: string[]
}

export interface FilterPeopleResponse {
    count: number,
    page_size: number,
    page: number,
    people: PersonWithRoles[]
}