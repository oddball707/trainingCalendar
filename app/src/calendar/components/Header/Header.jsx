import React from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import useStyles from './styles';
import Box from '@mui/material/Box';

export default function Header() {
  const classes = useStyles();

  return (
    <AppBar position="absolute" sx={{ ...classes.appBar, backgroundColor: 'black' }}>
      <Toolbar>
        <Box component="img"
          src="/logo.png"
          alt="Training Calendars Logo"
          sx={{ height: 100, width: 'auto', display: 'block' }}
        />
      </Toolbar>
    </AppBar>
  );
}
