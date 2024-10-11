import { z } from "zod";

export default z
  .object({
    VITE_PLACEKIT_API_KEY: z.string(),
    VITE_BACKEND_URL: z.string().url(),
  })
  .parse(import.meta.env);
