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
    options: {
      weeklyMileage: {
        name: 'weeklyMileage',
        label: 'What is your current weekly mileage?'
      },
      backToBacks: {
        name: 'backToBacks',
        label: 'Do you want to include back to back long runs?'
      },
      restDays: {
        name: 'restDays',
        label: 'How many rest days do you want to schedule per week?',
        minWarning: 'Warning, it is suggested to have at least 1 rest day per week',
        maxWarning: 'Warning, too many rest days'
      },
      increase: {
        name: 'increase',
        label: "What is your target week over week percentage increase in mileage?"
      },
      restWeekFreq: {
        name: 'restWeekFreq',
        label: 'How often do you want to schedule a rest week?'
      },
      restWeekLevel: {
        name: 'restWeekLevel',
        label: 'At what percent of a standard week do you want to plan each rest week?'
      },
      goalTime: {
        name: 'goalTime',
        label: 'What is your goal finish time? (mm:ss)'
      }
    }
  }
};
