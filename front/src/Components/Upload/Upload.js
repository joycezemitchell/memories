import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';

import Uppy from '@uppy/core'
import Tus from '@uppy/tus'
import { Dashboard } from '@uppy/react'

import '@uppy/core/dist/style.css'
import '@uppy/dashboard/dist/style.css'

const uppy = new Uppy({
	meta: { type: 'avatar' },
	restrictions: { maxNumberOfFiles: 1 },
	autoProceed: true
})

uppy.use(Tus, { endpoint: process.env.REACT_APP_TUSD_ENDPOINT })
// uppy.use(Tus, { endpoint: 'http://localhost:1080/files/' })

uppy.on('complete', (result) => {
	const url = result.successful[0].uploadURL
	React.store.dispatch({
		type: 'SET_USER_AVATAR_URL',
		payload: { url: url }
	})
})

const useStyles = makeStyles((theme) => ({
	root: {

	},
}));

let dOptions = {
	inline: true,
 	target: '.DashboardContainer',
  	replaceTargetContent: true,
  	showProgressDetails: true,
  	width:'100%',
	height: 470,

}


export default function Upload(props) {
	const classes = useStyles();


	return (
		<React.Fragment>
			<Grid container className={classes.root} justify="center">
				<Grid item lg={12} >
					<Dashboard
						uppy={uppy}
						locale={{
							strings: {
								// Text to show on the droppable area.
								// `%{browse}` is replaced with a link that opens the system file selection dialog.
								dropHereOr: 'Drop here or %{browse}',
								// Used as the label for the link that opens the system file selection dialog.
								browse: 'browse'
							}
						}}
						{...dOptions}					
					/>
				</Grid>
			</Grid>
		</React.Fragment>
	);
}    