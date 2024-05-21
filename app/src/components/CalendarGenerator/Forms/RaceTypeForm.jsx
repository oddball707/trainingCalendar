import React from 'react';
import { useFormikContext, Formik } from 'formik';
import { Grid, Typography, TextField, Switch } from '@material-ui/core';
import { NumberField, SelectField, SwitchField } from '../../FormFields';
const raceTypes = [
  {
    value: '1',
    label: 'Half Marathon',
    description: 'A 20 week schedule culminating in a half marathon. Suitable to jump into with very little current mileage, but should be able to run 15 miles/week to start'
  },
  {
    value: '2',
    label: 'Marathon',
    description: 'A 20 week schedule culminating in a marathon. Suitable to jump into, but should be comfortable running 20 miles/week to start'

  },
  {
    value: '3',
    label: '50K',
    description: 'A 20 week schedule culminating in a 50k ultra. Should have a solid base coming into this schedule, and be comfortable running 30+ miles/week'

  },
  {
    value: '4',
    label: '50 Mile',
    description: 'A 20 week schedule culminating in a 50 mile ultra. Should have a solid base coming into this schedule, and be comfortable running 20+ miles/week'
  },
  {
    value: '5',
    label: '100k',
    description: 'A 20 week schedule culminating in a 100k ultra. Should have a solid base coming into this schedule, and be comfortable running 30+ miles/week'
  },
  {
    value: '6',
    label: '100 Mile',
    description: 'A 27 week schedule culminating in a 100 mile ultra. Should have a solid base coming into this schedule, and be comfortable running 30+ miles/week'

  },
  {
    value: '7',
    label: 'Dynamic Schedule',
    description: 'Generate a schedule with a duration of your choosing, with the ability to tweak various parameters'
  },
]

export default function ScheduleForm(props) {
  const {
    formField: { raceType, weeklyMileage, backToBacks, restDays }
  } = props;
  const { values: formValues } = useFormikContext();

  return (
    <React.Fragment>
      <Typography variant="h5" gutterBottom>
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
            <Typography variant="h6" gutterBottom>
              What is your current weekly mileage?
            </Typography>
            <br/>
            <NumberField name={weeklyMileage.name}/>
          </Grid>
          <Grid item xs={12} md={12}>
            <Typography variant="h6" gutterBottom>
              How many rest days do you want to schedule per week?
            </Typography>
            <br/>
            <NumberField name={restDays.name}/>
          </Grid>
          <Grid item xs={12} md={12}>
            <Typography variant="h6" gutterBottom>
              Back to back long runs
            </Typography>
            <SwitchField name={backToBacks.name} />
          </Grid>
        </Grid>
      : null }
    </React.Fragment>
  );
}
