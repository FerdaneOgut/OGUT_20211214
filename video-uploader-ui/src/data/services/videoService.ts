import { Video } from "../models/Video";
import http from "./../http.common";

const getAll = () => {
  return http.get<Array<Video>>(`/video`);
};

const get = (id: any) => {
  return http.get<any>(`/video/${id}`);
};

const create = (data: FormData) => {
  return http.post<Video>("/video", data, { headers: { 'Content-Type': 'multipart/form-data' } });
};

const update = (id: any, data: Video) => {
  return http.put<any>(`/video/${id}`, data);
};

const remove = (id: any) => {
  return http.delete<any>(`/video/${id}`);
};




const VideoService = {
  getAll,
  get,
  create,
  update,
  remove,
};

export default VideoService;