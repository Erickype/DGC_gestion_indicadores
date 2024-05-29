export interface AcademicPeriod{
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    ID: number;
    name: string;
    description: string;
    abbreviation: string;
    start_date: string;
    end_date: string
}