import React from 'react';
import { useFormikContext } from 'formik';
import { Grid, Typography, TextField, Switch } from '@material-ui/core';
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
  {
    value: '7',
    label: 'Dynamic Schedule'
  },
]

export default function ScheduleForm(props) {
  const {
    formField: { raceType, weeklyMileage, backToBacks, restDays }
  } = props;
  const { values: formValues } = useFormikContext();

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
      { formValues["raceType"] == 7 ?
        <Grid container spacing={3}>
          <Grid item xs={12} md={12}>
            <Typography variant="h7" gutterBottom>
              What is your current weekly mileage?
            </Typography>
            <br/>
            <TextField 
              defaultValue="10"
              type="number"
              data
              onChange={formik.handleChange}
              value={formik.values.weeklyMileage}
            />
          </Grid>
          <Grid item xs={12} md={12}>
            <Typography variant="h7" gutterBottom>
              How many rest days do you want to schedule per week?
            </Typography>
            <br/>
            <TextField 
              defaultValue="2"
              type="number"
              onChange={formik.handleChange}
              value={formik.values.restDays}
            />
          </Grid>
          <Grid item xs={12} md={12}>
            <Typography variant="h7" gutterBottom>
              Back to back long runs
            </Typography>
            <Switch
              onChange={formik.handleChange}
              value={formik.values.backToBacks}
            />
          </Grid>
        </Grid>
      : null }
    </React.Fragment>
  );
}