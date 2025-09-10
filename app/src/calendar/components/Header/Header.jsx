import React from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import useStyles from './styles';

export default function Header() {
  const classes = useStyles();

  return (
    <AppBar position="absolute" color="default" sx={classes.appBar}>
      <Toolbar>
        <Typography variant="h6" color="inherit" noWrap>
          Training Calendars
        </Typography>
      </Toolbar>
    </AppBar>
  );
}
