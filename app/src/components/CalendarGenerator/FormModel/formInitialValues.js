import formModel from './formModel';
const {
  formField: {
    raceType,
    raceDate,
    weeklyMileage,
    backToBacks,
    restDays
  }
} = formModel;

export default {
  [raceType.name]: '',
  [raceDate.name]: '',
  [weeklyMileage]: 15,
  [backToBacks]: true,
  [restDays]: 2
};
