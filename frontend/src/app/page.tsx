export default function Home() {
  return (
    <article className="form-container">
      <header>
        <h2>
          Music Manager
        </h2>
      </header>
      <section>
        <h3>
          Register
        </h3>
        <form>
          <div>
            <label>Username</label>
            <input/>
          </div>
          <div>
            <label>Password</label>
            <input/>
          </div>
          <div>
            <label>Verify Password</label>
            <input/>
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
    </article>
  );
}
