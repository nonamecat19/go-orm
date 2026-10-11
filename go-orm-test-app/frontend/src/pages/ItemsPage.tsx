import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Pencil, Trash2 } from 'lucide-react';
import { type FormEvent, useState } from 'react';
import Layout from '../components/Layout';
import { Button } from '../components/ui/button.tsx';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from '../components/ui/form.tsx';
import { Input } from '../components/ui/input.tsx';
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '../components/ui/sheet.tsx';
import useZodForm from '../hooks/useZodForm.ts';
import { createItem, deleteItem, getItems, updateItem } from '../lib/api.ts';
import {
  type ItemSchema,
  type UpdateItemSchema,
  updateItemSchema,
} from '../models/item.ts';

export default function ItemsPage() {
  const [newItemName, setNewItemName] = useState('');
  const [editingItem, setEditingItem] = useState<ItemSchema | null>(null);
  const queryClient = useQueryClient();

  const {
    data: items = [],
    isLoading,
    error,
  } = useQuery({
    queryKey: ['items'],
    queryFn: getItems,
  });

  const createItemMutation = useMutation({
    mutationFn: createItem,
    onSuccess: () => {
      void queryClient.invalidateQueries({ queryKey: ['items'] });
      setNewItemName('');
    },
  });

  const updateItemMutation = useMutation({
    mutationFn: ({
      id,
      data,
    }: { id: number; data: { name: string; bought: boolean } }) =>
      updateItem(id, data),
    onSuccess: () => {
      void queryClient.invalidateQueries({ queryKey: ['items'] });
      setEditingItem(null);
    },
  });

  const deleteItemMutation = useMutation({
    mutationFn: deleteItem,
    onSuccess: () => {
      void queryClient.invalidateQueries({ queryKey: ['items'] });
    },
  });

  const handleCreateItem = (e: FormEvent) => {
    e.preventDefault();
    if (newItemName.trim()) {
      createItemMutation.mutate(newItemName);
    }
  };

  const handleUpdateItem = ({ id, ...data }: UpdateItemSchema) => {
    updateItemMutation.mutate({
      id,
      data,
    });
  };

  const handleDeleteItem = (id: number) => {
    if (window.confirm('Ви впевнені що хочете видалити цей товар?')) {
      deleteItemMutation.mutate(id);
    }
  };

  function EditItem({ item }: { item: ItemSchema }) {
    const [open, setOpen] = useState(false);

    const form = useZodForm(updateItemSchema, {
      defaultValues: item,
    });

    return (
      <Sheet open={open} onOpenChange={setOpen}>
        <SheetTrigger asChild>
          <Button disabled={!!editingItem} size='icon'>
            <Pencil />
          </Button>
        </SheetTrigger>
        <SheetContent className='w-[400px] sm:w-[540px]' side='left'>
          <SheetHeader>
            <SheetTitle>Редагування товару</SheetTitle>
          </SheetHeader>
          <Form {...form}>
            <form>
              <FormField
                control={form.control}
                name='...'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel />
                    <FormControl>
                      <Input {...field} placeholder='Назва товару' />
                    </FormControl>
                  </FormItem>
                )}
              />
              <div className='grid grid-cols-2 gap-2'>
                <Button type='button' disabled={updateItemMutation.isPending}>
                  {updateItemMutation.isPending ? 'Збереження...' : 'Зберегти'}
                </Button>
                <SheetClose asChild>
                  <Button variant='outline'>Відміна</Button>
                </SheetClose>
              </div>
            </form>
          </Form>
        </SheetContent>
      </Sheet>
    );
  }

  return (
    <Layout>
      <div className='space-y-6'>
        <div className='flex justify-between items-center'>
          <h1 className='text-2xl font-bold'>Товари</h1>
        </div>

        <div className='bg-white p-4 rounded shadow'>
          <h2 className='text-lg font-semibold mb-4'>Додати покупку</h2>
          <form onSubmit={handleCreateItem} className='flex gap-2'>
            <Input
              type='text'
              value={newItemName}
              onChange={(e) => setNewItemName(e.target.value)}
              placeholder='Назва товару'
              className='flex-1 px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500'
              required
            />
            <Button type='submit' disabled={createItemMutation.isPending}>
              {createItemMutation.isPending ? 'Створення...' : 'Створити товар'}
            </Button>
          </form>
        </div>

        <div className='bg-white p-4 rounded shadow'>
          <h2 className='text-lg font-semibold mb-4'>Всі товари</h2>

          {isLoading ? (
            <p>Loading items...</p>
          ) : error ? (
            <p className='text-red-500'>Error loading items</p>
          ) : items.length === 0 ? (
            <p>Товарів немає</p>
          ) : (
            <ul className='divide-y divide-gray-200'>
              {items.map((item) => (
                <li
                  key={item.id}
                  className='py-4 flex items-center justify-between'
                >
                  <div className='flex items-center'>
                    <Input
                      type='checkbox'
                      checked={item.bought}
                      onChange={() =>
                        handleUpdateItem({
                          id: item.id,
                          bought: !item.bought,
                          name: item.name,
                        })
                      }
                      className='h-5 w-5 text-blue-600 focus:ring-blue-500 border-gray-300 rounded'
                    />
                    <span
                      className={`ml-3 ${item.bought ? 'line-through text-gray-500' : ''}`}
                    >
                      {item.name}
                    </span>
                  </div>
                  <div className='flex space-x-2'>
                    <EditItem item={item} />

                    <Button
                      onClick={() => handleDeleteItem(item.id)}
                      disabled={deleteItemMutation.isPending}
                      size='icon'
                      variant='destructive'
                    >
                      <Trash2 />
                    </Button>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>
    </Layout>
  );
}
