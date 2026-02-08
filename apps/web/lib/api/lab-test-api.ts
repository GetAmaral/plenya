/**
 * API Client for Lab Test System
 * Handles Lab Test Definitions, Lab Result Values, and Score Mappings
 */

import { apiClient } from "@/lib/api-client";

// ============================================================
// Types
// ============================================================

export type LabTestCategory =
  | "hematology"
  | "biochemistry"
  | "hormones"
  | "immunology"
  | "microbiology"
  | "urine"
  | "imaging"
  | "functional"
  | "genetics"
  | "other";

export type ResultType = "numeric" | "text" | "boolean" | "categorical";

export type Gender = "male" | "female" | "other";

export interface LabTestDefinition {
  id: string;
  code: string;
  name: string;
  shortName?: string;
  tussCode?: string;
  loincCode?: string;
  category: LabTestCategory;
  isRequestable: boolean;
  parentTestId?: string;
  unit?: string;
  unitConversion?: string;
  resultType: ResultType;
  collectionMethod?: string;
  fastingHours?: number;
  specimenType?: string;
  description?: string;
  clinicalIndications?: string;
  displayOrder: number;
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
  parentTest?: LabTestDefinition;
  subTests?: LabTestDefinition[];
}

export interface LabResultValue {
  id: string;
  labResultId: string;
  labTestDefinitionId: string;
  numericValue?: number;
  textValue?: string;
  booleanValue?: boolean;
  unit?: string;
  notes?: string;
  createdAt: string;
  updatedAt: string;
  labTestDefinition?: LabTestDefinition;
}

export interface CreateLabTestDefinitionDTO {
  code: string;
  name: string;
  shortName?: string;
  tussCode?: string;
  loincCode?: string;
  category: LabTestCategory;
  isRequestable: boolean;
  parentTestId?: string;
  unit?: string;
  unitConversion?: string;
  resultType: ResultType;
  collectionMethod?: string;
  fastingHours?: number;
  specimenType?: string;
  description?: string;
  clinicalIndications?: string;
  displayOrder?: number;
  isActive?: boolean;
}

export interface CreateLabResultValueDTO {
  labResultId: string;
  labTestDefinitionId: string;
  numericValue?: number;
  textValue?: string;
  booleanValue?: boolean;
  unit?: string;
  notes?: string;
}


// ============================================================
// Lab Test Definitions API
// ============================================================

export const labTestDefinitionsApi = {
  /**
   * Get all lab test definitions
   */
  getAll: async (): Promise<LabTestDefinition[]> => {
    return apiClient.get<LabTestDefinition[]>(
      "/api/v1/lab-tests/definitions"
    );
  },

  /**
   * Get only requestable lab tests (can be ordered)
   */
  getRequestable: async (
    category?: LabTestCategory
  ): Promise<LabTestDefinition[]> => {
    const params = category ? `?category=${category}` : "";
    return apiClient.get<LabTestDefinition[]>(
      `/api/v1/lab-tests/requestable${params}`
    );
  },

  /**
   * Get lab test definition by ID
   */
  getById: async (id: string): Promise<LabTestDefinition> => {
    return apiClient.get<LabTestDefinition>(
      `/api/v1/lab-tests/definitions/${id}`
    );
  },

  /**
   * Get lab test definition by code
   */
  getByCode: async (code: string): Promise<LabTestDefinition> => {
    return apiClient.get<LabTestDefinition>(
      `/api/v1/lab-tests/definitions/code/${code}`
    );
  },

  /**
   * Get sub-tests of a parent test
   */
  getSubTests: async (parentId: string): Promise<LabTestDefinition[]> => {
    return apiClient.get<LabTestDefinition[]>(
      `/api/v1/lab-tests/definitions/${parentId}/sub-tests`
    );
  },

  /**
   * Search lab test definitions by name or code
   */
  search: async (query: string): Promise<LabTestDefinition[]> => {
    return apiClient.get<LabTestDefinition[]>(
      `/api/v1/lab-tests/definitions/search?q=${encodeURIComponent(query)}`
    );
  },

  /**
   * Create new lab test definition (admin only)
   */
  create: async (
    data: CreateLabTestDefinitionDTO
  ): Promise<LabTestDefinition> => {
    return apiClient.post<LabTestDefinition>(
      "/api/v1/lab-tests/definitions",
      data
    );
  },

  /**
   * Update lab test definition (admin only)
   */
  update: async (
    id: string,
    data: Partial<CreateLabTestDefinitionDTO>
  ): Promise<LabTestDefinition> => {
    return apiClient.put<LabTestDefinition>(
      `/api/v1/lab-tests/definitions/${id}`,
      data
    );
  },

  /**
   * Delete lab test definition (admin only)
   */
  delete: async (id: string): Promise<void> => {
    return apiClient.delete(`/api/v1/lab-tests/definitions/${id}`);
  },
};

// ============================================================
// Lab Result Values API
// ============================================================

