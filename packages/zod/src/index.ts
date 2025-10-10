import { z } from "zod";
import { extendZodWithOpenApi } from "@anatine/zod-openapi";

// Extend Zod with OpenAPI support
extendZodWithOpenApi(z as any);

// Example schema
export const ZHealthResponse = z.object({
  status: z.string(),
  uptime: z.number(),
});

// Export everything needed for downstream packages
export { z };
