import React from 'react';
import Collapse from '@mui/material/Collapse';
import PropTypes from 'prop-types';
import { at } from 'lodash';
import { useField, useFormikContext} from 'formik';
import NumberField from './NumberField';
import PercentField from './PercentField';
import SwitchField from './SwitchField';
import Box from '@mui/material/Box';
import FormControl from '@mui/material/FormControl';
import RadioGroup from '@mui/material/RadioGroup';
import Radio from '@mui/material/Radio';
import FormHelperText from '@mui/material/FormHelperText';
import FormControlLabel from '@mui/material/FormControlLabel';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';

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
        <Card key={index} sx={{ display: 'flex', flexDirection: 'column', height: '100%' }} >
          <Box sx={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start' }}>
            <CardActions sx={{ pt: 2, pb: 0, pl: 0, pr: 0 }}>
              <FormControlLabel value={item.value} control={<Radio />} />
            </CardActions>
            <CardContent sx={{ display: 'flex', flexDirection: 'column', height: '100%', pt: 2, pb: 0 }}>
              <Typography component="div" variant="h5">
                {item.label}
              </Typography>
              <Typography variant="subtitle1" color="secondary" component="div">
                {item.description}
              </Typography>
              <Collapse in={formValues["raceType"] == '7' && item.value == '7'} timeout="auto" unmountOnExit>
                <Grid container direction="column" spacing={2} >
                  <Grid item>
                    <Typography variant="h6" gutterBottom>
                      {options.weeklyMileage.label}
                    </Typography>
                    <NumberField name={options.weeklyMileage.name}/>
                  </Grid>
                  <Grid item>
                    <Typography variant="h6" gutterBottom>
                      {options.restDays.label}
                    </Typography>
                    <NumberField name={options.restDays.name}/>
                  </Grid>
                  <Grid item>
                    <Typography variant="h6" gutterBottom>
                      {options.backToBacks.label}
                    </Typography>
                    <SwitchField name={options.backToBacks.name} />
                  </Grid>
                  <Grid item>
                    <Typography variant="h6" gutterBottom>
                      {options.increase.label}
                    </Typography>
                    <PercentField name={options.increase.name} />
                  </Grid>
                  <Grid item>
                    <Typography variant="h6" gutterBottom>
                      {options.restWeekFreq.label}
                    </Typography>
                    <NumberField name={options.restWeekFreq.name} />
                  </Grid>
                  <Grid item>
                    <Typography variant="h6" gutterBottom>
                      {options.restWeekLevel.label}
                    </Typography>
                    <PercentField name={options.restWeekLevel.name} />
                  </Grid>
                </Grid>
              </Collapse>
            </CardContent>
          </Box>
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
