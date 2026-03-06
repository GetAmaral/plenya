package services

// NSCAReferenceData contém dados de referência NSCA para treino
var NSCAReferenceData = map[string]interface{}{
	"zonas_repeticao": map[string]interface{}{
		"1-5_RM": map[string]interface{}{
			"objetivo":    "Força máxima",
			"descanso":    "2-5 min",
			"series":      "3-5",
			"intensidade": "85-100% 1RM",
		},
		"6-12_RM": map[string]interface{}{
			"objetivo":    "Hipertrofia",
			"descanso":    "30-90s",
			"series":      "3-6",
			"intensidade": "67-85% 1RM",
		},
		"12+_RM": map[string]interface{}{
			"objetivo":    "Resistência muscular",
			"descanso":    "30s ou menos",
			"series":      "2-3",
			"intensidade": "<67% 1RM",
		},
	},
	"intervalos_descanso": map[string]interface{}{
		"forca":       "2-5 min",
		"potencia":    "2-5 min",
		"hipertrofia": "30-90s",
		"resistencia": "30s ou menos",
	},
	"frequencia_treino": map[string]interface{}{
		"iniciantes":    "2-3x/semana Full Body",
		"intermediarios": "3-4x/semana Upper/Lower ou Push/Pull/Legs",
		"avancados":     "4-6x/semana Split por grupo muscular",
	},
	"progressao_carga": map[string]interface{}{
		"iniciantes": map[string]interface{}{
			"upper_body": "2.5-5 lbs (1-2.5 kg) por sessão",
			"lower_body": "5-10 lbs (2.5-5 kg) por sessão",
		},
		"avancados": map[string]interface{}{
			"upper_body": "2.5-5 lbs (1-2.5 kg) por semana",
			"lower_body": "5-10 lbs (2.5-5 kg) por semana",
		},
	},
	"volume_semanal": map[string]interface{}{
		"iniciantes":    "10-20 séries por grupo muscular/semana",
		"intermediarios": "15-25 séries por grupo muscular/semana",
		"avancados":     "20-30+ séries por grupo muscular/semana",
	},
}
