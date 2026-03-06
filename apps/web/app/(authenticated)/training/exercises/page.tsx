"use client";

import { useState } from "react";
import { motion } from "framer-motion";
import { Search, Library } from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Button } from "@/components/ui/button";
import { useRequireAuth } from "@/lib/use-auth";
import { PageHeader } from "@/components/layout/page-header";
import { useExercises, type ExerciseFilters } from "@/lib/api/exercise-api";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001";

export default function ExercisesPage() {
  useRequireAuth();
  const [search, setSearch] = useState("");
  const [filters, setFilters] = useState<ExerciseFilters>({ limit: 20, offset: 0 });

  const { data, isLoading } = useExercises({ ...filters, search: search || undefined });

  return (
    <div className="space-y-6">
      <PageHeader
        title="Exercícios"
        subtitle={`${data?.total ?? 0} exercícios disponíveis`}
        icon={Library}
      />

      <div className="flex gap-4">
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            placeholder="Buscar exercícios..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="pl-10"
          />
        </div>
      </div>

      {isLoading ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
          {Array.from({ length: 8 }).map((_, i) => (
            <Skeleton key={i} className="h-64" />
          ))}
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
          {data?.exercises?.map((exercise) => (
            <motion.div
              key={exercise.id}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
            >
              <Card className="overflow-hidden hover:shadow-lg transition-shadow">
                <div className="aspect-square bg-muted flex items-center justify-center overflow-hidden">
                  {exercise.gifUrl ? (
                    <img
                      src={`${API_URL}${exercise.gifUrl}`}
                      alt={exercise.namePt || exercise.name}
                      className="w-full h-full object-cover"
                      loading="lazy"
                    />
                  ) : (
                    <Library className="h-12 w-12 text-muted-foreground" />
                  )}
                </div>
                <CardContent className="p-3">
                  <h3 className="font-medium text-sm truncate">
                    {exercise.namePt || exercise.name}
                  </h3>
                  <div className="flex flex-wrap gap-1 mt-2">
                    {exercise.targetMusclesPt?.slice(0, 2).map((m) => (
                      <Badge key={m} variant="secondary" className="text-xs">
                        {m}
                      </Badge>
                    ))}
                    {exercise.equipmentsPt?.slice(0, 1).map((e) => (
                      <Badge key={e} variant="outline" className="text-xs">
                        {e}
                      </Badge>
                    ))}
                  </div>
                </CardContent>
              </Card>
            </motion.div>
          ))}
        </div>
      )}

      {data && data.total > (filters.limit || 20) && (
        <div className="flex justify-center gap-2">
          <Button
            variant="outline"
            size="sm"
            disabled={(filters.offset || 0) === 0}
            onClick={() => setFilters(f => ({ ...f, offset: Math.max(0, (f.offset || 0) - (f.limit || 20)) }))}
          >
            Anterior
          </Button>
          <Button
            variant="outline"
            size="sm"
            disabled={(filters.offset || 0) + (filters.limit || 20) >= data.total}
            onClick={() => setFilters(f => ({ ...f, offset: (f.offset || 0) + (f.limit || 20) }))}
          >
            Próximo
          </Button>
        </div>
      )}
    </div>
  );
}
