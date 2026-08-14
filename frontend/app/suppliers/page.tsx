'use client';

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import {
  Box,
  Button,
  Card,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
  CircularProgress,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Container,
} from '@mui/material';
import { Add as AddIcon, Edit as EditIcon, Delete as DeleteIcon } from '@mui/icons-material';
import { useAuthStore } from '@/hooks/useAuth';
import { suppliersAPI } from '@/lib/api';

interface Supplier {
  id: number;
  supplier_code: string;
  name: string;
  phone: string;
  email: string;
  city: string;
  material_type: string;
  status: string;
}

export default function SuppliersPage() {
  const router = useRouter();
  const { isAuthenticated, loadFromLocalStorage } = useAuthStore();
  const [suppliers, setSuppliers] = useState<Supplier[]>([]);
  const [loading, setLoading] = useState(true);
  const [openDialog, setOpenDialog] = useState(false);
  const [formData, setFormData] = useState({
    supplier_code: '',
    name: '',
    phone: '',
    email: '',
    city: '',
    material_type: 'Raw Silk',
  });

  useEffect(() => {
    loadFromLocalStorage();
    if (!isAuthenticated) {
      router.push('/login');
      return;
    }
    fetchSuppliers();
  }, [isAuthenticated, router, loadFromLocalStorage]);

  const fetchSuppliers = async () => {
    try {
      const response = await suppliersAPI.getAll();
      setSuppliers(response.data.data || []);
    } catch (error) {
      console.error('Failed to fetch suppliers:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleAddSupplier = async () => {
    if (!formData.name || !formData.phone) {
      alert('Please fill required fields');
      return;
    }

    try {
      await suppliersAPI.create(formData);
      setOpenDialog(false);
      setFormData({ supplier_code: '', name: '', phone: '', email: '', city: '', material_type: 'Raw Silk' });
      fetchSuppliers();
    } catch (error) {
      alert('Failed to add supplier');
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
      <Box sx={{ mb: 3, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <Typography variant="h4">Suppliers</Typography>
        <Button
          variant="contained"
          startIcon={<AddIcon />}
          onClick={() => setOpenDialog(true)}
        >
          Add Supplier
        </Button>
      </Box>

      <TableContainer component={Card}>
        <Table>
          <TableHead>
            <TableRow sx={{ backgroundColor: '#f5f5f5' }}>
              <TableCell><strong>Code</strong></TableCell>
              <TableCell><strong>Name</strong></TableCell>
              <TableCell><strong>Phone</strong></TableCell>
              <TableCell><strong>Email</strong></TableCell>
              <TableCell><strong>City</strong></TableCell>
              <TableCell><strong>Type</strong></TableCell>
              <TableCell><strong>Status</strong></TableCell>
              <TableCell align="right"><strong>Actions</strong></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {suppliers.length === 0 ? (
              <TableRow>
                <TableCell colSpan={8} align="center" sx={{ py: 3 }}>
                  No suppliers found
                </TableCell>
              </TableRow>
            ) : (
              suppliers.map((supplier) => (
                <TableRow key={supplier.id} hover>
                  <TableCell>{supplier.supplier_code}</TableCell>
                  <TableCell>{supplier.name}</TableCell>
                  <TableCell>{supplier.phone}</TableCell>
                  <TableCell>{supplier.email}</TableCell>
                  <TableCell>{supplier.city}</TableCell>
                  <TableCell>{supplier.material_type}</TableCell>
                  <TableCell>
                    <Box
                      sx={{
                        display: 'inline-block',
                        px: 1.5,
                        py: 0.5,
                        backgroundColor: supplier.status === 'Active' ? '#c8e6c9' : '#ffcccc',
                        borderRadius: 1,
                        fontSize: '0.85rem',
                      }}
                    >
                      {supplier.status}
                    </Box>
                  </TableCell>
                  <TableCell align="right">
                    <EditIcon sx={{ cursor: 'pointer', mr: 1, fontSize: '1.2rem' }} />
                    <DeleteIcon sx={{ cursor: 'pointer', fontSize: '1.2rem' }} />
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </TableContainer>

      <Dialog open={openDialog} onClose={() => setOpenDialog(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Add New Supplier</DialogTitle>
        <DialogContent sx={{ display: 'flex', flexDirection: 'column', gap: 2, pt: 2 }}>
          <TextField
            label="Supplier Code"
            value={formData.supplier_code}
            onChange={(e) => setFormData({ ...formData, supplier_code: e.target.value })}
            fullWidth
            size="small"
          />
          <TextField
            label="Name"
            value={formData.name}
            onChange={(e) => setFormData({ ...formData, name: e.target.value })}
            fullWidth
            size="small"
            required
          />
          <TextField
            label="Phone"
            value={formData.phone}
            onChange={(e) => setFormData({ ...formData, phone: e.target.value })}
            fullWidth
            size="small"
            required
          />
          <TextField
            label="Email"
            type="email"
            value={formData.email}
            onChange={(e) => setFormData({ ...formData, email: e.target.value })}
            fullWidth
            size="small"
          />
          <TextField
            label="City"
            value={formData.city}
            onChange={(e) => setFormData({ ...formData, city: e.target.value })}
            fullWidth
            size="small"
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenDialog(false)}>Cancel</Button>
          <Button onClick={handleAddSupplier} variant="contained">
            Add Supplier
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
}
