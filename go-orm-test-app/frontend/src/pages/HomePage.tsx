import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Trash2 } from 'lucide-react';
import { type FormEvent, useState } from 'react';
import { Link } from 'react-router-dom';
import Layout from '../components/Layout';
import { Button } from '../components/ui/button.tsx';
import { Input } from '../components/ui/input.tsx';
import { createList, deleteList, getLists } from '../lib/api.ts';

export default function HomePage() {
  const [newListName, setNewListName] = useState('');
  const queryClient = useQueryClient();

  const {
    data: lists = [],
    isLoading,
    error,
  } = useQuery({
    queryKey: ['lists'],
    queryFn: getLists,
  });

  const createListMutation = useMutation({
    mutationFn: createList,
    onSuccess: () => {
      void queryClient.invalidateQueries({ queryKey: ['lists'] });
      setNewListName('');
    },
  });

  const deleteListMutation = useMutation({
    mutationFn: deleteList,
    onSuccess: () => {
      void queryClient.invalidateQueries({ queryKey: ['lists'] });
    },
  });

  const handleCreateList = (e: FormEvent) => {
    e.preventDefault();
    if (newListName.trim()) {
      createListMutation.mutate(newListName);
    }
  };

  const handleDeleteList = (id: number) => {
    if (window.confirm('Are you sure you want to delete this list?')) {
      deleteListMutation.mutate(id);
    }
  };

  return (
    <Layout>
      <div className='space-y-6'>
        <div className='flex justify-between items-center'>
          <h1 className='text-2xl font-bold'>Списки</h1>
        </div>

        <div className='bg-white p-4 rounded shadow'>
          <h2 className='text-lg font-semibold mb-4'>Створити новий список</h2>
          <form onSubmit={handleCreateList} className='flex gap-2'>
            <Input
              type='text'
              value={newListName}
              onChange={(e) => setNewListName(e.target.value)}
              placeholder='Назва списку'
              className='flex-1 px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500'
              required
            />
            <Button type='submit' disabled={createListMutation.isPending}>
              {createListMutation.isPending
                ? 'Створення...'
                : 'Створити список'}
            </Button>
          </form>
        </div>

        <div className='bg-white p-4 rounded shadow'>
          <h2 className='text-lg font-semibold mb-4'>Ваші списки покупок</h2>

          {isLoading ? (
            <p>Завантаження списків...</p>
          ) : error ? (
            <p className='text-red-500'>Виникла помилка завантаження списків</p>
          ) : lists.length === 0 ? (
            <p>Не знайдено списків покупок</p>
          ) : (
            <ul className='divide-y divide-gray-200'>
              {lists.map((list) => (
                <li
                  key={list.id}
                  className='py-4 flex justify-between items-center'
                >
                  <Link
                    to={`/lists/${list.id}`}
                    className='text-blue-500 hover:underline'
                  >
                    {list.name}
                  </Link>
                  <Button
                    onClick={() => handleDeleteList(list.id)}
                    disabled={deleteListMutation.isPending}
                    size='icon'
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
