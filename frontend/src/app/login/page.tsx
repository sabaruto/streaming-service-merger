function register() {
  return <>
    <section>
      <h2>
        Register
      </h2>
      <form>
        <div>
          <label>Username</label>
          <input />
        </div>
        <div>
          <label>Password</label>
          <input />
        </div>
        <div>
          <label>Verify Password</label>
          <input />
        </div>
      </form>
      <footer>
        <button>
          <h5>Register</h5>
        </button>
      </footer>
    </section>
    <footer>
      <small>Already have an account? <a>log in</a></small>
    </footer>
  </>
}

function serviceLogin() {
  return <>
    <section>
      <h2>
        Service Login
      </h2>
      <div>
        <button>
          <h6>Tidal Login</h6>
        </button>
        <button>
          <h6>Spotify Login</h6>
          <a href="/api/auth/signin">Sign in</a>
        </button>
      </div>
    </section>
  </>
}

export default function Home() {
  const showLogin = false

  // useEffect(() => {
  // init_tidal_sdk()
  //   .catch(console.error)
  //   .then(
  //     () => search("Eminem")
  //       .catch(console.error)
  //   )
  // }, [])

  return (
    <article className="form-container">
      <header>
        <h1>
          Music Manager
        </h1>
      </header>
      {showLogin ? register() : serviceLogin()}
    </article>
  );
}
