export interface Career {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    ID: number;
    faculty_id: number;
    name: string;
    abbreviation: string;
    description: string | null;
}
