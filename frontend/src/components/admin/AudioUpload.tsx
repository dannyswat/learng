import React, { useState, useRef } from 'react';
import { mediaService } from '@/services/media.service';
import { useAudioRecorder } from '@/hooks/useAudioRecorder';
import { Button } from '@/components/shared/Button';

interface AudioUploadProps {
  currentUrl?: string;
  onUploadSuccess: (url: string) => void;
  maxSizeMB?: number;
}

export const AudioUpload: React.FC<AudioUploadProps> = ({
  currentUrl,
  onUploadSuccess,
  maxSizeMB = 2,
}) => {
  const [uploading, setUploading] = useState(false);
  const [audioUrl, setAudioUrl] = useState<string | null>(currentUrl || null);
  const [error, setError] = useState<string | null>(null);
  const [mode, setMode] = useState<'upload' | 'record'>('upload');
  const fileInputRef = useRef<HTMLInputElement>(null);
  const audioRef = useRef<HTMLAudioElement>(null);

  const {
    isRecording,
    recordingTime,
    audioBlob,
    startRecording,
    stopRecording,
    reset: clearRecording,
  } = useAudioRecorder();

  const handleFileSelect = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (!file) return;

    // Validate file type
    if (!file.type.startsWith('audio/')) {
      setError('Please select an audio file (MP3, WAV, or WebM)');
      return;
    }

    // Validate file size
    const maxSize = maxSizeMB * 1024 * 1024;
    if (file.size > maxSize) {
      setError(`File size must be less than ${maxSizeMB}MB`);
      return;
    }

    await uploadFile(file);
  };

  const uploadFile = async (file: File | Blob, filename?: string) => {
    try {
      setUploading(true);
      setError(null);
      
      const fileToUpload = file instanceof Blob && !(file instanceof File)
        ? new File([file], filename || 'recording.webm', { type: file.type })
        : file as File;
      
      const response = await mediaService.uploadAudio(fileToUpload);
      setAudioUrl(response.url);
      onUploadSuccess(response.url);
      clearRecording();
    } catch (err: any) {
      setError(err.response?.data?.error || 'Failed to upload audio');
    } finally {
      setUploading(false);
    }
  };

  const handleRecordingComplete = async () => {
    if (audioBlob) {
      await uploadFile(audioBlob, `recording-${Date.now()}.webm`);
    }
  };

  const handleRemove = () => {
    setAudioUrl(null);
    setError(null);
    onUploadSuccess('');
    clearRecording();
    if (fileInputRef.current) {
      fileInputRef.current.value = '';
    }
  };

  const formatTime = (seconds: number) => {
    const mins = Math.floor(seconds / 60);
    const secs = seconds % 60;
    return `${mins}:${secs.toString().padStart(2, '0')}`;
  };

  return (
    <div className="space-y-3">
      <div className="flex items-center justify-between">
        <label className="block text-sm font-medium text-gray-700">
          Audio
        </label>
        <div className="flex gap-2">
          <button
            type="button"
            onClick={() => setMode('upload')}
            className={`px-3 py-1 text-sm rounded ${
              mode === 'upload'
                ? 'bg-blue-100 text-blue-700'
                : 'bg-gray-100 text-gray-600'
            }`}
          >
            Upload
          </button>
          <button
            type="button"
            onClick={() => setMode('record')}
            className={`px-3 py-1 text-sm rounded ${
              mode === 'record'
                ? 'bg-blue-100 text-blue-700'
                : 'bg-gray-100 text-gray-600'
            }`}
          >
            Record
          </button>
        </div>
      </div>

      {audioUrl ? (
        <div className="border border-gray-200 rounded-lg p-4 space-y-3">
          <audio
            ref={audioRef}
            src={audioUrl.startsWith('/') ? `http://localhost:8080${audioUrl}` : audioUrl}
            controls
            className="w-full"
          />
          <div className="flex gap-2">
            <Button
              type="button"
              variant="secondary"
              size="sm"
              onClick={() => {
                setAudioUrl(null);
                if (mode === 'upload') {
                  fileInputRef.current?.click();
                } else {
                  setMode('record');
                }
              }}
              disabled={uploading}
            >
              Replace
            </Button>
            <Button
              type="button"
              variant="danger"
              size="sm"
              onClick={handleRemove}
              disabled={uploading}
            >
              Remove
            </Button>
          </div>
        </div>
      ) : mode === 'upload' ? (
        <div
          onClick={() => fileInputRef.current?.click()}
          className="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center cursor-pointer hover:border-blue-400 transition-colors"
        >
          <div className="space-y-2">
            <svg
              className="mx-auto h-12 w-12 text-gray-400"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
              />
            </svg>
            <div className="text-sm text-gray-600">
              <span className="font-medium text-blue-600 hover:text-blue-500">
                Click to upload
              </span>
              {' '}audio file
            </div>
            <p className="text-xs text-gray-500">
              MP3, WAV, WebM up to {maxSizeMB}MB
            </p>
          </div>
        </div>
      ) : (
        <div className="border border-gray-200 rounded-lg p-6 space-y-4">
          {audioBlob ? (
            <>
              <audio
                src={URL.createObjectURL(audioBlob)}
                controls
                className="w-full"
              />
              <div className="flex gap-2">
                <Button
                  type="button"
                  variant="primary"
                  onClick={handleRecordingComplete}
                  disabled={uploading}
                  className="flex-1"
                >
                  {uploading ? 'Uploading...' : 'Save Recording'}
                </Button>
                <Button
                  type="button"
                  variant="secondary"
                  onClick={clearRecording}
                  disabled={uploading}
                >
                  Discard
                </Button>
              </div>
            </>
          ) : (
            <div className="text-center space-y-4">
              <div className="flex items-center justify-center">
                <div
                  className={`w-16 h-16 rounded-full flex items-center justify-center ${
                    isRecording ? 'bg-red-100' : 'bg-gray-100'
                  }`}
                >
                  <svg
                    className={`w-8 h-8 ${
                      isRecording ? 'text-red-600' : 'text-gray-600'
                    }`}
                    fill="currentColor"
                    viewBox="0 0 20 20"
                  >
                    <path
                      fillRule="evenodd"
                      d="M7 4a3 3 0 016 0v4a3 3 0 11-6 0V4zm4 10.93A7.001 7.001 0 0017 8a1 1 0 10-2 0A5 5 0 015 8a1 1 0 00-2 0 7.001 7.001 0 006 6.93V17H6a1 1 0 100 2h8a1 1 0 100-2h-3v-2.07z"
                      clipRule="evenodd"
                    />
                  </svg>
                </div>
              </div>
              {isRecording && (
                <div className="text-2xl font-mono text-gray-700">
                  {formatTime(recordingTime)}
                </div>
              )}
              <Button
                type="button"
                variant={isRecording ? 'danger' : 'primary'}
                onClick={isRecording ? stopRecording : startRecording}
                className="w-full"
              >
                {isRecording ? 'Stop Recording' : 'Start Recording'}
              </Button>
            </div>
          )}
        </div>
      )}

      <input
        ref={fileInputRef}
        type="file"
        accept="audio/mpeg,audio/wav,audio/webm,audio/mp3"
        onChange={handleFileSelect}
        className="hidden"
      />

      {uploading && (
        <div className="flex items-center justify-center gap-2 text-sm text-blue-600">
          <svg className="animate-spin h-4 w-4" viewBox="0 0 24 24">
            <circle
              className="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              strokeWidth="4"
              fill="none"
            />
            <path
              className="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
            />
          </svg>
          Uploading...
        </div>
      )}

      {error && (
        <div className="text-sm text-red-600 bg-red-50 border border-red-200 rounded p-2">
          {error}
        </div>
      )}
    </div>
  );
};
