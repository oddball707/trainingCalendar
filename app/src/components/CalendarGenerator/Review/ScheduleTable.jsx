import React from 'react';
import axios from 'axios';
import moment from 'moment';
import { Typography, Grid } from '@material-ui/core';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import useStyles from './styles';

const classes = useStyles;
const baseURL = process.env.REACT_APP_API_URL || ''

class ScheduleTable extends React.Component {

  constructor () {
      super();

      this.state = {
          tableData: [{
              weekStart: '',
              days: [],
          }],
      };
  }

  componentDidMount () {
    const payload = {
      "date": moment(this.props.formValues.raceDate).format("MM/D/YY"),
      "type": this.props.formValues.raceType,
      "options":
      {
        "weeklyMileage": this.props.formValues.weeklyMileage,
        "backToBacks": this.props.formValues.backToBacks,
        "restDays": this.props.formValues.restDays,
        "increase": this.props.formValues.increase,
        "restWeekFreq": this.props.formValues.restWeekFreq,
        "restWeekLevel": this.props.formValues.restWeekLevel
      }
    }
    axios({
      url: `${baseURL}/api/show`,
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: payload,
    }).then(response => {
      this.setState({ tableData: response.data });
      console.log(this.state.tableData)
    });
  }

  render () {
    const { tableData } = this.state;
    return (
      <TableContainer component={Paper}>
      <Table className={classes.table} size="large" aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Week</TableCell>
            <TableCell>Date</TableCell>
            <TableCell align="right">Monday</TableCell>
            <TableCell align="right">Tuesday</TableCell>
            <TableCell align="right">Wednesday</TableCell>
            <TableCell align="right">Thursday</TableCell>
            <TableCell align="right">Friday</TableCell>
            <TableCell align="right">Saturday</TableCell>
            <TableCell align="right">Sunday</TableCell>
            <TableCell align="right">Weekly Distance</TableCell>
            <TableCell align="right">Increase over Previous Week</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {tableData.map((week, weekNumber) => (
            <TableRow key={weekNumber}>
              <TableCell component="th" scope="row">
                {weekNumber+1}
              </TableCell>
              <TableCell component="th" scope="row">
                {moment(new Date(week.weekStart)).format("M/DD")}
              </TableCell>
              {week.days.map(day => (
                <TableCell align="right">{day.description}</TableCell>
              ))}
              <TableCell component="th" scope="row">{week.totalDistance}</TableCell>
              <TableCell component="th" scope="row">{week.wowIncrease}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
    );
  }
}

export default ScheduleTable;
