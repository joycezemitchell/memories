import React, { useState, useEffect } from 'react';

import clsx from 'clsx';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import CssBaseline from '@material-ui/core/CssBaseline';
import List from '@material-ui/core/List';
import Divider from '@material-ui/core/Divider';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Apps from '@material-ui/icons/Apps';
import Favorite from '@material-ui/icons/Favorite';
import FiberNew from '@material-ui/icons/FiberNew';
import CloudUpload from '@material-ui/icons/CloudUpload';
import ExitToApp from '@material-ui/icons/ExitToApp';
import Cookies from 'universal-cookie';
import { Link } from 'react-router-dom';

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  drawerPaper: {
    width: drawerWidth,
  },
  drawerHeader: {
    display: 'flex',
    alignItems: 'center',
    padding: theme.spacing(0, 1),
    // necessary for content to be below app bar
    ...theme.mixins.toolbar,
    justifyContent: 'flex-end',
  },
}));


export default function Maindrawer(props) {

  /* Token Cookie */
  const cookies = new Cookies();

  const classes = useStyles();
  const theme = useTheme();
  const [open, setOpen] = useState(true);

  const handleDrawerClose = (open, event) => {
    props.onCloseDrawer(false);
  };

  const handleLogout = (event) => {
    cookies.remove('token', { path: '/' });
    cookies.remove('email', { path: '/' });
    cookies.remove('level', { path: '/' });

    setTimeout(function () {
      window.location.href = "/admin";
    }, 500);

  }

  return (
    <Drawer
      className={classes.drawer}
      variant="persistent"
      anchor="left"
      open={props.open}
      classes={{
        paper: classes.drawerPaper,
      }}
    >
      <div className={classes.drawerHeader}>
        <IconButton onClick={e => handleDrawerClose(open, e)}>
          {theme.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
        </IconButton>
      </div>
      <Divider />
      <div onClick={e => handleDrawerClose(open, e)} onKeyDown={e => handleDrawerClose(open, e)}>
        <List>
          <ListItem button component={Link} to="/" >
            <ListItemIcon><Apps /></ListItemIcon>
            <ListItemText primary="Home" />
          </ListItem>
          <ListItem button component={Link} to="/">
            <ListItemIcon><Favorite /></ListItemIcon>
            <ListItemText primary="Favorite" />
          </ListItem>
          <ListItem button component={Link} to="/">
            <ListItemIcon><FiberNew /></ListItemIcon>
            <ListItemText primary="Recently Added" />
          </ListItem>
        </List>
        <Divider />
        {props.userlevel == 1 ?
          <List>
            <ListItem button component={Link} to="/upload" >
              <ListItemIcon><CloudUpload /></ListItemIcon>
              <ListItemText primary="Upload" />
            </ListItem>
            <ListItem button onClick={e => handleLogout(e)}>
              <ListItemIcon><ExitToApp /></ListItemIcon>
              <ListItemText primary="Logout" />
            </ListItem>
          </List>
          : ""
        }
      </div>
    </Drawer>
  );
}