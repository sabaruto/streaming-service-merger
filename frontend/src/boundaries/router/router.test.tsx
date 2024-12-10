import { test } from "vitest";
import { render } from "vitest-browser-react";
import Router from "./router";

test("renders", function () {
    render(<Router />);
});
