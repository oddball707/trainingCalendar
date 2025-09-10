import React, {useEffect, useState} from 'react';
import { useField } from 'formik';
import Grid from '@mui/material/Grid';
import Switch from '@mui/material/Switch';

export default function SwitchField(props) {
    const [field, meta, helper] = useField(props);
    const { setValue } = helper;
    const { value } = field;
    const [selection, setSelection] = useState(false);

    useEffect(() => {
        if (value) {
            setSelection(value);
        }
    }, [value]);

    function _onChange(val) {
        if (val) {
            setSelection(val);
        }
        setValue(val);
    }

    return (
        <Grid container>
            <Switch
                {...field}
                {...props}
                value={selection}
            />
        </Grid>
    );
}
