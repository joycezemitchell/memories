import React, {useState, useEffect} from 'react';
import CardActions from '@material-ui/core/CardActions';
import DeleteIcon from '@material-ui/icons/Delete';
import EditIcon from '@material-ui/icons/Edit';
import Button from '@material-ui/core/Button';
import Alert from './Alert'

export default function Actions(props) {
  const [Actionbutton, setActionbutton] = useState();
  const [open, setOpen] = useState(false);

  const handleDeleteVideo = () => {
    // alert(props.videoId)
    props.onDelete(props.videoId)
    setOpen(false)
  }

  useEffect(() => {
    /* Delete video */
    const handleDelete = (event) => {
      // event.preventDefault();
      // event.stopPropagation();
      
      setOpen(true)
    }

    if (props.userlevel === "1") {
      setActionbutton(
        <CardActions>
          <Button color="primary" size="small" startIcon={<DeleteIcon />}  onClick={(e) => handleDelete(e)}>Delete</Button>
          <Button color="primary" size="small" startIcon={<EditIcon />}>Edit</Button>      
        </CardActions>
      )
    }
  }, [props])

  return (
    <React.Fragment>
      {Actionbutton}
      <Alert open={open} onDelete={handleDeleteVideo}  onClose={e => setOpen(false)} />
    </React.Fragment>
  );

}