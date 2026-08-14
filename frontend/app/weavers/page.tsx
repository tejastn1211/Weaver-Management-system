'use client';

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { Box, Button, Card, Container, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography, CircularProgress } from '@mui/material';
import { Add as AddIcon } from '@mui/icons-material';
import { useAuthStore } from '@/hooks/useAuth';
import { weaversAPI } from '@/lib/api';

interface Weaver {
  id: number;
  weaver_code: string;
  name: string;
  phone: string;
  village: string;
  status: string;
}

export default function WeaversPage() {
  const router = useRouter();
  const { isAuthenticated, loadFromLocalStorage } = useAuthStore();
  const [weavers, setWeavers] = useState<Weaver[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadFromLocalStorage();
    if (!isAuthenticated) {
      router.push('/login');
      return;
    }
    fetchWeavers();
  }, [isAuthenticated, router, loadFromLocalStorage]);

  const fetchWeavers = async () => {
    try {
      const response = await weaversAPI.getAll();
      setWeavers(response.data.data || []);
    } catch (error) {
      console.error('Failed to fetch weavers:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return <Box sx={{ display: 'flex', justifyContent: 'center', minHeight: '80vh', alignItems: 'center' }}><CircularProgress /></Box>;
  }

  return (
    <Container maxWidth="lg">
      <Box sx={{ mb: 3, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <Typography variant="h4">Weavers</Typography>
        <Button variant="contained" startIcon={<AddIcon />}>Add Weaver</Button>
      </Box>

      <TableContainer component={Card}>
        <Table>
          <TableHead>
            <TableRow sx={{ backgroundColor: '#f5f5f5' }}>
              <TableCell><strong>Code</strong></TableCell>
              <TableCell><strong>Name</strong></TableCell>
              <TableCell><strong>Phone</strong></TableCell>
              <TableCell><strong>Village</strong></TableCell>
              <TableCell><strong>Status</strong></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {weavers.length === 0 ? (
              <TableRow><TableCell colSpan={5} align="center" sx={{ py: 3 }}>No weavers found</TableCell></TableRow>
            ) : (
              weavers.map((w) => (
                <TableRow key={w.id} hover>
                  <TableCell>{w.weaver_code}</TableCell>
                  <TableCell>{w.name}</TableCell>
                  <TableCell>{w.phone}</TableCell>
                  <TableCell>{w.village}</TableCell>
                  <TableCell><Box sx={{ display: 'inline-block', px: 1.5, py: 0.5, backgroundColor: '#c8e6c9', borderRadius: 1, fontSize: '0.85rem' }}>{w.status}</Box></TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </TableContainer>
    </Container>
  );
}
