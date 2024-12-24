import { useState } from "react";
import LocalAuth from "../../subpages/local_auth";
import ServiceAuth from "../../subpages/service_auth";

function Auth() {
    const [localAuthDone, setLocalAuth] = useState(false);
    fetch("http://localhost:8000/auth/csrf/")
        .then((data: Response) => { return data.json() })
        .then((data) => { console.log(data) })
        .catch((err: unknown) => { console.log(err) })

    function setLocalAuthDone() {
        setLocalAuth(true);
    }

    const currentForm = localAuthDone ? <ServiceAuth /> : <LocalAuth setAuthComplete={setLocalAuthDone.bind(Auth)} />;

    return (
        <div
            id="page"
            className="center"
        >
            <div className="main-card">{currentForm}</div>
        </div>
    );
}

export default Auth;
