import Navbar from "../../components/navbar";

import { BrowserRouter, Routes, Route } from "react-router-dom";
import WebPlayer from "../../pages/web_player";
import Auth from "../../pages/auth";
import AddSong from "../../pages/add_song";
import Playlists from "../../subpages/playlists";
import IndividualPlaylist from "../../subpages/individual_playlist";
import SongDetails from "../../subpages/song_details";
import SearchResults from "../../subpages/search_results";
import Title from "../title";

function Router() {
    return (
        <BrowserRouter>
            <header>
                <Title />
                <Navbar />
            </header>
            <Routes>
                <Route
                    path="/"
                    element={<Auth />}
                />
                <Route
                    path="/player"
                    element={<WebPlayer />}
                >
                    <Route
                        path="playlists"
                        element={<Playlists />}
                    />
                    <Route
                        path="playlist/:id"
                        element={<IndividualPlaylist />}
                    />
                    <Route
                        path="song/:id"
                        element={<SongDetails />}
                    />
                    <Route
                        path="search-results"
                        element={<SearchResults />}
                    />
                </Route>
                <Route
                    path="/add-song"
                    element={<AddSong />}
                />
            </Routes>
        </BrowserRouter>
    );
}

export default Router;
