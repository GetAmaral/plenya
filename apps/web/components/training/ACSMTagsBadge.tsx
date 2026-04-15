"use client";

import { Badge } from "@/components/ui/badge";
import type { ACSMTagDTO } from "@/lib/api/physical-assessment-api";

interface ACSMTagsBadgeProps {
  tags?: ACSMTagDTO[];
  fallbackTags?: string[];
}

export function ACSMTagsBadge({ tags, fallbackTags }: ACSMTagsBadgeProps) {
  if (tags && tags.length > 0) {
    return (
      <div className="flex flex-wrap gap-2">
        {tags.map((tag, i) => (
          <Badge
            key={`${tag.label}-${i}`}
            variant="outline"
            style={{ borderColor: tag.color, color: tag.color }}
          >
            {tag.label}
          </Badge>
        ))}
      </div>
    );
  }

  if (fallbackTags && fallbackTags.length > 0) {
    return (
      <div className="flex flex-wrap gap-2">
        {fallbackTags.map((tag) => (
          <Badge key={tag} variant="outline">{tag}</Badge>
        ))}
      </div>
    );
  }

  return null;
}
