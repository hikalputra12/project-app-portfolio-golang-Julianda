package repository

import (
	"project-portofolio/database"

	"go.uber.org/zap"
)

type Repository struct {
	Biodata    BiodataRepository
	About      AboutRepository
	Resume     ResumeRepository
	Portofolio PortofolioRepository
	Message    MessageRepository
	Skill      SkillRepository
}

func NewRepository(db database.PgxIface, logger *zap.Logger) Repository {
	return Repository{
		Biodata:    NewBiodataRepository(db, logger),
		About:      NewAboutRepository(db, logger),
		Resume:     NewResumeRepository(db, logger),
		Portofolio: NewPortofolioRepository(db, logger),
		Message:    NewMessageRepository(db, logger),
		Skill:      NewSkillRepository(db, logger),
	}
}
