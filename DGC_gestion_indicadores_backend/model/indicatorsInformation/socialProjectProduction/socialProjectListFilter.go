package model

type FilterSocialProjectListsByAcademicPeriodRequest struct {
	AcademicPeriodID int    `json:"academic_period_id"`
	Career           string `json:"career"`
	Name             string `json:"name"`
	PageSize         int    `json:"page_size"`
	Page             int    `json:"page"`
}

type SocialProjectListByAcademicPeriodJoined struct {
	AcademicPeriodID int    `json:"academic_period_id"`
	CareerID         uint   `json:"career_id"`
	Career           string `json:"career"`
	Name             string `json:"name"`
}

type FilterSocialProjectListsByAcademicPeriodResponse struct {
	Count              int64                                     `json:"count"`
	PageSize           int                                       `json:"page_size"`
	Page               int                                       `json:"page"`
	SocialProjectLists []SocialProjectListByAcademicPeriodJoined `json:"social_project_lists"`
}
