export default {
  formId: 'trainingCalendarForm',
  formField: {
    raceType: {
      name: 'raceType',
      label: 'Race Type*',
      requiredErrorMsg: 'Race type is required'
    },
    raceDate: {
      name: 'raceDate',
      label: 'Race Date*',
      requiredErrorMsg: 'Race date is required',
      futureErrorMsg: 'Race date must be in the future',
      weekendErrorMsg: 'Race date must be on a Saturday or Sunday'
    },
  }
};