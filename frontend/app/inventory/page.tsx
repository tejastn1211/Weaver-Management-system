'use client';

import { Box, Card, Container, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography } from '@mui/material';
import React from 'react';

export default function InventoryPage() {
  return (
    <Container maxWidth="lg">
      <Typography variant="h4" sx={{ mb: 3 }}>Inventory</Typography>

      <Box sx={{ mb: 3 }}>
        <Typography variant="h6" sx={{ mb: 2 }}>📦 Current Stock</Typography>
        <TableContainer component={Card}>
          <Table>
            <TableHead>
              <TableRow sx={{ backgroundColor: '#f5f5f5' }}>
                <TableCell><strong>Batch Reference</strong></TableCell>
                <TableCell><strong>Product</strong></TableCell>
                <TableCell><strong>Colour</strong></TableCell>
                <TableCell><strong>Quantity</strong></TableCell>
                <TableCell><strong>Location</strong></TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              <TableRow>
                <TableCell colSpan={5} align="center" sx={{ py: 3 }}>No inventory found</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>
      </Box>

      <Box>
        <Typography variant="h6" sx={{ mb: 2 }}>📋 Movement History</Typography>
        <TableContainer component={Card}>
          <Table>
            <TableHead>
              <TableRow sx={{ backgroundColor: '#f5f5f5' }}>
                <TableCell><strong>Date</strong></TableCell>
                <TableCell><strong>Type</strong></TableCell>
                <TableCell><strong>Quantity</strong></TableCell>
                <TableCell><strong>From Location</strong></TableCell>
                <TableCell><strong>To Location</strong></TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              <TableRow>
                <TableCell colSpan={5} align="center" sx={{ py: 3 }}>No movements found</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>
      </Box>
    </Container>
  );
}
