export type OperatorType = '=' | '>' | '>=' | '<' | '<=' | 'between';

export interface ScoreLevelRule {
  id: string;
  level: string;
  color: string;
  operatorType: OperatorType;
  minValue?: number;
  maxValue?: number;
  textValue?: string;
}

export interface ScoreItemRule {
  id: string;
  name: string;
  unit?: string;
  gender?: 'male' | 'female' | null;
  ageMin?: number;
  ageMax?: number;
  postMenopause?: boolean;
  levels: ScoreLevelRule[];
}

export interface PatientContext {
  gender?: 'male' | 'female';
  ageYears?: number;
  postMenopause?: boolean;
}

export function isItemEligible(item: ScoreItemRule, patient: PatientContext): boolean {
  if (item.gender && patient.gender && item.gender !== patient.gender) return false;
  if (item.ageMin !== undefined && (patient.ageYears ?? -1) < item.ageMin) return false;
  if (item.ageMax !== undefined && (patient.ageYears ?? Infinity) > item.ageMax) return false;
  if (item.postMenopause === true && patient.postMenopause !== true) return false;
  return true;
}

export function matchLevel(
  levels: ScoreLevelRule[],
  value: number | string,
): ScoreLevelRule | undefined {
  for (const lvl of levels) {
    if (typeof value === 'string') {
      if (lvl.operatorType === '=' && lvl.textValue === value) return lvl;
      continue;
    }

    const n = value;
    switch (lvl.operatorType) {
      case '=':
        if (lvl.minValue === n) return lvl;
        break;
      case '>':
        if (lvl.minValue !== undefined && n > lvl.minValue) return lvl;
        break;
      case '>=':
        if (lvl.minValue !== undefined && n >= lvl.minValue) return lvl;
        break;
      case '<':
        if (lvl.maxValue !== undefined && n < lvl.maxValue) return lvl;
        break;
      case '<=':
        if (lvl.maxValue !== undefined && n <= lvl.maxValue) return lvl;
        break;
      case 'between':
        if (
          lvl.minValue !== undefined &&
          lvl.maxValue !== undefined &&
          n >= lvl.minValue &&
          n <= lvl.maxValue
        ) {
          return lvl;
        }
        break;
    }
  }
  return undefined;
}
