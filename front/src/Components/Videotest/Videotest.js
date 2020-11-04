import React, { useState, useEffect } from 'react';
import Video from '../Videos/Video'
import Grid from '@material-ui/core/Grid';
import NdjsonParser from 'ndjson-parse';
import Cookies from 'universal-cookie';

let vsi = [];

/* Token Cookie */
const cookies = new Cookies();

export default function Videotest(props) {

	const handleVideoPlay = (id) => {
		props.onVideoPlayMain(id);
	};

	const [Videos, setVideos] = useState("")

	useEffect(() => {
		fetch(process.env.REACT_APP_API_URL + "/videos")
			.then(res => res.json())
			.then(
				(result) => {
					if (result != null) {
						let thumb = ""
						/*setVideos(vsi.map((v) => {
							thumb = process.env.REACT_APP_DIGITAL_SPACE_URL + v.Video + "/" + v.Thumbnail
							console.log(v.Video)
							return <Video key={v.Video} id={v.Video} title="" datex={v.CreatedAt} src={thumb} onVideoPlay={handleVideoPlay}  />
						}))*/
						
						setVideos(result.map((v) => {
							thumb = process.env.REACT_APP_DIGITAL_SPACE_URL + v.Video + "/" + v.Thumbnail
							return (<Video key={v.Video} id={v.Video} title="" datex={v.CreatedAt} src={thumb} onVideoPlay={handleVideoPlay}  />)
						}))
					}
				}
			)

	}, [])

	return (
		<React.Fragment>
			<Grid container>
        {Videos}
			</Grid>
		</React.Fragment>
	);
}