import React from 'react';
import { useFormikContext } from 'formik';
import { Typography, Grid } from '@material-ui/core';
import RaceType from './RaceType';
import RaceDate from './RaceDate';

export default function Review() {
  const { values: formValues } = useFormikContext();
  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        Training Plan Summary
      </Typography>
      <Grid container spacing={2}>
        <RaceType formValues={formValues} />
        <RaceDate formValues={formValues} />
      </Grid>
    </React.Fragment>
  );
}