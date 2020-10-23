import React from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import {useForm} from 'react-hook-form'
import Snackbar from '@material-ui/core/Snackbar';
import MuiAlert from '@material-ui/lab/Alert';
import Cookies from 'universal-cookie';

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

export default function SignIn() {
  const classes = useStyles();
  const { register, errors, handleSubmit} = useForm();
	const [open, setOpen] = React.useState(false);

  /* Token Cookie */
  const cookies = new Cookies();

  /* Login Form Handler */
	const onSubmit = (data) => {       

		// Form validation using React Hook Form
		// Basic Integration https://react-hook-form.com/get-started 
		// Integrating in Material UI https://www.youtube.com/watch?v=PquWexbGcVc
		// https://github.com/satansdeer/react-hook-form-material-ui/blob/master/src/App.js

    fetch(process.env.REACT_APP_API_URL + "login",{
      method: 'POST',
      body: JSON.stringify({
          Email: data.email,
          Password: data.password
      })
    })
    .then(res => res.json())
    .then(
      (result) => {
        cookies.set('token', result.Token, { path: '/' });
        cookies.set('email', result.Email, { path: '/' });
        cookies.set('level', result.Level, { path: '/' });
        console.log(result);
        setTimeout(function () {
          window.location.href = "/";
        }, 1);
      }
    )
    .catch(reason => {
      setOpen(true)
    })
		
    console.log(data)
	}    
	
	const handleClose = (event) => {
			setOpen(false);
	};

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form className={classes.form} noValidate onSubmit={handleSubmit(onSubmit)}>
          <TextField
            variant="outlined"
            margin="normal"
            inputRef={register({ required: true })}
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
          />
					<div>{errors.email && "Email is required"}</div>
          <TextField
            variant="outlined"
            margin="normal"
            inputRef={register({ required: true })}
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
          />
					<div>{errors.password && "Password is required"}</div>
          <FormControlLabel
            control={
              <Checkbox inputRef={register} name="remember" color="primary" defaultValue={false}/>}
            label="Remember me"
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            Sign In
          </Button>
          <Grid container>
            <Grid item xs>
              <Link href="#" variant="body2">
                Forgot password?
              </Link>
            </Grid>
            <Grid item>
              <Link href="#" variant="body2">
                {"Don't have an account? Sign Up"}
              </Link>
            </Grid>
          </Grid>
					<Snackbar open={open} autoHideDuration={10000} onClose={handleClose}>
						<Alert severity="error">Access Denied!</Alert>
        	</Snackbar>
        </form>
      </div>
    </Container>
  );
}