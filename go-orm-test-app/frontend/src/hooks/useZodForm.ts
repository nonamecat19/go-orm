import { zodResolver } from '@hookform/resolvers/zod';
import { type UseFormProps, useForm } from 'react-hook-form';
import type z from 'zod';

export type ZodFormOptions<TSchema extends z.Schema> = Omit<
  UseFormProps<z.infer<TSchema>>,
  'resolver'
>;

export default function useZodForm<TSchema extends z.Schema>(
  schema: TSchema,
  options?: ZodFormOptions<TSchema>,
) {
  return useForm({ ...options, resolver: zodResolver(schema) });
}
