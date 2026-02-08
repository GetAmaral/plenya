import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface ProcessingJob {
  id: string
  batchId: string
  startedAt: number
  lastStep?: number
}

interface ProcessingJobStore {
  activeJobs: ProcessingJob[]
  addJob: (id: string, batchId: string) => void
  removeJob: (id: string) => void
  updateJobStep: (id: string, step: number) => void
  getJob: (id: string) => ProcessingJob | undefined
}

export const useProcessingJobStore = create<ProcessingJobStore>()(
  persist(
    (set, get) => ({
      activeJobs: [],

      addJob: (id, batchId) =>
        set((state) => ({
          activeJobs: [
            ...state.activeJobs.filter((j) => j.id !== id),
            { id, batchId, startedAt: Date.now() },
          ],
        })),

      removeJob: (id) =>
        set((state) => ({
          activeJobs: state.activeJobs.filter((j) => j.id !== id),
        })),

      updateJobStep: (id, step) =>
        set((state) => ({
          activeJobs: state.activeJobs.map((j) =>
            j.id === id ? { ...j, lastStep: step } : j
          ),
        })),

      getJob: (id) => get().activeJobs.find((j) => j.id === id),
    }),
    {
      name: 'processing-jobs-storage',
    }
  )
)
