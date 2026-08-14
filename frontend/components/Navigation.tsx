'use client';

import React, { useState } from 'react';
import {
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Box,
  Typography,
  Divider,
  useMediaQuery,
  useTheme,
  IconButton,
} from '@mui/material';
import {
  Dashboard,
  People,
  Storefront,
  Inventory,
  LocalShipping,
  Settings,
  Menu,
  Close,
} from '@mui/icons-material';
import { useRouter, usePathname } from 'next/navigation';

const drawerWidth = 240;

const menuItems = [
  { label: 'Dashboard', icon: Dashboard, path: '/' },
  { label: 'Suppliers', icon: People, path: '/suppliers' },
  { label: 'Weavers', icon: People, path: '/weavers' },
  { label: 'Buyers', icon: Storefront, path: '/buyers' },
  { label: 'Raw Silk Purchases', icon: LocalShipping, path: '/raw-silk' },
  { label: 'Colouring', icon: Inventory, path: '/colouring' },
  { label: 'Inventory', icon: Inventory, path: '/inventory' },
  { label: 'Settings', icon: Settings, path: '/settings' },
];

export default function Navigation() {
  const theme = useTheme();
  const router = useRouter();
  const pathname = usePathname();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  const [open, setOpen] = useState(!isMobile);

  const handleNavClick = (path: string) => {
    router.push(path);
    if (isMobile) setOpen(false);
  };

  const drawerContent = (
    <Box sx={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
      <Box sx={{ p: 2, backgroundColor: 'primary.main', color: 'white' }}>
        <Typography variant="h6" sx={{ fontWeight: 600 }}>
          🪡 Weaver
        </Typography>
        <Typography variant="caption">Management System</Typography>
      </Box>
      <Divider />
      <List sx={{ flexGrow: 1 }}>
        {menuItems.map((item) => (
          <ListItem key={item.path} disablePadding>
            <ListItemButton
              onClick={() => handleNavClick(item.path)}
              selected={pathname === item.path}
              sx={{
                '&.Mui-selected': {
                  backgroundColor: 'primary.light',
                  color: 'white',
                  '& .MuiListItemIcon-root': {
                    color: 'white',
                  },
                },
              }}
            >
              <ListItemIcon>
                <item.icon />
              </ListItemIcon>
              <ListItemText primary={item.label} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
      <Divider />
      <Box sx={{ p: 2 }}>
        <Typography variant="caption" sx={{ color: 'text.secondary' }}>
          Version 1.0.0
        </Typography>
      </Box>
    </Box>
  );

  return (
    <>
      {isMobile && (
        <Box
          sx={{
            position: 'fixed',
            top: 0,
            left: 0,
            right: 0,
            backgroundColor: 'primary.main',
            color: 'white',
            p: 2,
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            zIndex: 1000,
          }}
        >
          <Typography variant="h6">Weaver Management</Typography>
          <IconButton
            color="inherit"
            onClick={() => setOpen(!open)}
            sx={{ display: { md: 'none' } }}
          >
            {open ? <Close /> : <Menu />}
          </IconButton>
        </Box>
      )}

      {/* Desktop Drawer */}
      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
            display: { xs: 'none', md: 'block' },
            mt: 0,
          },
        }}
        variant="permanent"
        anchor="left"
      >
        {drawerContent}
      </Drawer>

      {/* Mobile Drawer */}
      <Drawer
        sx={{
          display: { xs: 'block', md: 'none' },
          '& .MuiDrawer-paper': {
            boxSizing: 'border-box',
            width: drawerWidth,
            mt: 7,
          },
        }}
        anchor="left"
        open={open}
        onClose={() => setOpen(false)}
      >
        {drawerContent}
      </Drawer>

      {/* Mobile spacing */}
      {isMobile && <Box sx={{ height: 56 }} />}
    </>
  );
}
