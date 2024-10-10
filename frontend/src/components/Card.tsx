import { Slot } from "@radix-ui/react-slot";
import { ComponentPropsWithoutRef } from "react";
import { twMerge } from "tailwind-merge";

type CardProps = ComponentPropsWithoutRef<"div"> & {
  variant?: keyof typeof variants;
  asChild?: boolean;
};

const variants = {
  default: "bg-slate-200 p-5 rounded-xl shadow-md",
  dark: "backdrop-blur-lg card-dark rounded-xl p-5 text-white",
};

export function Card({
  children,
  className,
  asChild = false,
  variant = "default",
  ...props
}: CardProps) {
  const Comp = asChild ? Slot : "div";
  return (
    <Comp className={twMerge(variants[variant], className)} {...props}>
      {children}
    </Comp>
  );
}
