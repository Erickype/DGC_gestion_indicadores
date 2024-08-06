package model

type CommonDeleteResponse struct {
	Deleted    bool   `json:"deleted"`
	EntityName string `json:"entity_name"`
	IDDeleted  int    `json:"id_deleted"`
}

func CreateCommonDeleteResponse(EntityName string, IDDeleted int) *CommonDeleteResponse {
	return &CommonDeleteResponse{
		Deleted:    true,
		EntityName: EntityName,
		IDDeleted:  IDDeleted,
	}
}
