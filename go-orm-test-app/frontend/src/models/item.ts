import z from 'zod';

export const itemSchema = z.object({
  id: z.number(),
  name: z.string(),
  bought: z.boolean(),
});

export type ItemSchema = z.infer<typeof itemSchema>;

export const updateItemSchema = itemSchema.merge(
  z.object({
    id: z.number(),
  }),
);

export type UpdateItemSchema = z.infer<typeof updateItemSchema>;
