package models

// CIDCode é o catálogo (curado) de códigos CID-10 para autocomplete no problem
// list. Não é exaustivo (a ~14k oficial fica para um seed futuro); cobre as
// condições crônicas/preventivas mais comuns. Code é a PK (sem UUID).
//
// A descrição do problema (PatientProblem.Description) é texto livre e o código
// é opcional — então a completude deste catálogo não bloqueia o uso.
type CIDCode struct {
	Code        string `gorm:"type:varchar(10);primaryKey" json:"code"`
	Description string `gorm:"type:varchar(300);not null" json:"description"`
	Chapter     string `gorm:"type:varchar(160)" json:"chapter,omitempty"`
	Version     string `gorm:"type:varchar(8);not null;default:'10'" json:"version"`
}

func (CIDCode) TableName() string { return "cid_codes" }
