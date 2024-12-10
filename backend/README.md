# Database Design

## Arcitecture

```mermaid
architecture-beta
    group api(cloud)[API]

    service db(database)[Database] in api
    service user_db(database)[Database] in api
    service disk1(logos:aws-s3)[S3 Bucket] in api
    service server(server)[Server] in api

    db:L -- R:server
    disk1:T -- B:db
```

## Database Tables

```mermaid
erDiagram
    USER {
        string id pk
        string name
    }

    CREDS {
        string creds pk
        enum stream_type
    }

    SONG {
        string id pk
        string title
    }

    LOCAL_SONG {
        time duration
        string type
        int bitrate
        int channels
        float16 sample_rate
        int sample_size
        time create_date
    }

    STREAMED_SONG {
        enum stream_type
        string song_id
    }

    ALBUM {
        string id pk
        string title
        time duration
        time create_date
    }

    LOCAL_ALBUM {
        string image_location
    }

    ARTIST {
        string id pk
        string name
    }

    LOCAL_ARTIST {
        string bio
        string image_location
    }

    ARTIST_STREAM_ID {
        enum stream_type
        string artist_id
    }

    PLAYLIST {
        string id pk
        string name
        string image_location
    }

    PLAYLIST_STREAM_ID {
        enum stream_type
        string playlist_id
    }

    PLAYLIST 1..0+ PLAYLIST_STREAM_ID : has
    PLAYLIST 0+--0+ SONG : contains

    SONG 1..1 LOCAL_SONG : is
    SONG 1..1 STREAMED_SONG : is
    SONG 0+--0+ ARTIST : has

    USER 1..0+ CREDS : has
    USER 1..0+ PLAYLIST : has
    USER 1..0+ LOCAL_SONG : has
    USER 1..0+ LOCAL_ALBUM : has
    USER 1..0+ LOCAL_SONG : has
    USER 1..0+ LOCAL_ARTIST : has

    ALBUM 1--1+ ARTIST : has
    ALBUM 1--1+ SONG : has
    ALBUM 1..o| LOCAL_ALBUM : is

    ARTIST 1..1+ ARTIST_STREAM_ID : has
    ARTIST 1..o| LOCAL_ARTIST : is
```

## Applications

### Authorisation

Sets up User table and creates views for
- User
- Credentials
- Authorisation

### Collections

Sets up table and create views for
- Playlist
- Album
- Artist
- Playlist Stream ID

### Songs

Sets up Song, LocalSong, StreamedSong and create views for
- Songs
- Local Song
- Streamed Song