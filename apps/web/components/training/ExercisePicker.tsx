"use client";

import { useState } from "react";
import { Search, Plus } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Sheet, SheetContent, SheetHeader, SheetTitle, SheetTrigger } from "@/components/ui/sheet";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Badge } from "@/components/ui/badge";
import { useExercises, type Exercise } from "@/lib/api/exercise-api";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001";

interface ExercisePickerProps {
  onSelect: (exercise: Exercise) => void;
  trigger?: React.ReactNode;
}

export function ExercisePicker({ onSelect, trigger }: ExercisePickerProps) {
  const [open, setOpen] = useState(false);
  const [search, setSearch] = useState("");
  const [selectedBodyPart, setSelectedBodyPart] = useState("");

  const { data, isLoading } = useExercises({
    search: search || undefined,
    bodyParts: selectedBodyPart ? [selectedBodyPart] : undefined,
    limit: 30,
  });

  const bodyParts = [
    { value: "", label: "Todos" },
    { value: "back", label: "Costas" },
    { value: "cardio", label: "Cardio" },
    { value: "chest", label: "Peito" },
    { value: "lower arms", label: "Antebraço" },
    { value: "lower legs", label: "Panturrilha" },
    { value: "neck", label: "Pescoço" },
    { value: "shoulders", label: "Ombros" },
    { value: "upper arms", label: "Braços" },
    { value: "upper legs", label: "Pernas" },
    { value: "waist", label: "Abdome" },
  ];

  const handleSelect = (exercise: Exercise) => {
    onSelect(exercise);
    setOpen(false);
  };

  return (
    <Sheet open={open} onOpenChange={setOpen}>
      <SheetTrigger asChild>
        {trigger || (
          <Button type="button" variant="outline" size="sm">
            <Plus className="h-4 w-4 mr-1" /> Adicionar Exercicio
          </Button>
        )}
      </SheetTrigger>
      <SheetContent className="w-full sm:max-w-lg">
        <SheetHeader>
          <SheetTitle>Selecionar Exercicio</SheetTitle>
        </SheetHeader>
        <div className="space-y-3 mt-4">
          <div className="relative">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              placeholder="Buscar exercicio..."
              className="pl-9"
            />
          </div>

          <div className="flex flex-wrap gap-1">
            {bodyParts.map((bp) => (
              <Badge
                key={bp.value}
                variant={selectedBodyPart === bp.value ? "default" : "outline"}
                className="cursor-pointer text-xs"
                onClick={() => setSelectedBodyPart(bp.value)}
              >
                {bp.label}
              </Badge>
            ))}
          </div>

          <ScrollArea className="h-[calc(100vh-220px)]">
            {isLoading ? (
              <p className="text-center text-muted-foreground py-8">Carregando...</p>
            ) : !data?.exercises?.length ? (
              <p className="text-center text-muted-foreground py-8">Nenhum exercicio encontrado.</p>
            ) : (
              <div className="space-y-2 pr-4">
                {data.exercises.map((ex) => (
                  <button
                    key={ex.id}
                    type="button"
                    className="flex items-center gap-3 w-full p-3 rounded-lg border hover:bg-accent transition-colors text-left"
                    onClick={() => handleSelect(ex)}
                  >
                    {ex.gifUrl && (
                      <img
                        src={`${API_URL}${ex.gifUrl}`}
                        alt={ex.namePt || ex.name}
                        className="w-14 h-14 rounded object-cover flex-shrink-0 bg-muted"
                        loading="lazy"
                      />
                    )}
                    <div className="flex-1 min-w-0">
                      <p className="font-medium text-sm truncate">
                        {ex.namePt || ex.name}
                      </p>
                      <div className="flex flex-wrap gap-1 mt-1">
                        {ex.targetMusclesPt?.map((m, i) => (
                          <Badge key={i} variant="secondary" className="text-[10px]">{m}</Badge>
                        )) || ex.targetMuscles?.map((m, i) => (
                          <Badge key={i} variant="secondary" className="text-[10px]">{m}</Badge>
                        ))}
                        {ex.equipmentsPt?.map((e, i) => (
                          <Badge key={`eq-${i}`} variant="outline" className="text-[10px]">{e}</Badge>
                        )) || ex.equipments?.map((e, i) => (
                          <Badge key={`eq-${i}`} variant="outline" className="text-[10px]">{e}</Badge>
                        ))}
                      </div>
                    </div>
                  </button>
                ))}
              </div>
            )}
          </ScrollArea>
        </div>
      </SheetContent>
    </Sheet>
  );
}
