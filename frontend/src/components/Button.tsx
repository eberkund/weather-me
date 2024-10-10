import { Slot } from "@radix-ui/react-slot";
import { twMerge } from "tailwind-merge";

export function Button({ asChild = false, className = "", ...props }) {
  const Comp = asChild ? Slot : "button";
  return (
    <Comp
      className={twMerge(
        "bg-gray-200 hover:bg-gray-400 p-3 rounded-full inline-block",
        className
      )}
      {...props}
    />
  );
}
