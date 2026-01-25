package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Article representa um artigo científico armazenado no sistema
type Article struct {
	// Primary Key
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// Metadados Bibliográficos Essenciais
	// @minLength 3
	// @maxLength 1000
	Title string `gorm:"type:varchar(1000);not null;index" json:"title" validate:"required,min=3,max=1000"`

	// @minLength 1
	// @maxLength 2000
	Authors string `gorm:"type:varchar(2000);not null" json:"authors" validate:"required,min=1,max=2000"`

	// @minLength 1
	// @maxLength 500
	Journal string `gorm:"type:varchar(500);not null;index" json:"journal" validate:"required,min=1,max=500"`

	// @example 2024-03-15
	PublishDate time.Time `gorm:"type:date;not null;index" json:"publishDate" validate:"required"`

	// @minLength 2
	// @maxLength 10
	// @example en
	Language string `gorm:"type:varchar(10);not null;default:'en'" json:"language" validate:"required,min=2,max=10"`

	// Identificadores Únicos
	// @example 10.1038/s41586-024-07146-0
	// @maxLength 255
	DOI *string `gorm:"type:varchar(255);uniqueIndex" json:"doi,omitempty" validate:"omitempty,max=255"`

	// @example 38355808
	// @maxLength 50
	PMID *string `gorm:"type:varchar(50);index" json:"pmid,omitempty" validate:"omitempty,max=50"`

	// @example 0028-0836
	// @maxLength 20
	ISSN *string `gorm:"type:varchar(20)" json:"issn,omitempty" validate:"omitempty,max=20"`

	// Conteúdo
	Abstract    *string `gorm:"type:text" json:"abstract,omitempty"`
	FullContent *string `gorm:"type:text" json:"fullContent,omitempty"`
	Notes       *string `gorm:"type:text" json:"notes,omitempty"` // Notas pessoais do médico

	// Links
	// @example https://www.nature.com/articles/s41586-024-07146-0
	OriginalLink *string `gorm:"type:varchar(2048)" json:"originalLink,omitempty" validate:"omitempty,url,max=2048"`

	// @example /uploads/articles/550e8400-e29b-41d4-a716-446655440000.pdf
	InternalLink *string `gorm:"type:varchar(2048)" json:"internalLink,omitempty" validate:"omitempty,max=2048"`

	// Classificação
	// @enum research_article,review,meta_analysis,case_study,clinical_trial,editorial,letter,protocol,lecture
	ArticleType string `gorm:"type:varchar(50);not null;default:'research_article';check:article_type IN ('research_article','review','meta_analysis','case_study','clinical_trial','editorial','letter','protocol','lecture')" json:"articleType" validate:"required,oneof=research_article review meta_analysis case_study clinical_trial editorial letter protocol lecture"`

	// @items.type string
	Keywords []string `gorm:"type:jsonb" json:"keywords,omitempty"`

	// @items.type string
	MeshTerms []string `gorm:"type:jsonb" json:"meshTerms,omitempty"`

	// @example Cardiology
	// @maxLength 200
	Specialty *string `gorm:"type:varchar(200);index" json:"specialty,omitempty" validate:"omitempty,max=200"`

	// Gestão Pessoal
	Favorite bool `gorm:"default:false;index" json:"favorite"`

	// @minimum 0
	// @maximum 5
	// @example 4
	Rating *int `gorm:"type:smallint;check:rating >= 0 AND rating <= 5" json:"rating,omitempty" validate:"omitempty,min=0,max=5"`

	// Arquivo
	// @example a1b2c3d4e5f6...
	// @maxLength 64
	FileHash *string `gorm:"type:varchar(64);index" json:"fileHash,omitempty" validate:"omitempty,len=64"` // SHA-256

	// @example 2457600
	FileSize *int64 `gorm:"type:bigint" json:"fileSize,omitempty" validate:"omitempty,min=0"`

	// Metadados do Sistema
	IndexedAt      time.Time  `gorm:"autoCreateTime" json:"indexedAt"`
	LastAccessedAt *time.Time `gorm:"type:timestamp" json:"lastAccessedAt,omitempty"`

	// Audit
	CreatedBy *uuid.UUID `gorm:"type:uuid" json:"createdBy,omitempty"` // User ID que criou
	UpdatedBy *uuid.UUID `gorm:"type:uuid" json:"updatedBy,omitempty"` // User ID que atualizou

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

// TableName especifica o nome da tabela no banco de dados
func (Article) TableName() string {
	return "articles"
}

// BeforeCreate hook - executado antes de criar registro
func (a *Article) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}
