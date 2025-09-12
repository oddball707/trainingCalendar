import { createTheme, responsiveFontSizes } from '@mui/material/styles';
import { cyan } from '@mui/material/colors';

let theme = createTheme({
  palette: {
    mode: 'dark',
    primary: cyan,
    secondary: cyan
  }
});

theme = responsiveFontSizes(theme);

const useStyle = () => ({
  root: {
    width: 'auto',
    marginLeft: theme.spacing(2),
    marginRight: theme.spacing(2),
    backgroundColor: theme.palette.background.default,
    color: theme.palette.text.primary
  },
  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(3),
    padding: theme.spacing(2),
    // Responsive styles should be handled with sx prop in MUI v5+
  }
});

export { theme, useStyle };
