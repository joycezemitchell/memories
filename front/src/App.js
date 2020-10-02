import React, { useState, useEffect } from 'react';
import './App.css';

import Maintoolbar from './Components/Maintoolbar/Maintoolbar'
import Maindrawer from './Components/Maindrawer/Maindrawer'
import Videos from './Components/Videos/Videos'
import Videoplay from './Components/Videoplay/Videoplay'
import Login from './Components/Login/Login'

import { makeStyles, useTheme } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import clsx from 'clsx';

import { Route, Switch } from 'react-router-dom';

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
  },
  drawerHeader: {
    display: 'flex',
    alignItems: 'center',
    padding: theme.spacing(0, 1),
    // necessary for content to be below app bar
    ...theme.mixins.toolbar,
    justifyContent: 'flex-end',
  },
  content: {
    flexGrow: 999,
    padding: theme.spacing(3),
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    marginLeft: -drawerWidth,
    textAlign: 'left',
    [theme.breakpoints.down('sm')]: {
      padding: theme.spacing(0),
    },  
  },
  contentShift: {
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
    marginLeft: 0,
  },
}));


let vid_id = ""

function App() {
  const classes = useStyles();
  const theme = useTheme();
  const [open, setOpen] = useState(false);

  const handleChildDrawerOpen = (open) => {
    setOpen(open);
  };

  const handleChildDrawerClose = (id) => {
    vid_id = id 
    setOpen(false);
  };

  return (
    <div className={classes.root}>
      <CssBaseline />
      <Maintoolbar open={open} onOpenDrawer={handleChildDrawerOpen} />
      <Maindrawer open={open} onCloseDrawer={handleChildDrawerClose} />
      <main
        className={clsx(classes.content, {
          [classes.contentShift]: open,
        })}
      >
        <div className={classes.drawerHeader} />
        <Switch>
          <Route path="/Login" component={Login}    />
          <Route path="/Videoplay" component={() => <Videoplay id={vid_id} />}    />
          <Route
            path='/'
            exact
            render={(props) => (
              <Videos {...props} onLoadVideo={e => setOpen(true)} onVideoPlayMain={handleChildDrawerClose} />
            )}
          />
        </Switch>
      </main>
    </div>
  );
}

export default App;
