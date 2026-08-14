'use client';

import { Box, Button, Card, Container, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography } from '@mui/material';
import { Add as AddIcon } from '@mui/icons-material';
import React from 'react';

export default function BuyersPage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ mb: 3, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <Typography variant="h4">Buyers</Typography>
        <Button variant="contained" startIcon={<AddIcon />}>Add Buyer</Button>
      </Box>

      <TableContainer component={Card}>
        <Table>
          <TableHead>
            <TableRow sx={{ backgroundColor: '#f5f5f5' }}>
              <TableCell><strong>Code</strong></TableCell>
              <TableCell><strong>Name</strong></TableCell>
              <TableCell><strong>Phone</strong></TableCell>
              <TableCell><strong>City</strong></TableCell>
              <TableCell><strong>Status</strong></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            <TableRow>
              <TableCell colSpan={5} align="center" sx={{ py: 3 }}>No buyers found. Create your first buyer.</TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </TableContainer>
    </Container>
  );
}
