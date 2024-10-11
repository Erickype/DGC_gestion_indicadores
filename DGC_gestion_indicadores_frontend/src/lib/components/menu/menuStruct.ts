import CircleGauge from "lucide-svelte/icons/circle-gauge"
import FileText from "lucide-svelte/icons/file-text"
import Settings from "lucide-svelte/icons/settings"
import School from "lucide-svelte/icons/school";
import BookUser from "lucide-svelte/icons/book-user";

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
        name: 'Información indicadores',
        icon: FileText,
        links: [
            { name: 'Listas Docentes', route: '/information/teachers' },
            { name: 'Publicaciones Académicas', route: '/information/academic-publications' },
            { name: 'Libros o Capítulos', route: '/information/books' },
            { name: 'Proyectos vinculación', route: '/information/social-projects' },
        ],
        roles: [1, 2]
    },
    {
        name: "Institución",
        icon: School,
        links: [
            { name: "Facultades", route: "/institution/faculties" },
            { name: "Carreras", route: "/institution/careers" },
        ],
        roles: [1]
    },
    {
        name: "Actores",
        icon: BookUser,
        links: [
            { name: "Personas", route: "/actors/people" },
            { name: "Profesores", route: "/actors/teachers" },
            { name: "Autores", route: "/actors/authors" },
        ],
        roles: [1, 2]
    },
    {
        name: 'Administrador',
        icon: Settings,
        links: [
            { name: 'Periodos Evaluación', route: '/admin/evaluation-periods' },
            { name: 'Periodos Académicos', route: '/admin/academic-periods' },
            { name: 'Usuarios', route: '/admin/users' },
            { name: 'Admin', route: '/admin' },
        ],
        roles: [1]
    }
];