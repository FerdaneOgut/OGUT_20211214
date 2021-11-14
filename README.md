## VIDEO UPLOADER APP

The repository contains both frontend and backend for video uploading. Backend is written with Go. Frontend is written with React/Typescript.
#### Backend
The api has 4 endpoints. it uses Postgres for data storage. It will seed the categories when the app starts.

- GET /categories to get categories
- POST /video to save video. It saves the video file inside the app folder "data/videos/{key}" with it is unique key. 
It uses ffmpeg utility to create thumbnails images of the video. It takes the first image starting from 1 sec. Saves the images to both "data/images/{key}" and database with the specified sizes(64x64, 128x128, 256x256).
- GET /video/:id to serve the video. It gets the video by the id and serves the video from the path app stores.
- GET /video/ to get all videos from db.

#### Frontend 
The ui has 2 pages inside pages folder. HomePage to list the videos. VideoUpload page to upload videos. 

## Instructions
- Make sure to stop postgres if you have running on your computer before running the docker-compose. 
- RUN docker-compose up to run both frontend and backend together.
```sh
$ docker-compose up
```

- You can run the api and ui seperately if you prefer. The instructions are below; <br/>
```sh
#for backend
#make sure to change .env folder for postgres connection.
$ go run main.go
```
```sh
#for frontend
$ npm install
$ npm start
```
