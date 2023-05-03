import React from 'react';
import { useFormikContext } from 'formik';
import { Typography } from '@material-ui/core';
import ScheduleTable from './ScheduleTable';

export default function Review() {
  const { values: formValues } = useFormikContext();
  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        Training Plan Summary
      </Typography>
      <ScheduleTable formValues={formValues} />
    </React.Fragment>
  );
}
