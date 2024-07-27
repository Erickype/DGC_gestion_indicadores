import CircleGauge from "lucide-svelte/icons/circle-gauge"
import FileText from "lucide-svelte/icons/file-text"
import Settings from "lucide-svelte/icons/settings"
import School from "lucide-svelte/icons/school";

export interface Menu {
    name: string,
    links: Link[]
    icon: any
    roles: number[]
}

interface Link {
    name: string
    route: string
}

export const menus: Menu[] = [
    {
        name: 'Dashboards',
        icon: CircleGauge,
        links: [
            { name: 'Periodos Evaluación', route: '/dashboards/evaluation-periods' },
            { name: 'Periodos Académicos', route: '/dashboards/academic-periods' },
            { name: 'Tasas académicas', route: '/dashboards/academic-fees' }
        ],
        roles: [1, 2, 3]
    },
    {
        name: 'Información',
        icon: FileText,
        links: [
            { name: 'Docentes', route: '/information/teachers' },
            { name: 'Publicaciones Académicas', route: '/information/academic-publications' },
            { name: 'Libros', route: '/information/books' }
        ],
        roles: [1, 2]
    },
    {
        name: "Institución",
        icon: School,
        links: [
            { name: "Facultades", route: "/institution/faculties" }
        ],
        roles: [1]
    },
    {
        name: 'Administrador',
        icon: Settings,
        links: [
            { name: 'Periodos Evaluación', route: '/admin/evaluation-periods' },
            { name: 'Periodos Académicos', route: '/admin/academic-periods' },
            { name: 'Personas', route: '/admin/people' },
            { name: 'Usuarios', route: '/admin/users' },
            { name: 'Admin', route: '/admin' },
        ],
        roles: [1]
    }
];