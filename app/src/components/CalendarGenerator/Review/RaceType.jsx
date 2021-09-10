import React from 'react';
import { Typography, Grid } from '@material-ui/core';
import useStyles from './styles';

function RaceType(props) {
  const { formValues } = props;
  const classes = useStyles();
  const { raceType } = formValues;
  return (
    <Grid item xs={12} sm={6}>
      <Typography variant="h6" gutterBottom className={classes.title}>
        Race Type
      </Typography>
      <Typography gutterBottom>{`${raceType}`}</Typography>
    </Grid>
  );
}

export default RaceType;