import * as Yup from 'yup';
import moment from 'moment';
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

export default [
  Yup.object().shape({
    [raceType.name]: Yup.string().required(`${raceType.requiredErrorMsg}`),
  }),
  Yup.object().shape({
    [raceDate.name]: Yup.string()
      .nullable()
      .required(`${raceDate.requiredErrorMsg}`)
      .test('raceDate', raceDate.futureErrorMsg, val => {
        if (val) {
          const startDate = new Date();
          const endDate = new Date(2050, 12, 31);
          if (moment(val, moment.ISO_8601).isValid()) {
            return moment(val).isBetween(startDate, endDate);
          }
          return false;
        }
        return false;
      })
      .test('raceDate', raceDate.weekendErrorMsg, val => {
        if (val) {
          const day = moment(val).day()
          if (day == 0 || day == 6) {
            return true;
          }
        }
        return false;
      }),    
  })
];