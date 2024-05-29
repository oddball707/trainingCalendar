import React from 'react';
import PropTypes from 'prop-types';
import { at } from 'lodash';
import { useField, useFormikContext} from 'formik';
import { NumberField, SwitchField } from './';

import {
  Box,
  FormControl,
  RadioGroup,
  Radio,
  FormHelperText,
  FormControlLabel,
  Card,
  CardActions,
  CardContent,
  Typography,
  Grid,
} from '@material-ui/core';

function SelectRaceType(props) {
  const { label, data, options, ...rest } = props;
  const [field, meta, helpers] = useField(props);
  const { values: formValues } = useFormikContext();
  const { value: selectedValue } = field;
  const [touched, error] = at(meta, 'touched', 'error');
  const isError = touched && error && true;
  function _renderHelperText() {
    if (isError) {
      return <FormHelperText>{error}</FormHelperText>;
    }
  }

  const handleChange = (event) => {
    helpers.setValue(event.target.value)
  };

  return (
    <FormControl {...rest} error={isError}>
      <RadioGroup
        aria-labelledby="demo-radio-buttons-group-label"
        name="radio-buttons-group"
        value={selectedValue ? selectedValue : ''}
        onChange={handleChange}
      >
      {data.map((item, index) => (
        <Card key={index} sx={{ display: 'flex' }} >
          <Box sx={{ display: 'flex', flexDirection: 'row' }}>
            <CardActions>
              <FormControlLabel value={item.value} control={<Radio />} />
            </CardActions>
            <CardContent sx={{ display: 'flex', flexDirection: 'row' }}>
              <Typography component="div" variant="h5">
                {item.label}
              </Typography>
              <Typography variant="subtitle1" color="secondary" component="div">
                {item.description}
              </Typography>
            </CardContent>
          </Box>
          { formValues["raceType"] == '7' && item.value == '7' ?
            <Grid container spacing={1}>
              <Grid item xs={1} md={1}/>
              <Grid item xs={11} md={11}>
                <Typography variant="h6" gutterBottom>
                  {options.weeklyMileage.label}
                </Typography>
                <NumberField name={options.weeklyMileage.name}/>
                <br/>
              </Grid>
              <Grid item xs={1} md={1}/>
              <Grid item xs={11} md={11}>
                <Typography variant="h6" gutterBottom>
                  {options.restDays.label}
                </Typography>
                <NumberField name={options.restDays.name}/>
                <br/>
              </Grid>
              <Grid item xs={1} md={1}/>
              <Grid item xs={11} md={11}>
                <Typography variant="h6" gutterBottom>
                  {options.backToBacks.label}
                </Typography>
                <SwitchField name={options.backToBacks.name} />
              </Grid>
              <Grid item xs={1} md={1}/>
              <Grid item xs={11} md={11}>
                <Typography variant="h6" gutterBottom>
                  {options.increase.label}
                </Typography>
                <NumberField name={options.increase.name} />
              </Grid>
              <Grid item xs={1} md={1}/>
              <Grid item xs={11} md={11}>
                <Typography variant="h6" gutterBottom>
                  {options.restWeekFreq.label}
                </Typography>
                <NumberField name={options.restWeekFreq.name} />
              </Grid>
              <Grid item xs={1} md={1}/>
              <Grid item xs={11} md={11}>
                <Typography variant="h6" gutterBottom>
                  {options.restWeekLevel.label}
                </Typography>
                <NumberField name={options.restWeekLevel.name} />
              </Grid>
            </Grid>
          : null }
        </Card>
      ))}
      </RadioGroup>
      {_renderHelperText()}
    </FormControl>
  );
}

SelectRaceType.defaultProps = {
  data: []
};

SelectRaceType.propTypes = {
  data: PropTypes.array.isRequired
};

export default SelectRaceType;
