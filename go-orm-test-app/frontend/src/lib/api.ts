import axios from 'axios';
import type { ItemSchema } from '../models/item.ts';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export interface List {
  id: number;
  name: string;
}

export interface ListWithItems extends List {
  items: ItemSchema[];
}

export const getLists = async (): Promise<List[]> => {
  const response = await api.get('/lists');
  return response.data;
};

export const getList = async (id: number): Promise<ListWithItems> => {
  const response = await api.get<ListWithItems>(`/lists/${id}`);
  if (!response.data.items) {
    response.data.items = [];
  }
  return response.data;
};

export const createList = async (name: string): Promise<List> => {
  const response = await api.post('/lists', { name });
  return response.data;
};

export const deleteList = async (id: number): Promise<void> => {
  await api.delete(`/lists/${id}`);
};

export const getItems = async (): Promise<ItemSchema[]> => {
  const response = await api.get('/items');
  return response.data;
};

export const createItem = async (name: string): Promise<ItemSchema> => {
  const response = await api.post('/items', { name });
  return response.data;
};

export const updateItem = async (
  id: number,
  data: { name: string; bought: boolean },
): Promise<ItemSchema> => {
  const response = await api.patch(`/items/${id}`, data);
  return response.data;
};

export const deleteItem = async (id: number): Promise<void> => {
  await api.delete(`/items/${id}`);
};

export const addItemToList = async (
  listId: number,
  itemId: number,
): Promise<void> => {
  await api.post(`/lists/${listId}/items`, {
    listId,
    itemId,
  });
};

export const updateListItem = async (
  listId: number,
  itemId: number,
  data: { name?: string; bought?: boolean },
): Promise<void> => {
  await api.patch(`/lists/${listId}/items/${itemId}`, data);
};

export const removeItemFromList = async (
  listId: number,
  itemId: number,
): Promise<void> => {
  await api.delete(`/lists/${listId}/items/${itemId}`);
};

export default api;
