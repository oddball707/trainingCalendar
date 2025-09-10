import React from 'react';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import DatePickerField from '../../FormFields/DatePickerField';

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
        <Grid>
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
