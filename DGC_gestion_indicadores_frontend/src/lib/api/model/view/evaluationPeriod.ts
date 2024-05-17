export interface EvaluationPeriod {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    ID: number;
    name: string;
    description: string;
    abbreviation: string;
    start_year: Date;
    end_year: Date;
}

export interface PostEvaluationPeriodRequest {
    name: string;
    description: string;
    abbreviation: string;
    start_year: string;
    end_year: string;
}