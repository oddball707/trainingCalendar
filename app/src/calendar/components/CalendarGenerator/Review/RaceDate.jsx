import React from 'react';
import moment from 'moment';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import useStyles from './styles';

function RaceDate(props) {
  const { formValues } = props;
  const classes = useStyles();
  const { raceDate } = formValues;
  return (
  <Grid container direction="column">
      <Typography variant="h6" gutterBottom sx={classes.title}>
        Race Date
      </Typography>
      <Grid container>
        <React.Fragment>
          <Grid>
            <Typography gutterBottom>
              {moment(raceDate).format('MMM Do YYYY')}
            </Typography>
          </Grid>
        </React.Fragment>
      </Grid>
    </Grid>
  );
}

export default RaceDate;
