import React from 'react';
import { useFormikContext, Formik } from 'formik';
import { Grid, Typography, TextField, Switch } from '@material-ui/core';
import { SelectField } from '../../FormFields';
const raceTypes = [
  {
    value: '1',
    label: 'Half Marathon',
    description: 'A 12 week schedule culminating in a half marathon.You should be able to run 20 miles/week to start'
  },
  {
    value: '2',
    label: 'Marathon',
    description: 'A 20 week schedule culminating in a marathon. Suitable to jump into, but should be comfortable running 20 miles/week to start'

  },
  {
    value: '3',
    label: '50K',
    description: 'A 20 week schedule culminating in a 50k ultra. You should have a solid base coming into this schedule, and be comfortable running 30+ miles/week'

  },
  {
    value: '4',
    label: '50 Mile',
    description: 'A 20 week schedule culminating in a 50 mile ultra. You should have a solid base coming into this schedule, and be comfortable running 20+ miles/week'
  },
  {
    value: '5',
    label: '100k',
    description: 'A 20 week schedule culminating in a 100k ultra. You should have a solid base coming into this schedule, and be comfortable running 30+ miles/week'
  },
  {
    value: '6',
    label: '100 Mile',
    description: 'A 27 week schedule culminating in a 100 mile ultra. You should have a solid base coming into this schedule, and be comfortable running 30+ miles/week'

  },
  {
    value: '7',
    label: 'Dynamic Schedule',
    description: 'Generate a schedule with a duration of your choosing, with the ability to tweak various parameters'
  },
]

export default function ScheduleForm(props) {
  const {
    formField: { raceType, options }
  } = props;
  const { values: formValues } = useFormikContext();

  return (
    <React.Fragment>
      <Typography variant="h5" gutterBottom>
        What type of race are you training for?
      </Typography>
      <SelectField
        name={raceType.name}
        label={raceType.label}
        options={options}
        data={raceTypes}
      />
    </React.Fragment>
  );
}
