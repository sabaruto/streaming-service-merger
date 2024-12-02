# Streaming Service Merger

This project sets out to create an app that allow users to merge
multiple streaming platforms, including files you uploaded, into
one website.

## Objectives

- Login to Tidal and Spotify, connecting to your own account
- Play music from the platforms and locally uploaded files
- Create playlists from all different platforms

### Stretch Goal

- Send music recommendations based on music played

## User flow

```mermaid
stateDiagram
    [*] --> local_login
    local_login --> streaming_platforms_login
    streaming_platforms_login --> main_page
    main_page --> streaming_platforms_login

    state main_page {
        playlists --> individual_playlist
        individual_playlist --> song_details
        playlists --> song_details
    }

    main_page --> full_screen
    full_screen --> main_page
    main_page --> new_song

    state new_song {
        upload_songs --> add_metadata
    }

    new_song --> main_page
```
