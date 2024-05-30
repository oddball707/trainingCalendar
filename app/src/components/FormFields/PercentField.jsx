import React from 'react';
import { useField } from 'formik';
import { Grid, Slider } from '@material-ui/core';

export default function PercentField(props) {
    const [field, _, helpers] = useField(props.name);
    const { setValue } = helpers;
    const { value } = field;

    function valuetext(value) {
        return `${value}%`;
    }

    const handleChange = (event, newValue) => {
        setValue(newValue);
      };

    return (
        <Grid container spacing={4}>
            <Grid item xs={4} md={4}>
                <Slider
                    {...field}
                    name={props.name}
                    onChange={handleChange}
                    value={value}
                    valueLabelFormat={valuetext}
                    getAriaValueText={valuetext}
                    valueLabelDisplay="auto"
                    step={5}
                    min={0}
                    max={100}
                />
            </Grid>
        </Grid>
    );
}
