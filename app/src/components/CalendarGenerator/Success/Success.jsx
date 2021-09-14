import React from 'react';
import { Typography } from '@material-ui/core';

function Success() {
  return (
    <React.Fragment>
      <Typography variant="h5" gutterBottom>
        Your ical training calendar should begin downloading shortly
      </Typography>
      <Typography variant="subtitle1">
        Thank you using this tool!
      </Typography>
    </React.Fragment>
  );
}

export default Success;