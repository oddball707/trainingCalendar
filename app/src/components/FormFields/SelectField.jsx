import React from 'react';
import PropTypes from 'prop-types';
import { at } from 'lodash';
import { useField, formik} from 'formik';
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
  const [field, meta, helpers] = useField(props);
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
        </Card>
      ))}
      </RadioGroup>
      {_renderHelperText()}
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
