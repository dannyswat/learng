import api from './api';
import { MediaUploadResponse } from '@/types/api.types';

export const mediaService = {
  async uploadImage(file: File, wordId?: string): Promise<MediaUploadResponse> {
    const formData = new FormData();
    formData.append('file', file);
    if (wordId) {
      formData.append('wordId', wordId);
    }

    const response = await api.post<MediaUploadResponse>(
      '/api/v1/media/upload/image',
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      }
    );
    return response.data;
  },

  async uploadAudio(file: File, wordId?: string): Promise<MediaUploadResponse> {
    const formData = new FormData();
    formData.append('file', file);
    if (wordId) {
      formData.append('wordId', wordId);
    }

    const response = await api.post<MediaUploadResponse>(
      '/api/v1/media/upload/audio',
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      }
    );
    return response.data;
  },
};
