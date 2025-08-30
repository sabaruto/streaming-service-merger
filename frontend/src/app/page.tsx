import Player from '@/app/audio-player'

export default function Home() {
  return (
    <article className="player">
      <section className="left-sidebar">
        left-sidebar
      </section>
      <section className="main-content">
        main-page
      </section>
      <section className="right-sidebar">
        right-sidebar
      </section>
      <footer>
        <Player />
      </footer>
    </article>
  );
}
