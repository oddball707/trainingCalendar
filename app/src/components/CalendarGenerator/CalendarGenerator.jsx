import React, { useState } from 'react';
import axios from 'axios';
import moment from 'moment';
import {
  Stepper,
  Step,
  StepLabel,
  Button,
  Typography,
  CircularProgress
} from '@material-ui/core';
import { Formik, Form } from 'formik';

import RaceTypeForm from './Forms/RaceTypeForm';
import DateForm from './Forms/DateForm';
import Review from './Review';
import Success from './Success';

import validationSchema from './FormModel/validationSchema';
import formModel from './FormModel/formModel';
import formInitialValues from './FormModel/formInitialValues';

import useStyles from './styles';

const steps = ['Race Type', 'Date', 'Review'];
const { formId, formField } = formModel;
const baseURL = process.env.REACT_APP_API_URL || ''

function _renderStepContent(step) {
  switch (step) {
    case 0:
      return <RaceTypeForm formField={formField} />;
    case 1:
      return <DateForm formField={formField} />;
    case 2:
      return <Review />;
    default:
      return <div>Not Found</div>;
  }
}

export default function CalendarGenerator() {
  const classes = useStyles();
  const [activeStep, setActiveStep] = useState(0);
  const currentValidationSchema = validationSchema[activeStep];
  const isLastStep = activeStep === steps.length - 1;

  function _sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  async function _submitForm(values, actions) {
    const payload = {
      "date": moment(values.raceDate).format("MM/D/YY"),
      "type": values.raceType,
      "options":
      {
        "weeklyMileage": values.weeklyMileage,
        "backToBacks": values.backToBacks,
        "restDays": values.restDays,
        "increase": values.increase,
        "restWeekFreq": values.restWeekFreq,
        "restWeekLevel": values.restWeekLevel
      }
    }
    axios({
      url: `${baseURL}/api/create`,
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: payload,
      responseType: 'blob',
    }).then((response) => {
      const url = window.URL.createObjectURL(new Blob([response.data]));
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', 'training.ics');
      document.body.appendChild(link);
      link.click();
    });

    actions.setSubmitting(false);
    setActiveStep(activeStep + 1);
  }

  function _handleSubmit(values, actions) {
    console.log(values)
    if (isLastStep) {
      _submitForm(values, actions);
    } else {
      setActiveStep(activeStep + 1);
      actions.setTouched({});
      actions.setSubmitting(false);
    }
  }

  function _handleBack() {
    setActiveStep(activeStep - 1);
  }

  return (
    <>
      <Typography component="h1" variant="h4" align="center">
        Generate Training Calendar
      </Typography>
      <Stepper activeStep={activeStep} className={classes.stepper}>
        {steps.map(label => (
          <Step key={label}>
            <StepLabel>{label}</StepLabel>
          </Step>
        ))}
      </Stepper>
      <React.Fragment>
        {activeStep === steps.length ? (
          <Success />
        ) : (
          <Formik
            initialValues={formInitialValues}
            validationSchema={currentValidationSchema}
            onSubmit={_handleSubmit}
          >
            {({ isSubmitting }) => (
              <Form id={formId}>
                {_renderStepContent(activeStep)}

                <div className={classes.buttons}>
                  {activeStep !== 0 && (
                    <Button onClick={_handleBack} className={classes.button}>
                      Back
                    </Button>
                  )}
                  <div className={classes.wrapper}>
                    <Button
                      disabled={isSubmitting}
                      type="submit"
                      variant="contained"
                      color="primary"
                      className={classes.button}
                    >
                      {isLastStep ? 'Generate' : 'Next'}
                    </Button>
                    {isSubmitting && (
                      <CircularProgress
                        size={24}
                        className={classes.buttonProgress}
                      />
                    )}
                  </div>
                </div>
              </Form>
            )}
          </Formik>
        )}
      </React.Fragment>
    </>
  );
}
