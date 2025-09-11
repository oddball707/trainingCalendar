import React from 'react';
import axios from 'axios';
import moment from 'moment';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import useStyles from './styles';

const baseURL = import.meta.env.REACT_APP_API_URL || '';

class ScheduleTable extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      tableData: [{
        weekStart: '',
        days: [],
      }],
    };
  }

  componentDidMount() {
    const payload = {
      date: moment(this.props.formValues.raceDate).format('MM/D/YY'),
      type: this.props.formValues.raceType,
      options: {
        weeklyMileage: this.props.formValues.weeklyMileage,
        backToBacks: this.props.formValues.backToBacks,
        restDays: this.props.formValues.restDays,
        increase: this.props.formValues.increase,
        restWeekFreq: this.props.formValues.restWeekFreq,
        restWeekLevel: this.props.formValues.restWeekLevel
      }
    };
    axios({
      url: `${baseURL}/api/show`,
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: payload,
    }).then(response => {
      this.setState({ tableData: response.data });
    });
  }

  render() {
    const { tableData } = this.state;
    const classes = useStyles();
    return (
      <TableContainer component={Paper}>
        <Table size="large" aria-label="simple table">
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
            {(Array.isArray(tableData) ? tableData : []).map((week, weekNumber) => (
              <TableRow key={weekNumber}>
                <TableCell component="th" scope="row">
                  {weekNumber + 1}
                </TableCell>
                <TableCell component="th" scope="row">
                  {moment(new Date(week.weekStart)).format('M/DD')}
                </TableCell>
                {(Array.isArray(week.days) ? week.days : []).map((day, idx) => (
                  <TableCell align="right" key={idx}>{day.description}</TableCell>
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
