import React from 'react';
import Paper from '@mui/material/Paper';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider } from '@mui/material/styles';
import Header from '../Header/Header';
import Footer from '../Footer';
import { theme, useStyle } from './styles';

export default function MaterialLayout(props) {
  const { children } = props;
  const classes = useStyle();

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Header />
      <div style={classes.root}>
        <Paper sx={{
          mt: { xs: 3, sm: 6 },
          mb: { xs: 3, sm: 6 },
          p: { xs: 2, sm: 3 }
        }}>{children}</Paper>
      </div>
      <Footer />
    </ThemeProvider>
  );
}
