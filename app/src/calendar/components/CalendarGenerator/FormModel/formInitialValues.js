import formModel from './formModel';
const {
  formField: {
    raceType,
    raceDate,
    options
  }
} = formModel;

export default {
  [raceType.name]: '',
  [raceDate.name]: '',
  [options.weeklyMileage.name]: 15,
  [options.backToBacks.name]: true,
  [options.restDays.name]: 2,
  [options.increase.name]: 10,
  [options.restWeekFreq.name]: 3,
  [options.restWeekLevel.name]: 70
};
