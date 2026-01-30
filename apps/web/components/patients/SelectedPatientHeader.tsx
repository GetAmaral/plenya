"use client";

import { User2, Calendar, MapPin, RefreshCw, Loader2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { useSelectedPatient } from "@/lib/use-selected-patient";
import { useRouter } from "next/navigation";
import { differenceInYears } from "date-fns";
import { Skeleton } from "@/components/ui/skeleton";

export function SelectedPatientHeader() {
  const router = useRouter();
  const { selectedPatient, isLoading } = useSelectedPatient();

  if (isLoading) {
    return (
      <Card className="mb-6 border-blue-200 bg-blue-50/50">
        <CardContent className="p-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3">
              <Skeleton className="h-12 w-12 rounded-full" />
              <div>
                <Skeleton className="h-5 w-48 mb-2" />
                <Skeleton className="h-4 w-32" />
              </div>
            </div>
            <Skeleton className="h-10 w-40" />
          </div>
        </CardContent>
      </Card>
    );
  }

  if (!selectedPatient) {
    return null;
  }

  const age = differenceInYears(new Date(), new Date(selectedPatient.birthDate));

  const getGenderLabel = (gender: string) => {
    const labels = {
      male: "M",
      female: "F",
      other: "Outro",
    };
    return labels[gender as keyof typeof labels] || gender;
  };

  return (
    <Card className="mb-6 border-blue-200 bg-gradient-to-r from-blue-50 to-purple-50">
      <CardContent className="p-4">
        <div className="flex items-center justify-between flex-wrap gap-4">
          {/* Patient Info */}
          <div className="flex items-center gap-4">
            <div className="flex h-12 w-12 items-center justify-center rounded-full bg-blue-600 text-white">
              <User2 className="h-6 w-6" />
            </div>

            <div>
              <div className="flex items-center gap-2 mb-1">
                <h2 className="text-lg font-bold">{selectedPatient.name}</h2>
                <Badge variant="outline" className="text-xs">
                  Paciente Ativo
                </Badge>
              </div>

              <div className="flex items-center gap-4 text-sm text-muted-foreground">
                {/* Age and Gender */}
                <div className="flex items-center gap-1">
                  <Calendar className="h-3 w-3" />
                  <span>
                    {age} anos • {getGenderLabel(selectedPatient.gender)}
                  </span>
                </div>

                {/* Location */}
                {(selectedPatient.municipality || selectedPatient.state) && (
                  <div className="flex items-center gap-1">
                    <MapPin className="h-3 w-3" />
                    <span>
                      {selectedPatient.municipality}
                      {selectedPatient.municipality &&
                        selectedPatient.state &&
                        " - "}
                      {selectedPatient.state}
                    </span>
                  </div>
                )}
              </div>
            </div>
          </div>

          {/* Actions */}
          <Button
            variant="outline"
            className="gap-2"
            onClick={() => router.push("/patients/select")}
          >
            <RefreshCw className="h-4 w-4" />
            Trocar Paciente
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}
