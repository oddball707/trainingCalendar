import React from 'react';
import { createRoot } from 'react-dom/client';
import CalendarApp from './calendar/App.jsx';
import theme from './calendar/theme.js';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider } from '@mui/material/styles';

createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <CalendarApp />
    </ThemeProvider>
  </React.StrictMode>
);
