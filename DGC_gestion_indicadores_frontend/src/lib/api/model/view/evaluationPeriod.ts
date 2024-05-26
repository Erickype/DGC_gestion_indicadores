export interface EvaluationPeriod {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    ID: number;
    name: string;
    description: string;
    abbreviation: string;
    start_year: string;
    end_year: string;
}

export interface PostEvaluationPeriodRequest {
    name: string;
    description: string;
    abbreviation: string;
    start_year: string;
    end_year: string;
}

export interface UpdateEvaluationPeriodRequest {
    ID: number;
    name: string;
    description: string;
    abbreviation: string;
    start_year: string;
    end_year: string;
}