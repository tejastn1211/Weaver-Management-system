'use client';

import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import {
  Box,
  Card,
  TextField,
  Button,
  Typography,
  Alert,
  CircularProgress,
  Container,
} from '@mui/material';
import { authAPI } from '@/lib/api';
import { useAuthStore } from '@/hooks/useAuth';

export default function LoginPage() {
  const router = useRouter();
  const { login, isAuthenticated } = useAuthStore();
  const [username, setUsername] = useState('admin');
  const [password, setPassword] = useState('admin123');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (isAuthenticated) {
      router.push('/');
    }
  }, [isAuthenticated, router]);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      const response = await authAPI.login(username, password);
      const { token, user } = response.data.data;
      login(user, token);
      router.push('/');
    } catch (err: any) {
      setError(
        err.response?.data?.message || 'Login failed. Please try again.'
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box
      sx={{
        minHeight: '100vh',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#f5f5f5',
      }}
    >
      <Container maxWidth="sm">
        <Card sx={{ p: 4, boxShadow: 3, borderRadius: 2 }}>
          <Box sx={{ textAlign: 'center', mb: 3 }}>
            <Typography variant="h3" sx={{ fontSize: '2.5rem', mb: 1 }}>
              🪡
            </Typography>
            <Typography variant="h4" sx={{ mb: 1 }}>
              Weaver Management
            </Typography>
            <Typography variant="body2" sx={{ color: 'text.secondary' }}>
              Sign in to your account
            </Typography>
          </Box>

          {error && (
            <Alert severity="error" sx={{ mb: 2 }}>
              {error}
            </Alert>
          )}

          <Box component="form" onSubmit={handleLogin} sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
            <TextField
              label="Username"
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              fullWidth
              disabled={loading}
              placeholder="admin, manager, or accountant"
              helperText="Demo accounts available"
            />
            <TextField
              label="Password"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              fullWidth
              disabled={loading}
              placeholder="admin123, manager123, or accountant123"
            />

            <Box sx={{ mt: 1 }}>
              <Typography variant="caption" sx={{ color: 'text.secondary', display: 'block', mb: 1 }}>
                Demo Credentials:
              </Typography>
              <Typography variant="caption" sx={{ color: 'text.secondary' }}>
                • admin / admin123 (Admin)<br/>
                • manager / manager123 (Manager)<br/>
                • accountant / accountant123 (Accountant)
              </Typography>
            </Box>

            <Button
              type="submit"
              variant="contained"
              size="large"
              fullWidth
              disabled={loading}
              sx={{ mt: 2 }}
            >
              {loading ? <CircularProgress size={24} sx={{ color: 'white' }} /> : 'Sign In'}
            </Button>
          </Box>
        </Card>

        <Box sx={{ mt: 3, textAlign: 'center' }}>
          <Typography variant="body2" sx={{ color: 'text.secondary' }}>
            This is a demo version. Use any of the demo credentials above to login.
          </Typography>
        </Box>
      </Container>
    </Box>
  );
}
