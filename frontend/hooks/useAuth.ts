import { create } from 'zustand';

interface User {
  id: number;
  username: string;
  email: string;
  full_name: string;
  role: string;
  status: string;
}

interface AuthStore {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  login: (user: User, token: string) => void;
  logout: () => void;
  setUser: (user: User) => void;
  loadFromLocalStorage: () => void;
}

export const useAuthStore = create<AuthStore>((set) => ({
  user: null,
  token: null,
  isAuthenticated: false,

  login: (user: User, token: string) => {
    set({ user, token, isAuthenticated: true });
    if (typeof window !== 'undefined') {
      localStorage.setItem('user', JSON.stringify(user));
      localStorage.setItem('token', token);
    }
  },

  logout: () => {
    set({ user: null, token: null, isAuthenticated: false });
    if (typeof window !== 'undefined') {
      localStorage.removeItem('user');
      localStorage.removeItem('token');
    }
  },

  setUser: (user: User) => {
    set({ user });
  },

  loadFromLocalStorage: () => {
    if (typeof window !== 'undefined') {
      const user = localStorage.getItem('user');
      const token = localStorage.getItem('token');
      if (user && token) {
        set({
          user: JSON.parse(user),
          token,
          isAuthenticated: true,
        });
      }
    }
  },
}));
