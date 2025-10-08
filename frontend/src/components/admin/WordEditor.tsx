import React, { useState } from 'react';
import { Word, CreateWordRequest } from '@/types/api.types';
import { ImageUpload } from './ImageUpload';
import { AudioUpload } from './AudioUpload';
import { Button } from '@/components/shared/Button';
import { Input } from '@/components/shared/Input';

interface WordEditorProps {
  word?: Word;
  scenarioId: string;
  onSave: (data: CreateWordRequest) => Promise<void>;
  onCancel: () => void;
}

export const WordEditor: React.FC<WordEditorProps> = ({
  word,
  scenarioId,
  onSave,
  onCancel,
}) => {
  const [formData, setFormData] = useState<CreateWordRequest>({
    scenarioId: word?.scenarioId || scenarioId,
    targetText: word?.targetText || '',
    sourceText: word?.sourceText || '',
    displayOrder: word?.displayOrder || 0,
    imageUrl: word?.imageUrl || '',
    audioUrl: word?.audioUrl || '',
  });
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!formData.targetText.trim()) {
      setError('Target text is required');
      return;
    }
    
    if (!formData.sourceText.trim()) {
      setError('Source text is required');
      return;
    }

    try {
      setSaving(true);
      setError(null);
      await onSave(formData);
    } catch (err: any) {
      setError(err.response?.data?.error || 'Failed to save word');
    } finally {
      setSaving(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="space-y-4">
          <Input
            label="Target Language Text"
            value={formData.targetText}
            onChange={(e) =>
              setFormData({ ...formData, targetText: e.target.value })
            }
            placeholder="e.g., 蘋果"
            required
          />

          <Input
            label="Source Language Text"
            value={formData.sourceText}
            onChange={(e) =>
              setFormData({ ...formData, sourceText: e.target.value })
            }
            placeholder="e.g., apple"
            required
          />

          <Input
            label="Display Order"
            type="number"
            value={formData.displayOrder}
            onChange={(e) =>
              setFormData({ ...formData, displayOrder: parseInt(e.target.value) || 0 })
            }
            min={0}
          />
        </div>

        <div className="space-y-4">
          <ImageUpload
            currentUrl={formData.imageUrl || undefined}
            onUploadSuccess={(url) =>
              setFormData({ ...formData, imageUrl: url })
            }
          />

          <AudioUpload
            currentUrl={formData.audioUrl || undefined}
            onUploadSuccess={(url) =>
              setFormData({ ...formData, audioUrl: url })
            }
          />
        </div>
      </div>

      {error && (
        <div className="bg-red-50 border border-red-200 rounded-lg p-3 text-sm text-red-700">
          {error}
        </div>
      )}

      <div className="flex justify-end gap-3">
        <Button
          type="button"
          variant="secondary"
          onClick={onCancel}
          disabled={saving}
        >
          Cancel
        </Button>
        <Button
          type="submit"
          variant="primary"
          disabled={saving}
        >
          {saving ? 'Saving...' : word ? 'Update Word' : 'Create Word'}
        </Button>
      </div>
    </form>
  );
};
