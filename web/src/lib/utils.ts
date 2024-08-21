import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { ZodError } from "zod";
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const translateFormError = (error: unknown) => {
  if (error instanceof ZodError) {
    return error.errors[0].message;
  } else if (error instanceof Error) {
    return error.message;
  } else {
    return "An unknown error occurred";
  }
};
