import React from 'react';
import { Button, Typography } from '@material-ui/core';

function Success() {
  return (
    <React.Fragment>
      <Typography variant="h5" gutterBottom>
        Your ical training calendar should begin downloading shortly
      </Typography>
      <Typography variant="subtitle1">
        Thank you for using this tool!
      </Typography>
      <Button
        onClick={() => {
          location.reload();
        }}
        variant="contained"
        color="primary"
      >
        Back
      </Button>
    </React.Fragment>
  );
}

export default Success;
