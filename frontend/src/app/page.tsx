import Player from '@/app/audio-player'

export default function Home() {
  return (
    <div className="player">
      <article className="left-sidebar">
        left-sidebar
      </article>
      <article className="main-content">
        main-page
      </article>
      <article className="right-sidebar">
        right-sidebar
      </article>
      <footer>
	  	<article>
		player
        <Player />
		</article>
      </footer>
    </div>
  );
}
