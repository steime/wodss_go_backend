package persistence

type ModuleVisit struct {
	ID uint `json:"id,omitempty,string" validate:"required"`
	Grade float32 `json:"grade" validate:"min=0,max=6,omitempty"`
	State string `json:"state,omitempty" validate:"oneof=passed failed ongoing planned,omitempty"`
	Student uint `json:"student,omitempty,string" validate:"required"`
	Module string `json:"module,omitempty" validate:"required,len=7"`
	Semester string `json:"semester,omitempty" validate:"required,len=4"`
	Weekday int `json:"weekday" validate:"min=0,max=6,omitempty"`
	TimeStart string `json:"timeStart,omitempty" validate:"required,contains=:,len=5"`
	TimeEnd string `json:"timeEnd,omitempty" validate:"required,contains=:,len=5"`
}

type ModuleVisitCreateBody struct {
	Grade float32 `json:"grade,omitempty" validate:"min=0,max=6,omitempty"`
	State string `json:"state,omitempty" validate:"oneof=passed failed ongoing planned,omitempty"`
	Student uint `json:"student,omitempty,string" validate:"required"`
	Module string `json:"module,omitempty" validate:"required,len=7"`
	Semester string `json:"semester,omitempty" validate:"required,len=4"`
	Weekday int `json:"weekday,omitempty" validate:"min=0,max=6,omitempty"`
	TimeStart string `json:"timeStart,omitempty" validate:"required,contains=:,len=5"`
	TimeEnd string `json:"timeEnd,omitempty" validate:"required,contains=:,len=5"`
}
