export const queryKeys = {
  all: ['plenya'] as const,

  me: () => [...queryKeys.all, 'me'] as const,
  meSessions: () => [...queryKeys.me(), 'sessions'] as const,

  mobileConfig: () => [...queryKeys.all, 'mobile-config'] as const,

  patients: {
    all: () => [...queryKeys.all, 'patients'] as const,
    list: (params?: Record<string, unknown>) =>
      [...queryKeys.patients.all(), 'list', params ?? {}] as const,
    detail: (id: string) => [...queryKeys.patients.all(), 'detail', id] as const,
    anamnesis: (id: string) => [...queryKeys.patients.detail(id), 'anamnesis'] as const,
    labResults: (id: string) => [...queryKeys.patients.detail(id), 'lab-results'] as const,
    prescriptions: (id: string) => [...queryKeys.patients.detail(id), 'prescriptions'] as const,
    scores: (id: string) => [...queryKeys.patients.detail(id), 'scores'] as const,
    workoutPlans: (id: string) => [...queryKeys.patients.detail(id), 'workout-plans'] as const,
  },

  scoreGroups: {
    all: () => [...queryKeys.all, 'score-groups'] as const,
    tree: () => [...queryKeys.scoreGroups.all(), 'tree'] as const,
    item: (id: string) => [...queryKeys.scoreGroups.all(), 'item', id] as const,
  },

  anamnesis: {
    all: () => [...queryKeys.all, 'anamnesis'] as const,
    detail: (id: string) => [...queryKeys.anamnesis.all(), id] as const,
    templates: () => [...queryKeys.anamnesis.all(), 'templates'] as const,
  },

  labResults: {
    all: () => [...queryKeys.all, 'lab-results'] as const,
    detail: (id: string) => [...queryKeys.labResults.all(), id] as const,
    definitions: () => [...queryKeys.labResults.all(), 'definitions'] as const,
  },

  prescriptions: {
    all: () => [...queryKeys.all, 'prescriptions'] as const,
    detail: (id: string) => [...queryKeys.prescriptions.all(), id] as const,
  },

  workoutPlans: {
    all: () => [...queryKeys.all, 'workout-plans'] as const,
    detail: (id: string) => [...queryKeys.workoutPlans.all(), id] as const,
    public: (code: string) => [...queryKeys.workoutPlans.all(), 'public', code] as const,
  },

  exercises: () => [...queryKeys.all, 'exercises'] as const,

  leads: {
    all: () => [...queryKeys.all, 'leads'] as const,
    list: (params?: Record<string, unknown>) =>
      [...queryKeys.leads.all(), 'list', params ?? {}] as const,
    detail: (id: string) => [...queryKeys.leads.all(), id] as const,
    dashboard: () => [...queryKeys.leads.all(), 'dashboard'] as const,
  },

  notifications: {
    all: () => [...queryKeys.all, 'notifications'] as const,
    unread: () => [...queryKeys.notifications.all(), 'unread'] as const,
  },

  appointments: {
    all: () => [...queryKeys.all, 'appointments'] as const,
    byRange: (from: string, to: string) =>
      [...queryKeys.appointments.all(), 'range', from, to] as const,
  },

  users: {
    all: () => [...queryKeys.all, 'users'] as const,
    staff: () => [...queryKeys.users.all(), 'staff'] as const,
    doctors: () => [...queryKeys.users.all(), 'doctors'] as const,
  },

  // Endpoints da área do paciente logado (app paciente)
  patientMe: {
    all: () => [...queryKeys.all, 'patient-me'] as const,
    profile: () => [...queryKeys.patientMe.all(), 'profile'] as const,
    dashboard: () => [...queryKeys.patientMe.all(), 'dashboard'] as const,
    appointments: () => [...queryKeys.patientMe.all(), 'appointments'] as const,
    workoutPlans: () => [...queryKeys.patientMe.all(), 'workout-plans'] as const,
    workoutPlanDetail: (id: string) =>
      [...queryKeys.patientMe.workoutPlans(), id] as const,
    workoutSessions: () => [...queryKeys.patientMe.all(), 'workout-sessions'] as const,
    workoutSessionDetail: (id: string) =>
      [...queryKeys.patientMe.workoutSessions(), id] as const,
    checkIns: () => [...queryKeys.patientMe.all(), 'check-ins'] as const,
    checkInToday: () => [...queryKeys.patientMe.checkIns(), 'today'] as const,
    notificationPreferences: () =>
      [...queryKeys.patientMe.all(), 'notification-preferences'] as const,
    labBatches: () => [...queryKeys.patientMe.all(), 'lab-batches'] as const,
    messages: () => [...queryKeys.patientMe.all(), 'messages'] as const,
    scores: () => [...queryKeys.patientMe.all(), 'scores'] as const,
  },
} as const;

export type QueryKeys = typeof queryKeys;
