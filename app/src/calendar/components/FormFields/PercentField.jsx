import React from 'react';
import { useField } from 'formik';
import Box from '@mui/material/Box';
import Slider from '@mui/material/Slider';

export default function PercentField(props) {
    const [field, , helpers] = useField(props.name);
    const { setValue } = helpers;
    const { value } = field;

    function valuetext(value) {
        return `${value}%`;
    }

    const handleChange = (event, newValue) => {
        setValue(newValue);
      };

    return (
        <Box sx={{ width: '100%', px: 2 }}>
            <Slider
                {...field}
                name={props.name}
                onChange={handleChange}
                value={typeof value === 'number' && !isNaN(value) ? value : 0}
                valueLabelFormat={valuetext}
                getAriaValueText={valuetext}
                valueLabelDisplay="auto"
                step={5}
                min={0}
                max={100}
            />
        </Box>
    );
}
