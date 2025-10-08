import api from './api';
import { Word, CreateWordRequest, UpdateWordRequest } from '@/types/api.types';

export const wordService = {
  async getWordById(id: string): Promise<Word> {
    const response = await api.get<Word>(`/api/v1/words/${id}`);
    return response.data;
  },

  async createWord(data: CreateWordRequest): Promise<Word> {
    const response = await api.post<Word>('/api/v1/words', data);
    return response.data;
  },

  async updateWord(id: string, data: UpdateWordRequest): Promise<Word> {
    const response = await api.put<Word>(`/api/v1/words/${id}`, data);
    return response.data;
  },

  async deleteWord(id: string): Promise<void> {
    await api.delete(`/api/v1/words/${id}`);
  },
};
