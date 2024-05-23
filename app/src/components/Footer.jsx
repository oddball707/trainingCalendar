import { Link, Typography } from '@material-ui/core/';
export default function Footer() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      <Link color="inherit" href="https://material-ui.com/">
        </Link>
      {new Date().getFullYear()}
    </Typography>
  );
}
