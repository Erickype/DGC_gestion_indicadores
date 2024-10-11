export interface AcademicPeriod {
    CreatedAt: string;
    UpdatedAt: string;
    ID: number;
    name: string;
    description: string;
    abbreviation: string;
    start_date: string;
    end_date: string
}

export interface PostAcademicPeriodRequest {
    name: string;
    description: string;
    abbreviation: string;
    start_date: string;
    end_date: string;
}

export interface UpdateAcademicPeriodRequest {
    ID: number;
    name: string;
    description: string;
    abbreviation: string;
    start_date: string;
    end_date: string;
}

export interface FilterAcademicPeriodsRequest {
    name: string;
    description: string;
    abbreviation: string;
    start_date: string;
    end_date: string;
    page_size: number;
    page: number;
}

export interface FilterAcademicPeriodsResponse {
    count: number,
    page_size: number,
    page: number,
    academic_periods: AcademicPeriod[]
}