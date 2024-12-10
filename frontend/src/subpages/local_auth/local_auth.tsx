import React, { useRef, useState } from "react";

function LocalAuth(props: { setAuthComplete: () => void }) {
    const [isRegister, setShowRegister] = useState(true);
    const formRef = useRef<HTMLFormElement>(null);
    const usernameRef = useRef<HTMLInputElement>(null);
    const passwordRef = useRef<HTMLInputElement>(null);

    const title = isRegister ? "Register" : "Log In";
    const footerLabel = isRegister ? "Already have user?" : "New User?";
    const footerButton = isRegister ? "Log in" : "Register";

    function validateCredentials(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        if (!formRef.current || !usernameRef.current || !passwordRef.current) {
            throw new Error("form not fully initialised");
        }

        const username = usernameRef.current.value;
        const password = passwordRef.current.value;

        if (username.length < 1 || password.length < 1) {
            console.error("Username is empty");
            return;
        }

        console.log("Username", username);
        console.log("Password", password);

        if (isRegister) {
            setShowRegister(false);
        } else {
            props.setAuthComplete();
        }

        formRef.current.reset();
    }

    return (
        <>
            <header>
                <h1>{title}</h1>
            </header>
            <form
                ref={formRef}
                onSubmit={(event: React.FormEvent<HTMLFormElement>) => {
                    validateCredentials(event);
                }}
            >
                <label>
                    Username
                    <input ref={usernameRef} />
                </label>
                <label>
                    Password
                    <input
                        type="password"
                        ref={passwordRef}
                    />
                </label>
                <button type="submit">Submit</button>
            </form>
            <footer>
                {footerLabel}
                <button
                    onClick={() => {
                        formRef.current?.reset();
                        setShowRegister(!isRegister);
                    }}
                >
                    {footerButton}
                </button>
            </footer>
        </>
    );
}

export default LocalAuth;
