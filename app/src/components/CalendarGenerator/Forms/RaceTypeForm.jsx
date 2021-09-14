import React from 'react';
import { Grid, Typography } from '@material-ui/core';
import { SelectField, DatePickerField } from '../../FormFields';

const raceTypes = [
  {
    value: undefined,
    label: 'None'
  },
  {
    value: '1',
    label: 'Half Marathon'
  },
  {
    value: '2',
    label: 'Marathon'
  },
  {
    value: '3',
    label: '50K'
  },
  {
    value: '4',
    label: '50 Mile'
  },
  {
    value: '5',
    label: '100k'
  },
  {
    value: '6',
    label: '100 Mile'
  },
  
  
]

export default function PaymentForm(props) {
  const {
    formField: { raceType }
  } = props;

  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        What type of race are you training for?
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} md={6}>
          <SelectField
            name={raceType.name}
            label={raceType.label}
            data={raceTypes}
            fullWidth
          />
        </Grid>
      </Grid>
    </React.Fragment>
  );
}