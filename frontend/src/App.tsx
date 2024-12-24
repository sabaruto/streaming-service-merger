import { createRoot } from "react-dom/client";
import Main from "./boundaries/main";

const element = document.getElementById("root");
if (element === null) {
    throw new Error("Cannot find element with root id");
}

createRoot(element).render(<Main />);
