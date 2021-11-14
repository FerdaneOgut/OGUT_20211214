CREATE USER video_api_user WITH PASSWORD 'root123';
CREATE DATABASE video_db;
GRANT ALL PRIVILEGES ON DATABASE video_db TO video_api_user;