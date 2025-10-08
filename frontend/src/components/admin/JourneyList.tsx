import React, { useState, useEffect } from 'react';
import { journeyService } from '@/services/journey.service';
import { Journey } from '@/types/api.types';
import { Button } from '@/components/shared/Button';

export const JourneyList: React.FC = () => {
  const [journeys, setJourneys] = useState<Journey[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [filter, setFilter] = useState<string>('all');

  useEffect(() => {
    loadJourneys();
  }, [filter]);

  const loadJourneys = async () => {
    try {
      setLoading(true);
      const params = filter === 'all' ? {} : { status: filter };
      const response = await journeyService.getJourneys(params);
      setJourneys(response.journeys);
      setError(null);
    } catch (err: any) {
      setError(err.response?.data?.error || 'Failed to load journeys');
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id: string) => {
    if (!confirm('Are you sure you want to delete this journey? This will also delete all scenarios and words.')) {
      return;
    }

    try {
      await journeyService.deleteJourney(id);
      await loadJourneys();
    } catch (err: any) {
      alert(err.response?.data?.error || 'Failed to delete journey');
    }
  };

  const handlePublish = async (id: string) => {
    try {
      await journeyService.publishJourney(id);
      await loadJourneys();
    } catch (err: any) {
      alert(err.response?.data?.error || 'Failed to publish journey');
    }
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="text-gray-500">Loading journeys...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="bg-red-50 border border-red-200 rounded-lg p-4">
        <p className="text-red-800">{error}</p>
        <Button onClick={loadJourneys} className="mt-2">Retry</Button>
      </div>
    );
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-800">Learning Journeys</h2>
        <div className="flex gap-4">
          <select
            value={filter}
            onChange={(e) => setFilter(e.target.value)}
            className="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="all">All Journeys</option>
            <option value="draft">Drafts</option>
            <option value="published">Published</option>
            <option value="archived">Archived</option>
          </select>
          <Button onClick={() => window.location.href = '/admin/journeys/new'}>
            + Create Journey
          </Button>
        </div>
      </div>

      {journeys.length === 0 ? (
        <div className="text-center py-12 bg-gray-50 rounded-lg">
          <p className="text-gray-500 mb-4">No journeys found</p>
          <Button onClick={() => window.location.href = '/admin/journeys/new'}>
            Create Your First Journey
          </Button>
        </div>
      ) : (
        <div className="grid gap-4">
          {journeys.map((journey) => (
            <div
              key={journey.id}
              className="bg-white border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
            >
              <div className="flex justify-between items-start">
                <div className="flex-1">
                  <div className="flex items-center gap-3 mb-2">
                    <h3 className="text-xl font-semibold text-gray-800">
                      {journey.title}
                    </h3>
                    <span
                      className={`px-2 py-1 text-xs font-medium rounded ${
                        journey.status === 'published'
                          ? 'bg-green-100 text-green-800'
                          : journey.status === 'archived'
                          ? 'bg-gray-100 text-gray-800'
                          : 'bg-yellow-100 text-yellow-800'
                      }`}
                    >
                      {journey.status}
                    </span>
                  </div>
                  <p className="text-gray-600 mb-3">{journey.description}</p>
                  <div className="flex gap-4 text-sm text-gray-500">
                    <span>
                      {journey.sourceLanguage} → {journey.targetLanguage}
                    </span>
                    <span>•</span>
                    <span>{journey.scenarioCount || 0} scenarios</span>
                    <span>•</span>
                    <span>{journey.wordCount || 0} words</span>
                  </div>
                </div>
                <div className="flex gap-2">
                  <Button
                    variant="secondary"
                    size="sm"
                    onClick={() => window.location.href = `/admin/journeys/${journey.id}`}
                  >
                    Edit
                  </Button>
                  {journey.status === 'draft' && (
                    <Button
                      variant="primary"
                      size="sm"
                      onClick={() => handlePublish(journey.id)}
                    >
                      Publish
                    </Button>
                  )}
                  <Button
                    variant="danger"
                    size="sm"
                    onClick={() => handleDelete(journey.id)}
                  >
                    Delete
                  </Button>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};
