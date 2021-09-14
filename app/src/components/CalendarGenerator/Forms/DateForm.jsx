import React from 'react';
import { Grid, Typography } from '@material-ui/core';
import { DatePickerField } from '../../FormFields';


export default function DateForm(props) {
  const {
    formField: { raceDate }
  } = props;
  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        What is the date of your race?
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} md={6}>
          <DatePickerField
            name={raceDate.name}
            label={raceDate.label}
            format="MM/dd/yyyy"
            minDate={new Date()}
            maxDate={new Date('2050/12/31')}
            fullWidth
          />
        </Grid>
      </Grid>
    </React.Fragment>
  );
}