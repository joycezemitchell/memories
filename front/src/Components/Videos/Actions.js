import React, {useState, useEffect} from 'react';
import CardActions from '@material-ui/core/CardActions';
import DeleteIcon from '@material-ui/icons/Delete';
import EditIcon from '@material-ui/icons/Edit';
import Button from '@material-ui/core/Button';

export default function Actions(props) {
  const [Actionbutton, setActionbutton] = useState();

  useEffect(() => {
    if (props.userlevel === "1") {
      setActionbutton(
        <CardActions>
          <Button color="primary" size="small" startIcon={<DeleteIcon />}>Delete</Button>
          <Button color="primary" size="small" startIcon={<EditIcon />}>Edit</Button>      
        </CardActions>
      )
    }
  }, [props.userlevel])

  return (
    <React.Fragment>
      {Actionbutton}
    </React.Fragment>
  );

}