import api from './api';
import {
  AuthResponse,
  LoginRequest,
  RegisterRequest,
} from '@/types/api.types';

export const authService = {
  async login(credentials: LoginRequest): Promise<AuthResponse> {
    const response = await api.post<AuthResponse>('/api/v1/auth/login', credentials);
    const { user, token } = response.data;
    
    // Store token and user in localStorage
    localStorage.setItem('authToken', token);
    localStorage.setItem('user', JSON.stringify(user));
    
    return response.data;
  },

  async register(data: RegisterRequest): Promise<AuthResponse> {
    const response = await api.post<AuthResponse>('/api/v1/auth/register', data);
    const { user, token } = response.data;
    
    // Store token and user in localStorage
    localStorage.setItem('authToken', token);
    localStorage.setItem('user', JSON.stringify(user));
    
    return response.data;
  },

  logout(): void {
    localStorage.removeItem('authToken');
    localStorage.removeItem('user');
  },

  getCurrentUser(): string | null {
    const userStr = localStorage.getItem('user');
    return userStr;
  },

  isAuthenticated(): boolean {
    return !!localStorage.getItem('authToken');
  },
};
