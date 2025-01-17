package profile

import (
	"CareerCenter/utils/exceptions"
	"github.com/google/uuid"
	"time"
)

type WorkExperience struct {
	id              string
	email           string
	skillExperience string
	name            string
	stillWorking    bool
	dateRange       DateRangeWork
	description     string
}
type DateRangeWork struct {
	start time.Time
	end   time.Time
}

type WorkExperienceDTO struct {
	Id              string
	Email           string
	SkillExperience string
	Name            string
	StillWorking    bool
	DateRange       DateRangeWorkDTO
	Description     string
}
type DateRangeWorkDTO struct {
	Start time.Time
	End   time.Time
}

func NewWorkExperience(dto *WorkExperienceDTO) (*WorkExperience, error) {
	err := dto.Validation()
	if err != nil {
		return nil, err
	}
	dateRange := NewWorkDateRange(dto.DateRange)
	return &WorkExperience{
		id:              dto.Id,
		email:           dto.Email,
		skillExperience: dto.SkillExperience,
		name:            dto.Name,
		stillWorking:    dto.StillWorking,
		dateRange:       dateRange,
		description:     dto.Description,
	}, nil
}

func NewWorkDateRange(dto DateRangeWorkDTO) DateRangeWork {
	return DateRangeWork{
		start: dto.Start,
		end:   dto.End,
	}
}

func NewListWorkExperience(dto []*WorkExperienceDTO) ([]*WorkExperience, error) {
	listWE := make([]*WorkExperience, 0)
	for _, data := range dto {
		workExperience, err := NewWorkExperience(data)
		if err != nil {
			return nil, err
		}
		listWE = append(listWE, workExperience)
	}
	return listWE, nil
}

func (dto *WorkExperienceDTO) Validation() error {
	if len(dto.Id) <= 0 {
		dto.Id = uuid.NewString()
	}

	if dto.DateRange.Start.IsZero() {
		return exceptions.ErrorStartDate
	}

	if dto.DateRange.End.IsZero() && dto.StillWorking == false {
		return exceptions.ErrCustomString("still working cant be false")
	}

	if !dto.DateRange.Start.IsZero() && !dto.DateRange.End.IsZero() {
		if dto.DateRange.End.Unix() < dto.DateRange.Start.Unix() {
			return exceptions.ErrorEndDate
		}
	}
	return nil
}

func (data *WorkExperience) GetId() string {
	return data.id
}
func (data *WorkExperience) GetEmail() string {
	return data.email
}
func (data *WorkExperience) GetSkillExperience() string {
	return data.skillExperience
}
func (data *WorkExperience) GetName() string {
	return data.name
}
func (data *WorkExperience) GetStillWorking() bool {
	return data.stillWorking
}
func (data *WorkExperience) GetStartWork() time.Time {
	return data.dateRange.start
}
func (data *WorkExperience) GetEndWork() time.Time {
	return data.dateRange.end
}
func (data *WorkExperience) GetDescription() string {
	return data.description
}

func (data *WorkExperience) SetEmail(email string) {
	data.email = email
}
