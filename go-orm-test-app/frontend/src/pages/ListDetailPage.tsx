import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Trash2 } from 'lucide-react';
import { type FormEvent, useState } from 'react';
import { useParams } from 'react-router-dom';
import Layout from '../components/Layout';
import { Button } from '../components/ui/button.tsx';
import { Checkbox } from '../components/ui/checkbox.tsx';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '../components/ui/select.tsx';
import {
  addItemToList,
  getItems,
  getList,
  removeItemFromList,
  updateListItem,
} from '../lib/api.ts';

export default function ListDetailPage() {
  const { id } = useParams<{ id: string }>();
  const listId = Number.parseInt(id || '0', 10);
  const queryClient = useQueryClient();
  const [selectedItemId, setSelectedItemId] = useState<number | ''>('');

  const {
    data: list,
    isLoading: isLoadingList,
    error: listError,
  } = useQuery({
    queryKey: ['list', listId],
    queryFn: () => getList(listId),
    enabled: !!listId,
  });

  const { data: allItems = [], isLoading: isLoadingItems } = useQuery({
    queryKey: ['items'],
    queryFn: getItems,
  });

  const addItemMutation = useMutation({
    mutationFn: ({ listId, itemId }: { listId: number; itemId: number }) =>
      addItemToList(listId, itemId),
    onSuccess: () => {
      void queryClient.invalidateQueries({ queryKey: ['list', listId] });
      setSelectedItemId('');
    },
  });

  const toggleItemMutation = useMutation({
    mutationFn: ({
      listId,
      itemId,
      bought,
    }: { listId: number; itemId: number; bought: boolean }) =>
      updateListItem(listId, itemId, { bought }),
    onSuccess: () => {
      void queryClient.invalidateQueries({ queryKey: ['list', listId] });
    },
  });

  const removeItemMutation = useMutation({
    mutationFn: ({ listId, itemId }: { listId: number; itemId: number }) =>
      removeItemFromList(listId, itemId),
    onSuccess: () => {
      void queryClient.invalidateQueries({ queryKey: ['list', listId] });
    },
  });

  const handleAddItem = (e: FormEvent) => {
    e.preventDefault();
    if (selectedItemId) {
      addItemMutation.mutate({ listId, itemId: selectedItemId });
    }
  };

  const handleToggleItem = (itemId: number, currentBought: boolean) => {
    toggleItemMutation.mutate({
      listId,
      itemId,
      bought: !currentBought,
    });
  };

  const handleRemoveItem = (itemId: number) => {
    if (
      window.confirm('Are you sure you want to remove this item from the list?')
    ) {
      removeItemMutation.mutate({ listId, itemId });
    }
  };

  const availableItems = allItems.filter(
    (item) => !list?.items.some((listItem) => listItem.id === item.id),
  );

  if (isLoadingList)
    return (
      <Layout>
        <div>Завантаження інформації про список...</div>
      </Layout>
    );

  if (listError)
    return (
      <Layout>
        <div className='text-red-500'>Помилка завантаження списку</div>
      </Layout>
    );

  if (!list)
    return (
      <Layout>
        <div>Список не знайдено</div>
      </Layout>
    );

  return (
    <Layout>
      <div className='space-y-6'>
        <div className='flex justify-between items-center'>
          <h1 className='text-2xl font-bold'>{list.name}</h1>
        </div>

        <div className='bg-white p-4 rounded shadow'>
          <h2 className='text-lg font-semibold mb-4'>
            Додати товари до списку
          </h2>
          <form onSubmit={handleAddItem} className='flex gap-2'>
            <Select
              value={selectedItemId.toString()}
              onValueChange={(value) => {
                setSelectedItemId(+value);
              }}
            >
              <SelectTrigger>
                <SelectValue placeholder='Оберіть товар' />
              </SelectTrigger>
              <SelectContent>
                {(isLoadingItems || !availableItems?.length) && (
                  <SelectItem value='-' disabled>
                    {isLoadingItems
                      ? 'Товарів немає'
                      : 'Завантаження товарів...'}
                  </SelectItem>
                )}
                {availableItems?.map((item) => (
                  <SelectItem key={item.id} value={item.id.toString()}>
                    {item.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            <Button
              type='submit'
              disabled={addItemMutation.isPending || !selectedItemId}
            >
              {addItemMutation.isPending ? 'Додавання...' : 'Додати до списку'}
            </Button>
          </form>
        </div>

        <div className='bg-white p-4 rounded shadow'>
          <h2 className='text-lg font-semibold mb-4'>Товари в списку</h2>

          {list?.items?.length === 0 ? (
            <p>В цьому списку немає товарів. Ви можете додати їх вручну</p>
          ) : (
            <ul className='divide-y divide-gray-200'>
              {list?.items?.map((item) => (
                <li
                  key={item?.id}
                  className='py-4 flex items-center justify-between'
                >
                  <div className='flex items-center'>
                    <Checkbox
                      checked={item?.bought}
                      onCheckedChange={() =>
                        handleToggleItem(item?.id, item?.bought)
                      }
                      className='h-5 w-5 text-blue-600 focus:ring-blue-500 border-gray-300 rounded'
                    />
                    <span
                      className={`ml-3 ${item?.bought ? 'line-through text-gray-500' : ''}`}
                    >
                      {item?.name}
                    </span>
                  </div>
                  <Button
                    onClick={() => handleRemoveItem(item?.id)}
                    disabled={removeItemMutation.isPending}
                    variant='destructive'
                  >
                    <Trash2 />
                  </Button>
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>
    </Layout>
  );
}
