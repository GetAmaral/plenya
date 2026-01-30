/**
 * API Client for User endpoints
 */

import { apiClient } from "@/lib/api-client";
import { User } from "@/lib/auth-store";

export const userApi = {
  /**
   * Get current user with selectedPatient populated
   */
  getMe: async (): Promise<User> => {
    return apiClient.get<User>("/api/v1/users/me");
  },

  /**
   * Update selected patient for current user
   */
  updateSelectedPatient: async (patientId: string): Promise<User> => {
    return apiClient.put<User>("/api/v1/users/me/selected-patient", {
      patientId,
    });
  },
};
