import api from './api';
import {
  Journey,
  JourneyListResponse,
  CreateJourneyRequest,
  Scenario,
} from '@/types/api.types';

export const journeyService = {
  async getJourneys(params?: {
    status?: string;
    page?: number;
    limit?: number;
  }): Promise<JourneyListResponse> {
    const response = await api.get<JourneyListResponse>('/api/v1/journeys', { params });
    return response.data;
  },

  async getJourneyById(id: string): Promise<Journey> {
    const response = await api.get<Journey>(`/api/v1/journeys/${id}`);
    return response.data;
  },

  async createJourney(data: CreateJourneyRequest): Promise<Journey> {
    const response = await api.post<Journey>('/api/v1/journeys', data);
    return response.data;
  },

  async updateJourney(id: string, data: Partial<CreateJourneyRequest>): Promise<Journey> {
    const response = await api.put<Journey>(`/api/v1/journeys/${id}`, data);
    return response.data;
  },

  async deleteJourney(id: string): Promise<void> {
    await api.delete(`/api/v1/journeys/${id}`);
  },

  async publishJourney(id: string): Promise<Journey> {
    return this.updateJourney(id, { status: 'published' } as Partial<CreateJourneyRequest>);
  },
};
