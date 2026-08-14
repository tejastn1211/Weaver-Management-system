import axios, { AxiosInstance, AxiosError } from 'axios';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api';

const apiClient: AxiosInstance = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
  timeout: 30000,
});

// Request interceptor
apiClient.interceptors.request.use(
  (config) => {
    const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null;
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor
apiClient.interceptors.response.use(
  (response) => response,
  (error: AxiosError) => {
    if (error.response?.status === 401) {
      if (typeof window !== 'undefined') {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        window.location.href = '/login';
      }
    }
    return Promise.reject(error);
  }
);

// API functions
export const authAPI = {
  login: (username: string, password: string) =>
    apiClient.post('/auth/login', { username, password }),
  logout: () => apiClient.post('/auth/logout'),
  getProfile: () => apiClient.get('/auth/profile'),
};

export const suppliersAPI = {
  getAll: () => apiClient.get('/suppliers'),
  getById: (id: number) => apiClient.get(`/suppliers/${id}`),
  create: (data: any) => apiClient.post('/suppliers', data),
  update: (id: number, data: any) => apiClient.put(`/suppliers/${id}`, data),
  delete: (id: number) => apiClient.delete(`/suppliers/${id}`),
};

export const weaversAPI = {
  getAll: () => apiClient.get('/weavers'),
  create: (data: any) => apiClient.post('/weavers', data),
};

export const buyersAPI = {
  getAll: () => apiClient.get('/buyers'),
  create: (data: any) => apiClient.post('/buyers', data),
};

export const rawSilkAPI = {
  getPurchases: () => apiClient.get('/raw-silk-purchases'),
  createPurchase: (data: any) => apiClient.post('/raw-silk-purchases', data),
};

export const colouringAPI = {
  getBatches: () => apiClient.get('/colouring'),
  createBatch: (data: any) => apiClient.post('/colouring', data),
};

export const inventoryAPI = {
  getStock: () => apiClient.get('/inventory/stock'),
  getMovements: () => apiClient.get('/inventory/movements'),
};

export const dashboardAPI = {
  getStats: () => apiClient.get('/dashboard/stats'),
};

export default apiClient;