export const labResultValuesApi = {
  /**
   * Create single lab result value
   */
  create: async (data: CreateLabResultValueDTO): Promise<LabResultValue> => {
    return apiClient.post<LabResultValue>("/api/v1/lab-results/values", data);
  },

  /**
   * Create multiple lab result values in batch
   */
  createBatch: async (
    data: CreateLabResultValueDTO[]
  ): Promise<LabResultValue[]> => {
    return apiClient.post<LabResultValue[]>(
      "/api/v1/lab-results/values/batch",
      data
    );
  },

  /**
   * Get lab result value by ID
   */
  getById: async (id: string): Promise<LabResultValue> => {
    return apiClient.get<LabResultValue>(`/api/v1/lab-results/values/${id}`);
  },

  /**
   * Get all values for a specific lab result
   */
  getByLabResult: async (labResultId: string): Promise<LabResultValue[]> => {
    return apiClient.get<LabResultValue[]>(
      `/api/v1/lab-results/${labResultId}/values`
    );
  },

  /**
   * Get all lab values for a patient
   */
  getByPatient: async (patientId: string): Promise<LabResultValue[]> => {
    return apiClient.get<LabResultValue[]>(
      `/api/v1/patients/${patientId}/lab-values`
    );
  },

  /**
   * Get latest value for a specific test and patient
   */
  getLatestForTest: async (
    patientId: string,
    testId: string
  ): Promise<LabResultValue> => {
    return apiClient.get<LabResultValue>(
      `/api/v1/patients/${patientId}/lab-values/test/${testId}/latest`
    );
  },

  /**
   * Update lab result value
   */
  update: async (
    id: string,
    data: Partial<CreateLabResultValueDTO>
  ): Promise<LabResultValue> => {
    return apiClient.put<LabResultValue>(
      `/api/v1/lab-results/values/${id}`,
      data
    );
  },

  /**
   * Delete lab result value (admin only)
   */
  delete: async (id: string): Promise<void> => {
    return apiClient.delete(`/api/v1/lab-results/values/${id}`);
  },
};


// ============================================================
// Lab Results API (Parent Container)
// ============================================================

export interface LabResult {
  id: string;
  patientId: string;
  doctorId: string;
  testType: string;
  testDate: string;
  results?: string;
  normalRange?: string;
  status: "pending" | "completed" | "reviewed";
  notes?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateLabResultDTO {
  patientId: string;
  doctorId: string;
  testType: string;
  testDate: string;
  notes?: string;
  status?: "pending" | "completed" | "reviewed";
}

export const labResultsApi = {
  /**
   * Create a lab result container
   */
  create: async (data: CreateLabResultDTO): Promise<LabResult> => {
    return apiClient.post<LabResult>("/api/v1/lab-results", data);
  },

  /**
   * Create lab result with structured values in one call
   */
  createWithValues: async (data: {
    labResult: CreateLabResultDTO;
    values: CreateLabResultValueDTO[];
  }): Promise<{ labResult: LabResult; values: LabResultValue[] }> => {
    // First create the lab result
    const labResult = await apiClient.post<LabResult>(
      "/api/v1/lab-results",
      data.labResult
    );

    // Then create values in batch
    const valuesWithLabResultId = data.values.map((v) => ({
      ...v,
      labResultId: labResult.id,
    }));

    const values = await apiClient.post<LabResultValue[]>(
      "/api/v1/lab-results/values/batch",
      valuesWithLabResultId
    );

    return { labResult, values };
  },
};

// ============================================================
// Helper Functions
// ============================================================

export const labTestHelpers = {
  /**
   * Get category label in Portuguese
   */
  getCategoryLabel: (category: LabTestCategory): string => {
    const labels: Record<LabTestCategory, string> = {
      hematology: "Hematologia",
      biochemistry: "Bioquímica",
      hormones: "Hormônios",
      immunology: "Imunologia",
      microbiology: "Microbiologia",
      urine: "Urinálise",
      imaging: "Imagem",
      functional: "Medicina Funcional",
      genetics: "Genética",
      other: "Outros",
    };
    return labels[category];
  },

  /**
   * Get result type label in Portuguese
   */
  getResultTypeLabel: (resultType: ResultType): string => {
    const labels: Record<ResultType, string> = {
      numeric: "Numérico",
      text: "Texto",
      boolean: "Sim/Não",
      categorical: "Categórico",
    };
    return labels[resultType];
  },

  /**
   * Format lab result value for display
   */
  formatValue: (value: LabResultValue): string => {
    if (value.numericValue !== undefined && value.numericValue !== null) {
      const formatted = value.numericValue.toLocaleString("pt-BR", {
        maximumFractionDigits: 2,
      });
      return value.unit ? `${formatted} ${value.unit}` : formatted;
    }
    if (value.textValue) return value.textValue;
    if (value.booleanValue !== undefined && value.booleanValue !== null) {
      return value.booleanValue ? "Sim" : "Não";
    }
    return "—";
  },

  /**
   * Build hierarchical tree from flat lab test definitions
   */
  buildTree: (tests: LabTestDefinition[]): LabTestDefinition[] => {
    const testMap = new Map(tests.map((t) => [t.id, { ...t, subTests: [] }]));
    const roots: LabTestDefinition[] = [];

    tests.forEach((test) => {
      const node = testMap.get(test.id)!;
      if (test.parentTestId) {
        const parent = testMap.get(test.parentTestId);
        if (parent) {
          parent.subTests = parent.subTests || [];
          parent.subTests.push(node);
        }
      } else {
        roots.push(node);
      }
    });

    // Sort by display order
    const sortByOrder = (arr: LabTestDefinition[]) => {
      arr.sort((a, b) => a.displayOrder - b.displayOrder);
      arr.forEach((item) => {
        if (item.subTests && item.subTests.length > 0) {
          sortByOrder(item.subTests);
        }
      });
    };

    sortByOrder(roots);
    return roots;
  },
};
