import api from './api';
import { Scenario, CreateScenarioRequest } from '@/types/api.types';

export const scenarioService = {
  async getScenarioById(id: string): Promise<Scenario> {
    const response = await api.get<Scenario>(`/api/v1/scenarios/${id}`);
    return response.data;
  },

  async createScenario(data: CreateScenarioRequest): Promise<Scenario> {
    const response = await api.post<Scenario>('/api/v1/scenarios', data);
    return response.data;
  },

  async updateScenario(id: string, data: Partial<CreateScenarioRequest>): Promise<Scenario> {
    const response = await api.put<Scenario>(`/api/v1/scenarios/${id}`, data);
    return response.data;
  },

  async deleteScenario(id: string): Promise<void> {
    await api.delete(`/api/v1/scenarios/${id}`);
  },
};
