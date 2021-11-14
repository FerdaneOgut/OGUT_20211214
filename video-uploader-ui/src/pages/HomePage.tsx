import { useCallback, useEffect, useMemo, useState } from "react";
import { Button, Container, Stack, Table } from "react-bootstrap";
import { useNavigate } from 'react-router-dom';
import { Video } from "../data/models/Video";
import VideoService from "../data/services/videoService";
import HtmlVideo from "../components/HtmlVideo";
import Image from "../components/Image";
import Modal from "../components/Modal";

function HomePage(props: any) {
  const [videos, setVideos] = useState<Video[]>([]);
  const [selectedVideo, setSelectedVideo] = useState<Video>();
  const [showVideo, setShowVideo] = useState(false);
  const navigate = useNavigate();
  const uploadVideo = useCallback(() => {
    navigate("/video")
  }, [navigate]);

  useEffect(() => {
    //handle error
    VideoService.getAll()
      .then(r => { setVideos(r.data); });
  }, []);
  const onImageClick = useCallback((v: Video) => {
    setShowVideo(true);
    setSelectedVideo(v);
  }, []);
  const close = () => {
    setShowVideo(false);
  }
  const getImage = useCallback((v: Video) => <Image video={v} onImageClick={onImageClick} />, [onImageClick]);

  const renderVideos = useMemo(() =>
    videos.map((v, i) => <tr key={i}>
      <td>{v.id}</td>
      <td>{v.title}</td>
      <td>{getImage(v)}</td>
      <td >{v.Category?.name}</td>
    </tr>)
    , [videos, getImage]);

  return <Container>

    <h3 style={{marginTop:"15px"}}>Videos</h3>

    <Stack>
      <Stack style={{ alignItems: "flex-start" }}>
        <Button variant="primary" onClick={uploadVideo}>Upload New Video</Button>
      </Stack>
      <Table style={{marginTop:"15px"}} striped bordered hover >
        <thead>
          <tr>
            <th>#</th>
            <th>Video Title</th>
            <th>Video Thumbnail</th>
            <th>Video Category</th>
          </tr>
        </thead>
        <tbody>
          {renderVideos}
        </tbody>
      </Table>
    </Stack>

    <Modal show={showVideo} close={close}>
      {selectedVideo && <HtmlVideo video={selectedVideo} />}
    </Modal>

  </Container>
}
export default HomePage;

