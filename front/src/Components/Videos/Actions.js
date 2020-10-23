import React, {useState, useEffect} from 'react';
import clsx from 'clsx';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import CardActions from '@material-ui/core/CardActions';
import DeleteIcon from '@material-ui/icons/Delete';
import EditIcon from '@material-ui/icons/Edit';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles((theme) => ({
  root:{
  },
}));

export default function Actions(props) {

  const classes = useStyles();
  const [expanded, setExpanded] = React.useState(false);
  const theme = useTheme();
  const [Actionbutton, setActionbutton] = useState();

  useEffect(() => {
    if (props.userlevel == 1) {
      setActionbutton(
        <CardActions>
          <Button color="primary" size="small" startIcon={<DeleteIcon />}>Delete</Button>
          <Button color="primary" size="small" startIcon={<EditIcon />}>Edit</Button>      
        </CardActions>
      )
    }
  }, [])

  return (
    <React.Fragment>
      {Actionbutton}
    </React.Fragment>
  );

}