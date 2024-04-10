interface Menu{
    name: string,
    links: Link[]
    icon: any
    roles: number[]
}

interface Link{
    name: string
    route: string
}