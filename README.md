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
    [*] --> login

    state login {
        [*] --> local_login
        [*] --> streaming_platforms_login
        local_login --> streaming_platforms_login: Log in
        local_login --> local_signup: New user
        local_signup --> streaming_platforms_login
        streaming_platforms_login --> [*]
    }
    
    login --> main_page

    state main_page {
        [*] --> playlists
        playlists --> individual_playlist: Choose playlist
        individual_playlist --> song_details: Choose song
        search_results --> playlists: Choose from list
        search_results --> individual_playlist: Choose from list
        search_results --> song_details: Choose from list
        song_details --> individual_playlist: Associated playlists
        [*] --> search_results: Search bar
        state states {
            album_view
            info_view
        }
    }

    main_page --> add_song

    state add_song {
        upload_songs --> add_metadata
    }

    add_song --> main_page
    [*] --> error_page: An error has occured
```
