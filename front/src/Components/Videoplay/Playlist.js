import React, {useEffect, useState } from 'react';
import Nextvidoes from './Nextvideos'
import Cookies from 'universal-cookie';
import InfiniteScroll from "react-infinite-scroll-component"; // https://github.com/ankeetmaini/react-infinite-scroll-component#readme


/* Token Cookie */
const cookies = new Cookies();

export default function Playlist(props) {
  
  const [Videos, setVideos] = useState("")
  
  const handleVideoPlay = (id) => {
    props.onVideoPlayMain(id);
  };
  
  useEffect(() => {

    fetch(process.env.REACT_APP_API_URL + "/videos", {
      headers: {
        Authorization: `Bearer ` + cookies.get('token'),
      },
    })
      .then(res => res.json())
      .then(
        (result) => {
          if (result != null) {
            let thumb = ""
            setVideos(result.map((v) => {
              thumb = process.env.REACT_APP_DIGITAL_SPACE_URL + v.Video + "/" + v.Thumbnail
              return <Nextvidoes onVideoPlay={handleVideoPlay} img={thumb}  key={v.Video} id={v.Video} date={v.CreatedAt}  />
            }))
          }
        }
      )

  }, [props])
  
  const fetchMoreData = () => {
    fetch(process.env.REACT_APP_API_URL + "/videos", {
      headers: {
        Authorization: `Bearer ` + cookies.get('token'),
      },
    })
      .then(res => res.json())
      .then(
        (result) => {
          if (result != null) {
            let thumb = ""
            setVideos([...Videos,result.map((v) => {
              thumb = process.env.REACT_APP_DIGITAL_SPACE_URL + v.Video + "/" + v.Thumbnail
              return <Nextvidoes onVideoPlay={handleVideoPlay} img={thumb}  key={v.Video} id={v.Video} date={v.CreatedAt}  />
            })])
          }
        }
      )
  };

  return (
    <React.Fragment>
      <InfiniteScroll
          dataLength={Videos.length}
          next={fetchMoreData}
          hasMore={true}
          loader={<h4>Loading...</h4>}
        >
        {Videos}
      </InfiniteScroll>
    </React.Fragment>
  );

}