import React from 'react';
import { useField } from 'formik';
import { at } from 'lodash';
import TextField from '@mui/material/TextField';
import FormHelperText from '@mui/material/FormHelperText';

function TimeField(props) {
  const { label = 'Time (mm:ss)', ...rest } = props;
  const [field, meta] = useField(props);
  const { name, value } = field;
  const [touched, error] = at(meta, 'touched', 'error');
  const isError = touched && error && true;

  const handleChange = (event) => {
    let inputValue = event.target.value.replace(/[^0-9:]/g, '');

    // Auto-format as the user types
    if (inputValue.length === 2 && !inputValue.includes(':')) {
      inputValue = inputValue + ':';
    }

    // Limit to mm:ss format (5 characters max)
    if (inputValue.length > 5) {
      inputValue = inputValue.slice(0, 5);
    }

    field.onChange({ target: { name, value: inputValue } });
  };

  return (
    <>
      <TextField
        {...rest}
        label={label}
        name={name}
        value={value || ''}
        onChange={handleChange}
        placeholder="mm:ss"
        error={isError}
        helperText={isError ? error : 'Enter time in mm:ss format'}
      />
    </>
  );
}

export default TimeField;
