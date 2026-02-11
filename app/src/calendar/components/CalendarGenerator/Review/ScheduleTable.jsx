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

const baseURL = import.meta.env.VITE_API_URL || '';

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
        restWeekLevel: this.props.formValues.restWeekLevel,
        goalTime: this.props.formValues.goalTime
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
    const raceType = this.props.formValues?.raceType;
    return (
      <>
        <TableContainer component={Paper}>
          <Table size="large" aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell>Week</TableCell>
                <TableCell>Date</TableCell>
                <TableCell align="center">Monday</TableCell>
                <TableCell align="center">Tuesday</TableCell>
                <TableCell align="center">Wednesday</TableCell>
                <TableCell align="center">Thursday</TableCell>
                <TableCell align="center">Friday</TableCell>
                <TableCell align="center">Saturday</TableCell>
                <TableCell align="center">Sunday</TableCell>
                {raceType !== '1' && <TableCell align="center">Weekly Distance</TableCell>}
                {raceType !== '1' && <TableCell align="center">Increase over Previous Week</TableCell>}
              </TableRow>
            </TableHead>
            <TableBody>
              {(Array.isArray(tableData) ? tableData : []).map((week, weekNumber) => (
                <TableRow key={weekNumber} sx={week.wowIncrease === '-' && weekNumber > 1 ? { backgroundColor: '#575757ff' } : {}}>
                  <TableCell component="th" scope="row">
                    {weekNumber + 1}
                  </TableCell>
                  <TableCell component="th" scope="row">
                    {moment(new Date(week.weekStart)).format('M/DD')}
                  </TableCell>
                  {(Array.isArray(week.days) ? week.days : []).map((day, idx) => (
                    <TableCell align="center" key={idx}>{day.title}</TableCell>
                  ))}
                  {raceType !== '1' && (
                    <TableCell align="center" component="th" scope="row">{week.totalDistance}</TableCell>
                  )}
                  {raceType !== '1' && (
                    <TableCell align="center" component="th" scope="row">{week.wowIncrease}</TableCell>
                  )}
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        {raceType !== '1' && (
          <div style={{ marginTop: 16, display: 'flex', alignItems: 'center' }}>
            <div style={{ width: 24, height: 24, backgroundColor: '#575757ff', border: '1px solid #575757ff', marginRight: 8 }} />
            <Typography variant="body2">Rest/Taper week</Typography>
          </div>
        )}
      </>
    );
  }
}

export default ScheduleTable;
