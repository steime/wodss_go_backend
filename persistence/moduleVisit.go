package persistence

type ModuleVisit struct {
	ID uint `json:"id,omitempty,string"`
	Grade float32 `json:"grade"`
	State string `json:"state,omitempty" `
	Student uint `json:"student,omitempty,string" `
	Module string `json:"module,omitempty" `
	Semester string `json:"semester,omitempty" `
	Weekday int `json:"weekday" `
	TimeStart string `json:"timeStart,omitempty" `
	TimeEnd string `json:"timeEnd,omitempty" `
}

type ModuleVisitCreateBody struct {
	Grade float32 `json:"grade,omitempty" validate:"min=0,max=6,omitempty"`
	State string `json:"state,omitempty" validate:"required"`
	Student uint `json:"student,omitempty,string" validate:"required"`
	Module string `json:"module,omitempty" validate:"required"`
	Semester string `json:"semester,omitempty" validate:"required"`
	Weekday int `json:"weekday,omitempty" validate:"min=0,max=6,omitempty"`
	TimeStart string `json:"timeStart,omitempty" validate:"required"`
	TimeEnd string `json:"timeEnd,omitempty" validate:"required"`
}
