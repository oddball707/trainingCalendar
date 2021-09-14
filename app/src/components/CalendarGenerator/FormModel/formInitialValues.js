import formModel from './formModel';
const {
  formField: {
    raceType,
    raceDate
  }
} = formModel;

export default {
  [raceType.name]: '',
  [raceDate.name]: '',
};
