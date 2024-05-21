import React from 'react';
import PropTypes from 'prop-types';
import { at } from 'lodash';
import { useField } from 'formik';
import {
  Box,
  FormLabel,
  FormControl,
  RadioGroup,
  Radio,
  FormHelperText,
  FormControlLabel,
  Card,
  CardActions,
  CardContent,
  Typography,
} from '@material-ui/core';

function SelectField(props) {
  const { label, data, ...rest } = props;
  const [field, meta] = useField(props);
  const { value: selectedValue } = field;
  const [touched, error] = at(meta, 'touched', 'error');
  const isError = touched && error && true;
  function _renderHelperText() {
    if (isError) {
      return <FormHelperText>{error}</FormHelperText>;
    }
  }

  return (
    <FormControl {...rest} error={isError}>
      <RadioGroup
        aria-labelledby="demo-radio-buttons-group-label"
        defaultValue=''
        name="radio-buttons-group"
        value={selectedValue ? selectedValue : ''}
      >
      {data.map((item, index) => (
        <Card sx={{ display: 'flex' }} >
          <Box sx={{ display: 'flex', flexDirection: 'column' }}>
            <CardActions sx={{ flex: '1 0 auto' }}>
              <FormControlLabel value={item.value} control={<Radio />} />
            </CardActions>
          </Box>
          <CardContent sx={{ display: 'flex', flexDirection: 'column' }}>
            <Typography component="div" variant="h5">
              {item.label}
            </Typography>
            <Typography variant="subtitle1" color="text.secondary" component="div">
              {item.description}
            </Typography>
          </CardContent>
        </Card>
      ))}
      </RadioGroup>
    </FormControl>
  );
}

SelectField.defaultProps = {
  data: []
};

SelectField.propTypes = {
  data: PropTypes.array.isRequired
};

export default SelectField;
