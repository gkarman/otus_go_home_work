package responsedto

import "github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/entity"

type Events struct {
	Events []entity.Event `json:"events"`
}
