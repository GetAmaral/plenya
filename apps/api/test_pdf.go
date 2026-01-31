package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/services"
)

func main() {
	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repositories and services
	patientRepo := repository.NewPatientRepository(database.DB)
	labRequestRepo := repository.NewLabRequestRepository(database.DB)
	pdfService := services.NewPDFService()
	labRequestService := services.NewLabRequestService(labRequestRepo, patientRepo, pdfService)

	// Get first lab request
	id := uuid.MustParse("019c11bd-03f9-7a3d-b625-ab7f30023ae8")

	req, err := labRequestService.GetLabRequestByID(id)
	if err != nil {
		log.Fatalf("Failed to get lab request: %v", err)
	}

	fmt.Printf("Generating PDF for lab request: %s\n", id)
	fmt.Printf("Patient: %s\n", req.Patient.Name)
	fmt.Printf("Date: %s\n", req.Date.Format("2006-01-02"))

	// Generate PDF
	pdfURL, err := pdfService.GenerateLabRequestPDF(req)
	if err != nil {
		log.Fatalf("Failed to generate PDF: %v", err)
	}

	fmt.Printf("✅ PDF generated: %s\n", pdfURL)

	// Check file size
	pdfPath := fmt.Sprintf("/app/uploads/lab-requests/%s", pdfURL[len("/uploads/lab-requests/"):])
	info, err := os.Stat(pdfPath)
	if err != nil {
		log.Fatalf("Failed to stat PDF file: %v", err)
	}

	sizeKB := float64(info.Size()) / 1024.0
	sizeMB := sizeKB / 1024.0

	fmt.Printf("📊 File size: %.2f KB (%.2f MB)\n", sizeKB, sizeMB)

	if sizeMB > 2.0 {
		fmt.Printf("⚠️  WARNING: PDF is larger than 2MB\n")
	} else {
		fmt.Printf("✅ File size acceptable\n")
	}
}
