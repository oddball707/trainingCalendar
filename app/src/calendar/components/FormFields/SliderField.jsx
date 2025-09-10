import React from 'react';
import { useField } from 'formik';
import Box from '@mui/material/Box';
import Slider from '@mui/material/Slider';

export default function SmallNumberSliderField(props) {
    const { max = 10, min = 0, step = 1, width = '50%' } = props;
    const [field, , helpers] = useField(props.name);
    const { setValue } = helpers;
    const { value } = field;

    function valuetext(value) {
        return `${value}`;
    }

    const handleChange = (event, newValue) => {
        setValue(newValue);
    };

    return (
        <Box sx={{ width: width, px: 2 }}>
            <Slider
                {...field}
                name={props.name}
                onChange={handleChange}
                value={typeof value === 'number' && !isNaN(value) ? value : 0}
                valueLabelFormat={valuetext}
                getAriaValueText={valuetext}
                valueLabelDisplay="auto"
                step={step}
                min={min}
                max={max}
            />
        </Box>
    );
}
