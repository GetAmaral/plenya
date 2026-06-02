package dto

// CreateConsultationVitalsRequest — payload de aferição de sinais vitais.
// AppointmentID ancora à consulta; MeasuredAt default = agora.
type CreateConsultationVitalsRequest struct {
	AppointmentID      *string  `json:"appointmentId,omitempty" validate:"omitempty,uuid"`
	SystolicBP         *int     `json:"systolicBp,omitempty"`
	DiastolicBP        *int     `json:"diastolicBp,omitempty"`
	HeartRate          *int     `json:"heartRate,omitempty"`
	RespRate           *int     `json:"respRate,omitempty"`
	Temperature        *float64 `json:"temperature,omitempty"`
	SpO2               *int     `json:"spo2,omitempty"`
	Weight             *float64 `json:"weight,omitempty"`
	Height             *float64 `json:"height,omitempty"`
	WaistCircumference *float64 `json:"waistCircumference,omitempty"`
	MeasuredAt         *string  `json:"measuredAt,omitempty"` // RFC3339; default agora
}

// ConsultationVitalsResponse
type ConsultationVitalsResponse struct {
	ID                 string   `json:"id"`
	AppointmentID      *string  `json:"appointmentId,omitempty"`
	PatientID          string   `json:"patientId"`
	SystolicBP         *int     `json:"systolicBp,omitempty"`
	DiastolicBP        *int     `json:"diastolicBp,omitempty"`
	HeartRate          *int     `json:"heartRate,omitempty"`
	RespRate           *int     `json:"respRate,omitempty"`
	Temperature        *float64 `json:"temperature,omitempty"`
	SpO2               *int     `json:"spo2,omitempty"`
	Weight             *float64 `json:"weight,omitempty"`
	Height             *float64 `json:"height,omitempty"`
	WaistCircumference *float64 `json:"waistCircumference,omitempty"`
	BMI                *float64 `json:"bmi,omitempty"`
	MeasuredByUserID   string   `json:"measuredByUserId"`
	MeasuredByName     string   `json:"measuredByName,omitempty"`
	MeasuredAt         string   `json:"measuredAt"`
	CreatedAt          string   `json:"createdAt"`
}
