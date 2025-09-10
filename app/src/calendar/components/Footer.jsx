import Typography from '@mui/material/Typography';
import Link from '@mui/material/Link';

export default function Footer() {
  return (
    <Typography variant="body2" color="text.secondary" align="center">
      <Link color="inherit" href="https://mui.com/">
        MUI
      </Link>{' '}
      {new Date().getFullYear()}
    </Typography>
  );
}
