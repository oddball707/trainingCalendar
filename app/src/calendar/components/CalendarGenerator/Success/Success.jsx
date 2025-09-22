import React from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

function Success() {
  return (
    <React.Fragment>
      <Box textAlign='center'>
        <Typography variant="h5" gutterBottom>
          Your ical training calendar should begin downloading shortly
        </Typography>
        <Typography variant="subtitle1">
          Thank you for using this tool!
        </Typography>
        <br/>
        <Button
          onClick={() => {
            location.reload();
          }}
          variant="contained"
          color="primary"
        >
          Start Over
        </Button>
      </Box>
    </React.Fragment>
  );
}

export default Success;
