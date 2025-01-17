package profile

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

func (u UseCaseProfileInteractor) GetProfileByEmail(ctx context.Context, email string) (*profile.ProfileUserDTO, error) {
	data, err := u.repoProfile.GetProfileByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	dataWorkExperience, err := u.repoProfile.GetListWorkExperience(ctx, email)
	if err != nil {
		return nil, err
	}

	dataEducation, err := u.repoProfile.GetListEducation(ctx, email)
	if err != nil {
		return nil, err
	}

	data.SetWorkExperience(dataWorkExperience)
	data.SetEducation(dataEducation)

	return data, nil
}
