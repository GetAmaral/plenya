"use client";

import { motion } from "framer-motion";
import {
  Activity,
  Calendar,
  FileText,
  Users,
  Clock,
  TrendingUp,
} from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { AreaChart, BarChart, DonutChart } from "@tremor/react";

// Animation variants
const container = {
  hidden: { opacity: 0 },
  show: {
    opacity: 1,
    transition: {
      staggerChildren: 0.1,
    },
  },
};

const item = {
  hidden: { opacity: 0, y: 20 },
  show: { opacity: 1, y: 0 },
};

// Mock data
const stats = [
  {
    title: "Pacientes Ativos",
    value: "127",
    change: "+12%",
    trend: "up",
    icon: Users,
    color: "text-blue-600",
    bgColor: "bg-blue-50 dark:bg-blue-950/30",
  },
  {
    title: "Consultas Hoje",
    value: "24",
    change: "+8",
    trend: "up",
    icon: Calendar,
    color: "text-emerald-600",
    bgColor: "bg-emerald-50 dark:bg-emerald-950/30",
  },
  {
    title: "Prescrições",
    value: "18",
    change: "Pendentes",
    icon: FileText,
    color: "text-amber-600",
    bgColor: "bg-amber-50 dark:bg-amber-950/30",
  },
  {
    title: "Exames",
    value: "9",
    change: "Novos",
    icon: Activity,
    color: "text-purple-600",
    bgColor: "bg-purple-50 dark:bg-purple-950/30",
  },
];

const recentPatients = [
  {
    name: "João Santos",
    status: "stable",
    time: "10:30",
    avatar: "JS",
  },
  {
    name: "Maria Silva",
    status: "observation",
    time: "11:00",
    avatar: "MS",
  },
  {
    name: "Pedro Costa",
    status: "stable",
    time: "14:00",
    avatar: "PC",
  },
];

const upcomingAppointments = [
  { patient: "Ana Oliveira", time: "15:00", type: "Consulta" },
  { patient: "Carlos Mendes", time: "15:30", type: "Retorno" },
  { patient: "Lucia Ferreira", time: "16:00", type: "Exame" },
];

// Chart data
const patientTrendData = [
  { month: "Jan", pacientes: 98 },
  { month: "Fev", pacientes: 102 },
  { month: "Mar", pacientes: 115 },
  { month: "Abr", pacientes: 118 },
  { month: "Mai", pacientes: 123 },
  { month: "Jun", pacientes: 127 },
];

const appointmentsByTypeData = [
  { tipo: "Consulta", quantidade: 42 },
  { tipo: "Retorno", quantidade: 28 },
  { tipo: "Exame", quantidade: 18 },
  { tipo: "Emergência", quantidade: 12 },
];

const patientStatusData = [
  { status: "Estável", pacientes: 87 },
  { status: "Observação", pacientes: 32 },
  { status: "Crítico", pacientes: 8 },
];

