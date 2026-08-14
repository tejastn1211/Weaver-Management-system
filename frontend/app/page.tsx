'use client';

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import {
  Box,
  Grid,
  Card,
  CardContent,
  Typography,
  CircularProgress,
  Container,
} from '@mui/material';
import {
  Storefront,
  People,
  LocalShipping,
  Inventory,
  TrendingUp,
} from '@mui/icons-material';
import { useAuthStore } from '@/hooks/useAuth';
import { dashboardAPI } from '@/lib/api';

interface Stats {
  total_suppliers: number;
  total_weavers: number;
  total_buyers: number;
  inventory_value: number;
  pending_purchases: number;
  processing_batches: number;
}

const StatCard = ({
  title,
  value,
  icon: Icon,
  color,
}: {
  title: string;
  value: number | string;
  icon: any;
  color: string;
}) => (
  <Card sx={{ height: '100%' }}>
    <CardContent>
      <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
        <Box>
          <Typography color="textSecondary" gutterBottom>
            {title}
          </Typography>
          <Typography variant="h5">{value}</Typography>
        </Box>
        <Box
          sx={{
            backgroundColor: `${color}20`,
            borderRadius: '50%',
            p: 2,
            display: 'flex',
          }}
        >
          <Icon sx={{ color, fontSize: 32 }} />
        </Box>
      </Box>
    </CardContent>
  </Card>
);

export default function Dashboard() {
  const router = useRouter();
  const { isAuthenticated, loadFromLocalStorage } = useAuthStore();
  const [stats, setStats] = useState<Stats | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadFromLocalStorage();
    if (!isAuthenticated) {
      router.push('/login');
      return;
    }

    fetchStats();
  }, [isAuthenticated, router, loadFromLocalStorage]);

  const fetchStats = async () => {
    try {
      const response = await dashboardAPI.getStats();
      setStats(response.data);
    } catch (error) {
      console.error('Failed to fetch stats:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '80vh' }}>
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Container maxWidth="lg">
      <Box sx={{ mb: 4 }}>
        <Typography variant="h4" sx={{ mb: 1 }}>
          Dashboard
        </Typography>
        <Typography variant="body2" sx={{ color: 'text.secondary' }}>
          Welcome back! Here's an overview of your business metrics.
        </Typography>
      </Box>

      <Grid container spacing={3}>
        <Grid item xs={12} sm={6} md={4}>
          <StatCard
            title="Total Suppliers"
            value={stats?.total_suppliers || 0}
            icon={Storefront}
            color="#1976d2"
          />
        </Grid>
        <Grid item xs={12} sm={6} md={4}>
          <StatCard
            title="Total Weavers"
            value={stats?.total_weavers || 0}
            icon={People}
            color="#2e7d32"
          />
        </Grid>
        <Grid item xs={12} sm={6} md={4}>
          <StatCard
            title="Total Buyers"
            value={stats?.total_buyers || 0}
            icon={Storefront}
            color="#f57c00"
          />
        </Grid>
        <Grid item xs={12} sm={6} md={4}>
          <StatCard
            title="Inventory Value"
            value={`₹${stats?.inventory_value || 0}`}
            icon={Inventory}
            color="#c62828"
          />
        </Grid>
        <Grid item xs={12} sm={6} md={4}>
          <StatCard
            title="Pending Purchases"
            value={stats?.pending_purchases || 0}
            icon={LocalShipping}
            color="#6a1b9a"
          />
        </Grid>
        <Grid item xs={12} sm={6} md={4}>
          <StatCard
            title="Processing Batches"
            value={stats?.processing_batches || 0}
            icon={TrendingUp}
            color="#00838f"
          />
        </Grid>
      </Grid>

      <Grid container spacing={3} sx={{ mt: 2 }}>
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <Typography variant="h6" sx={{ mb: 2 }}>
                📊 Quick Links
              </Typography>
              <Box sx={{ display: 'flex', gap: 2, flexWrap: 'wrap' }}>
                <Typography
                  variant="body2"
                  sx={{
                    p: 1.5,
                    backgroundColor: '#e3f2fd',
                    borderRadius: 1,
                    cursor: 'pointer',
                    '&:hover': { backgroundColor: '#bbdefb' },
                  }}
                  onClick={() => router.push('/suppliers')}
                >
                  👥 Manage Suppliers
                </Typography>
                <Typography
                  variant="body2"
                  sx={{
                    p: 1.5,
                    backgroundColor: '#f3e5f5',
                    borderRadius: 1,
                    cursor: 'pointer',
                    '&:hover': { backgroundColor: '#e1bee7' },
                  }}
                  onClick={() => router.push('/raw-silk')}
                >
                  📦 Raw Silk Purchases
                </Typography>
                <Typography
                  variant="body2"
                  sx={{
                    p: 1.5,
                    backgroundColor: '#fff3e0',
                    borderRadius: 1,
                    cursor: 'pointer',
                    '&:hover': { backgroundColor: '#ffe0b2' },
                  }}
                  onClick={() => router.push('/colouring')}
                >
                  🎨 Colouring Batches
                </Typography>
                <Typography
                  variant="body2"
                  sx={{
                    p: 1.5,
                    backgroundColor: '#e0f2f1',
                    borderRadius: 1,
                    cursor: 'pointer',
                    '&:hover': { backgroundColor: '#b2dfdb' },
                  }}
                  onClick={() => router.push('/inventory')}
                >
                  📚 Inventory
                </Typography>
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Container>
  );
}
