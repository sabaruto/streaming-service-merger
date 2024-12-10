import { StrictMode } from "react";
import Router from "../router";
import { ErrorBoundary } from "react-error-boundary";
import DefaultErrorBoundary from "../../pages/error";

function Main() {
    return (
        <StrictMode>
            <ErrorBoundary FallbackComponent={DefaultErrorBoundary}>
                <Router />
            </ErrorBoundary>
        </StrictMode>
    );
}

export default Main;