export default function DashboardPage() {
  return (
    <div className="min-h-screen p-6">
      <div className="mx-auto max-w-7xl">
        {/* Header */}
        <motion.div
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          className="mb-8"
        >
          <h1 className="text-3xl font-bold tracking-tight text-gray-900 dark:text-white">
            Dashboard
          </h1>
          <p className="mt-2 text-gray-600 dark:text-gray-400">
            Bem-vindo de volta! Aqui está um resumo do seu dia.
          </p>
        </motion.div>

        {/* Stats Grid - Bento Layout */}
        <motion.div
          variants={container}
          initial="hidden"
          animate="show"
          className="grid gap-6 md:grid-cols-2 lg:grid-cols-4 mb-6"
        >
          {stats.map((stat) => {
            const Icon = stat.icon;
            return (
              <motion.div key={stat.title} variants={item}>
                <Card className="group relative overflow-hidden border-0 shadow-md hover:shadow-xl transition-all duration-300">
                  <CardContent className="p-6">
                    <div className="flex items-center justify-between">
                      <div>
                        <p className="text-sm font-medium text-muted-foreground">
                          {stat.title}
                        </p>
                        <div className="flex items-baseline gap-2 mt-2">
                          <h3 className="text-3xl font-bold">{stat.value}</h3>
                          {stat.change && (
                            <span
                              className={`text-sm font-medium ${
                                stat.trend === "up"
                                  ? "text-emerald-600"
                                  : "text-gray-600"
                              }`}
                            >
                              {stat.change}
                            </span>
                          )}
                        </div>
                      </div>
                      <div
                        className={`${stat.bgColor} ${stat.color} p-3 rounded-2xl group-hover:scale-110 transition-transform`}
                      >
                        <Icon className="h-6 w-6" />
                      </div>
                    </div>
                  </CardContent>
                </Card>
              </motion.div>
            );
          })}
        </motion.div>

        {/* Charts Section */}
        <motion.div
          variants={container}
          initial="hidden"
          animate="show"
          className="grid gap-6 lg:grid-cols-3 mb-6"
        >
          {/* Patient Trend Chart */}
          <motion.div variants={item} className="lg:col-span-2">
            <Card className="border-0 shadow-md">
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <TrendingUp className="h-5 w-5 text-blue-600" />
                  Crescimento de Pacientes
                </CardTitle>
              </CardHeader>
              <CardContent>
                <AreaChart
                  className="h-72"
                  data={patientTrendData}
                  index="month"
                  categories={["pacientes"]}
                  colors={["blue"]}
                  valueFormatter={(number: number) =>
                    `${Intl.NumberFormat("pt-BR").format(number)} pacientes`
                  }
                  showLegend={false}
                  showGridLines={false}
                  curveType="monotone"
                />
              </CardContent>
            </Card>
          </motion.div>

          {/* Patient Status Donut Chart */}
          <motion.div variants={item}>
            <Card className="border-0 shadow-md">
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <Activity className="h-5 w-5 text-purple-600" />
                  Status dos Pacientes
                </CardTitle>
              </CardHeader>
              <CardContent>
                <DonutChart
                  className="h-72"
                  data={patientStatusData}
                  category="pacientes"
                  index="status"
                  colors={["emerald", "amber", "red"]}
                  valueFormatter={(number: number) =>
                    `${Intl.NumberFormat("pt-BR").format(number)} pacientes`
                  }
                  showAnimation={true}
                />
              </CardContent>
            </Card>
          </motion.div>

          {/* Appointments by Type Chart */}
          <motion.div variants={item} className="lg:col-span-3">
            <Card className="border-0 shadow-md">
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <Calendar className="h-5 w-5 text-emerald-600" />
                  Consultas por Tipo
                </CardTitle>
              </CardHeader>
              <CardContent>
                <BarChart
                  className="h-64"
                  data={appointmentsByTypeData}
                  index="tipo"
                  categories={["quantidade"]}
                  colors={["emerald"]}
                  valueFormatter={(number: number) =>
                    `${Intl.NumberFormat("pt-BR").format(number)} consultas`
                  }
                  showLegend={false}
                  showGridLines={false}
                />
              </CardContent>
            </Card>
          </motion.div>
        </motion.div>

        {/* Bento Grid - Main Content */}
        <div className="grid gap-6 lg:grid-cols-3">
          {/* Recent Patients - Large Card */}
          <motion.div
            variants={item}
            initial="hidden"
            animate="show"
            className="lg:col-span-2"
          >
            <Card className="h-full border-0 shadow-md">
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <Users className="h-5 w-5 text-blue-600" />
                  Pacientes Recentes
                </CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {recentPatients.map((patient, index) => (
                    <motion.div
                      key={patient.name}
                      initial={{ opacity: 0, x: -20 }}
                      animate={{ opacity: 1, x: 0 }}
                      transition={{ delay: index * 0.1 }}
                      className="flex items-center justify-between p-4 rounded-lg border border-border hover:bg-accent/50 transition-colors"
                    >
                      <div className="flex items-center gap-3">
                        <Avatar className="h-10 w-10">
                          <AvatarFallback className="bg-gradient-to-br from-blue-500 to-blue-700 text-white">
                            {patient.avatar}
                          </AvatarFallback>
                        </Avatar>
                        <div>
                          <p className="font-medium">{patient.name}</p>
                          <p className="text-sm text-muted-foreground">
                            Última consulta: {patient.time}
                          </p>
                        </div>
                      </div>
                      <Badge
                        variant={
                          patient.status === "stable" ? "stable" : "observation"
                        }
                      >
                        {patient.status === "stable"
                          ? "Estável"
                          : "Observação"}
                      </Badge>
                    </motion.div>
                  ))}
                </div>
              </CardContent>
            </Card>
          </motion.div>

          {/* Upcoming Appointments */}
          <motion.div
            variants={item}
            initial="hidden"
            animate="show"
            transition={{ delay: 0.2 }}
          >
            <Card className="h-full border-0 shadow-md">
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <Calendar className="h-5 w-5 text-emerald-600" />
                  Próximas Consultas
                </CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {upcomingAppointments.map((apt, index) => (
                    <motion.div
                      key={apt.patient}
                      initial={{ opacity: 0, scale: 0.9 }}
                      animate={{ opacity: 1, scale: 1 }}
                      transition={{ delay: 0.3 + index * 0.1 }}
                      className="flex items-start gap-3 p-3 rounded-lg bg-accent/30"
                    >
                      <div className="flex h-10 w-10 items-center justify-center rounded-full bg-emerald-100 dark:bg-emerald-950/30">
                        <Clock className="h-5 w-5 text-emerald-600" />
                      </div>
                      <div className="flex-1">
                        <p className="font-medium text-sm">{apt.patient}</p>
                        <p className="text-xs text-muted-foreground">
                          {apt.type}
                        </p>
                        <p className="text-xs font-medium text-emerald-600 mt-1">
                          {apt.time}
                        </p>
                      </div>
                    </motion.div>
                  ))}
                </div>
              </CardContent>
            </Card>
          </motion.div>
        </div>

        {/* Quick Actions */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.5 }}
          className="mt-6"
        >
          <Card className="border-0 shadow-md bg-gradient-to-br from-blue-500 to-blue-700 text-white">
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <h3 className="text-lg font-semibold">Ações Rápidas</h3>
                  <p className="text-sm text-blue-100 mt-1">
                    Acesse rapidamente as funcionalidades principais
                  </p>
                </div>
                <div className="flex gap-2">
                  <button className="px-4 py-2 bg-white/20 hover:bg-white/30 rounded-lg text-sm font-medium transition-colors">
                    Novo Paciente
                  </button>
                  <button className="px-4 py-2 bg-white/20 hover:bg-white/30 rounded-lg text-sm font-medium transition-colors">
                    Nova Consulta
                  </button>
                </div>
              </div>
            </CardContent>
          </Card>
        </motion.div>
      </div>
    </div>
  );
}
