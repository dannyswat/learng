import { useState, useCallback, useRef } from 'react';

interface RecorderState {
  isRecording: boolean;
  isPaused: boolean;
  recordingTime: number;
  audioBlob: Blob | null;
  audioUrl: string | null;
  error: string | null;
}

export const useAudioRecorder = () => {
  const [state, setState] = useState<RecorderState>({
    isRecording: false,
    isPaused: false,
    recordingTime: 0,
    audioBlob: null,
    audioUrl: null,
    error: null,
  });

  const mediaRecorderRef = useRef<MediaRecorder | null>(null);
  const chunksRef = useRef<Blob[]>([]);
  const timerRef = useRef<number | null>(null);

  const startRecording = useCallback(async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      const mediaRecorder = new MediaRecorder(stream);
      mediaRecorderRef.current = mediaRecorder;
      chunksRef.current = [];

      mediaRecorder.ondataavailable = (e) => {
        if (e.data.size > 0) {
          chunksRef.current.push(e.data);
        }
      };

      mediaRecorder.onstop = () => {
        const blob = new Blob(chunksRef.current, { type: 'audio/webm' });
        const url = URL.createObjectURL(blob);
        setState((prev) => ({
          ...prev,
          isRecording: false,
          audioBlob: blob,
          audioUrl: url,
        }));

        // Stop all tracks
        stream.getTracks().forEach((track) => track.stop());
        
        // Clear timer
        if (timerRef.current) {
          clearInterval(timerRef.current);
          timerRef.current = null;
        }
      };

      mediaRecorder.start();
      setState((prev) => ({
        ...prev,
        isRecording: true,
        isPaused: false,
        recordingTime: 0,
        audioBlob: null,
        audioUrl: null,
        error: null,
      }));

      // Start timer
      timerRef.current = window.setInterval(() => {
        setState((prev) => ({
          ...prev,
          recordingTime: prev.recordingTime + 1,
        }));
      }, 1000);
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Failed to start recording';
      setState((prev) => ({
        ...prev,
        error: errorMessage,
      }));
    }
  }, []);

  const stopRecording = useCallback(() => {
    if (mediaRecorderRef.current && state.isRecording) {
      mediaRecorderRef.current.stop();
    }
  }, [state.isRecording]);

  const pauseRecording = useCallback(() => {
    if (mediaRecorderRef.current && state.isRecording && !state.isPaused) {
      mediaRecorderRef.current.pause();
      setState((prev) => ({ ...prev, isPaused: true }));
      
      if (timerRef.current) {
        clearInterval(timerRef.current);
        timerRef.current = null;
      }
    }
  }, [state.isRecording, state.isPaused]);

  const resumeRecording = useCallback(() => {
    if (mediaRecorderRef.current && state.isRecording && state.isPaused) {
      mediaRecorderRef.current.resume();
      setState((prev) => ({ ...prev, isPaused: false }));
      
      // Resume timer
      timerRef.current = window.setInterval(() => {
        setState((prev) => ({
          ...prev,
          recordingTime: prev.recordingTime + 1,
        }));
      }, 1000);
    }
  }, [state.isRecording, state.isPaused]);

  const reset = useCallback(() => {
    if (state.audioUrl) {
      URL.revokeObjectURL(state.audioUrl);
    }
    setState({
      isRecording: false,
      isPaused: false,
      recordingTime: 0,
      audioBlob: null,
      audioUrl: null,
      error: null,
    });
  }, [state.audioUrl]);

  return {
    ...state,
    startRecording,
    stopRecording,
    pauseRecording,
    resumeRecording,
    reset,
  };
};
