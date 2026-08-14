'use client';

import React from 'react';
import { ThemeProvider, createTheme, CssBaseline, Box } from '@mui/material';
import Navigation from '@/components/Navigation';

const theme = createTheme({
  palette: {
    primary: {
      main: '#1976d2',
      light: '#42a5f5',
      dark: '#1565c0',
    },
    secondary: {
      main: '#dc004e',
    },
    success: {
      main: '#66bb6a',
    },
    warning: {
      main: '#ffa726',
    },
    error: {
      main: '#ef5350',
    },
    background: {
      default: '#f5f5f5',
      paper: '#ffffff',
    },
  },
  typography: {
    fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
    h1: { fontSize: '2.5rem', fontWeight: 600 },
    h2: { fontSize: '2rem', fontWeight: 600 },
    h3: { fontSize: '1.75rem', fontWeight: 600 },
    h4: { fontSize: '1.5rem', fontWeight: 600 },
    h5: { fontSize: '1.25rem', fontWeight: 600 },
  },
  shape: {
    borderRadius: 8,
  },
});

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <Box sx={{ display: 'flex' }}>
            <Navigation />
            <Box
              component="main"
              sx={{
                flexGrow: 1,
                p: 3,
                backgroundColor: '#f5f5f5',
                minHeight: '100vh',
              }}
            >
              {children}
            </Box>
          </Box>
        </ThemeProvider>
      </body>
    </html>
  );
}
