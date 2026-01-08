import * as React from "react";
import { cva, type VariantProps } from "class-variance-authority";
import { cn } from "@/lib/utils";

const badgeVariants = cva(
  "inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2",
  {
    variants: {
      variant: {
        default:
          "border-transparent bg-primary text-primary-foreground shadow hover:bg-primary/80",
        secondary:
          "border-transparent bg-secondary text-secondary-foreground hover:bg-secondary/80",
        destructive:
          "border-transparent bg-destructive text-destructive-foreground shadow hover:bg-destructive/80",
        outline: "text-foreground",
        // Medical-specific variants
        stable:
          "border-transparent bg-medical-stable/10 text-medical-stable border-medical-stable/20",
        observation:
          "border-transparent bg-medical-observation/10 text-medical-observation border-medical-observation/20",
        critical:
          "border-transparent bg-medical-critical/10 text-medical-critical border-medical-critical/20",
        // Priority variants
        urgent:
          "border-transparent bg-priority-urgent/10 text-priority-urgent border-priority-urgent/20",
        high: "border-transparent bg-priority-high/10 text-priority-high border-priority-high/20",
        normal:
          "border-transparent bg-priority-normal/10 text-priority-normal border-priority-normal/20",
        low: "border-transparent bg-priority-low/10 text-priority-low border-priority-low/20",
      },
    },
    defaultVariants: {
      variant: "default",
    },
  }
);

export interface BadgeProps
  extends React.HTMLAttributes<HTMLDivElement>,
    VariantProps<typeof badgeVariants> {}

function Badge({ className, variant, ...props }: BadgeProps) {
  return (
    <div className={cn(badgeVariants({ variant }), className)} {...props} />
  );
}

export { Badge, badgeVariants };
