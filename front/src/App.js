import React, { useState, useEffect } from 'react';
import './App.css';

import Maintoolbar from './Components/Maintoolbar/Maintoolbar'
import Maindrawer from './Components/Maindrawer/Maindrawer'
import Videos from './Components/Videos/Videos'
import Videoplay from './Components/Videoplay/Videoplay'
import Login from './Components/Login/Login'
import Blog from './Components/Blog/Blog'
import Upload from './Components/Upload/Upload'
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import clsx from 'clsx';
import { Route, Switch } from 'react-router-dom';
import Cookies from 'universal-cookie';
import InputContext from './Utilities/InputContext/InputContext'
import { useHistory } from "react-router-dom";

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


function App() {
  const classes = useStyles();
  const [open, setOpen] = useState(false);
  const [Routers, setRouters] = useState();
  const [userlevel, setUserlevel] = useState();
  const history = useHistory();


  /* Token Cookie */
  const cookies = new Cookies();

  const handleChildDrawerOpen = (open) => {
    setOpen(open);
  };

  const handleChildDrawerClose = (id) => {
    setOpen(false);
  };

  const handlePlayVideo = (id) => {
    history.push('/videoplay/' + id )
  };

  const gd = {
    User: {
      Level: "",
    }
  }

  useEffect(() => {
    const getSession = async () => {

      try {
        const response = await fetch(process.env.REACT_APP_API_URL + "session", {
          headers: {
            Authorization: `Bearer ` + cookies.get('token'),
          },
        });
        
        const responseData = await response.json();

        console.log(responseData)

        gd.User.Level = cookies.get('level');
        setUserlevel(gd.User.Level)
        setRouters(
          <Route>
            <InputContext.Provider value={gd}>
              <Switch>
                <Route path="/Upload" component={Upload} />
                <Route path="/Videoplay/:idu" component={Videoplay} />
                <Route path='/' exact render={(props) => (<Videos {...props} onLoadVideo={e => setOpen(true)} onVideoPlayMain={handlePlayVideo} />)} />
              </Switch>
            </InputContext.Provider>
          </Route>
        )

      } catch (error) {
        console.log("test")
        console.log(error)
        setRouters(
          <Route>
            <Switch>Videotest
              <Route path="/Blog" component={Blog} />
              <Route path="/" component={Login} />
            </Switch>
          </Route>
        )
      }

    };

    getSession()

  }, [])

  return (
    <div className={classes.root}>
      <CssBaseline />
      <Maintoolbar open={open} onOpenDrawer={handleChildDrawerOpen} />
      <Maindrawer userlevel={userlevel} open={open} onCloseDrawer={handleChildDrawerClose} />
      <main
        className={clsx(classes.content, {
          [classes.contentShift]: open,
        })}
      >
        <div className={classes.drawerHeader} />
        {Routers}
      </main>
    </div>
  );
}

export default App;