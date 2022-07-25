export default {
  formId: 'trainingCalendarForm',
  formField: {
    raceType: {
      name: 'raceType',
      label: 'Race Type*',
      requiredErrorMsg: 'Race type is required'
    },
    weeklyMileage: {
      name: 'weeklyMileage',
      label: 'Current weeklyMileage*'
    },
    backToBacks: {
      name: 'backToBacks',
      label: 'Do you want to include back to back long runs?*'
    },
    restDays: {
      name: 'restDays',
      label: 'How many rest days per week?*'
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