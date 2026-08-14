'use client';

import { Box, Button, Card, Container, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography } from '@mui/material';
import { Add as AddIcon } from '@mui/icons-material';
import React from 'react';

export default function ColouringPage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ mb: 3, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <Typography variant="h4">Colouring Batches</Typography>
        <Button variant="contained" startIcon={<AddIcon />}>New Batch</Button>
      </Box>

      <TableContainer component={Card}>
        <Table>
          <TableHead>
            <TableRow sx={{ backgroundColor: '#f5f5f5' }}>
              <TableCell><strong>Reference</strong></TableCell>
              <TableCell><strong>Colour</strong></TableCell>
              <TableCell><strong>Quantity</strong></TableCell>
              <TableCell><strong>Date Sent</strong></TableCell>
              <TableCell><strong>Status</strong></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            <TableRow>
              <TableCell colSpan={5} align="center" sx={{ py: 3 }}>No colouring batches found. Create your first batch.</TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </TableContainer>
    </Container>
  );
}
