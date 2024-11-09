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
            { name: 'Proyectos innovación', route: '/information/research-innovation-projects' },
            { name: 'Publicaciones Académicas', route: '/information/academic-publications' },
            { name: 'Producción artística', route: '/information/artistic-production' },
            { name: 'Libros o Capítulos', route: '/information/books' },
            { name: 'Proyectos vinculación', route: '/information/social-projects' },
            { name: 'Tasas de grado', route: '/information/grade-rates' },
            { name: 'Tasas de posgrado', route: '/information/postgraduate-rates' },
        ],
        roles: [1, 2]
    },
    {
        name: "Institución",
        icon: School,
        links: [
            { name: "Facultades", route: "/institution/faculties" },
            { name: "Carreras", route: "/institution/careers" },
            { name: "Programas posgrado", route: "/institution/postgraduate/programs" },
        ],
        roles: [1]
    },
    {
        name: "Actores",
        icon: BookUser,
        links: [
            { name: "Personas", route: "/actors/people" },
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
        ],
        roles: [1]
    }
];