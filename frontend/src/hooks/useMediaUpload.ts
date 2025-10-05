import { useState, useCallback } from 'react';
import { mediaService } from '@/services/media.service';

interface UploadState {
  isUploading: boolean;
  progress: number;
  error: string | null;
  url: string | null;
}

export const useMediaUpload = () => {
  const [state, setState] = useState<UploadState>({
    isUploading: false,
    progress: 0,
    error: null,
    url: null,
  });

  const uploadImage = useCallback(async (file: File, wordId?: string) => {
    setState({ isUploading: true, progress: 0, error: null, url: null });

    try {
      const response = await mediaService.uploadImage(file, wordId);
      setState({ isUploading: false, progress: 100, error: null, url: response.url });
      return response.url;
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Upload failed';
      setState({ isUploading: false, progress: 0, error: errorMessage, url: null });
      throw error;
    }
  }, []);

  const uploadAudio = useCallback(async (file: File, wordId?: string) => {
    setState({ isUploading: true, progress: 0, error: null, url: null });

    try {
      const response = await mediaService.uploadAudio(file, wordId);
      setState({ isUploading: false, progress: 100, error: null, url: response.url });
      return response.url;
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Upload failed';
      setState({ isUploading: false, progress: 0, error: errorMessage, url: null });
      throw error;
    }
  }, []);

  const reset = useCallback(() => {
    setState({ isUploading: false, progress: 0, error: null, url: null });
  }, []);

  return {
    ...state,
    uploadImage,
    uploadAudio,
    reset,
  };
};
