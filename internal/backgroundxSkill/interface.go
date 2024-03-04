package backgroundxSkill

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryBackgroundXSkills interface {
	CreateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error)
	GetAllBackgroundXSkills() ([]domain.BackgroundXSkills, error)
	GetByIdBackgroundXSkills(id int64) (domain.BackgroundXSkills, error)
	GetByBackgroundId(id int64) ([]domain.BackgroundXSkills, error)
	UpdateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error)
	DeleteBackgroundXSkills(id int64) error
	DeleteByBackgroundId(id int64) error
}

type ServiceBackgroundXSkills interface {
	CreateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error)
	GetAllBackgroundXSkills() ([]domain.BackgroundXSkills, error)
	GetByIdBackgroundXSkills(id int64) (domain.BackgroundXSkills, error)
	GetByBackgroundId(id int64) ([]domain.BackgroundXSkills, error)
	UpdateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error)
	DeleteBackgroundXSkills(id int64) error
	DeleteByBackgroundId(id int64) error
}
