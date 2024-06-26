import React, {useEffect, useState} from 'react';
import { useField } from 'formik';
import { Grid, TextField } from '@material-ui/core';

export default function NumberField(props) {
    const [field, meta, helper] = useField(props);
    const { value } = field;
    const { touched, error } = meta;
    const isError = touched && error && true;
    const [selection, setSelection] = useState(0);

    useEffect(() => {
        if (value) {
            setSelection(value);
        }
    }, [value]);

    return (
        <Grid container>
            <TextField
                {...field}
                {...props}
                type="number"
                error={isError}
                value={selection}
            />
        </Grid>
    );
}
