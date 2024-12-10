import { FallbackProps } from "react-error-boundary";
function DefaultFallbackComponent(props: FallbackProps) {
    const error = props.error as Error;

    return (
        <div
            id="page"
            className="center"
        >
            <div className="main-card">
                <header>
                    <h1>{error.name}</h1>
                </header>
                <span>{error.message}</span>
            </div>
        </div>
    );
}

export default DefaultFallbackComponent;
