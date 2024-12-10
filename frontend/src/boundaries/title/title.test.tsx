import { expect, test } from "vitest";
import { render } from "vitest-browser-react";
import Title from "./title";

test("loads and displays title", function () {
    const screen = render(<Title />);

    const title_element = screen.getByTitle("tab-name").element();
    expect(title_element).toHaveProperty("text", "Music streaming manager");
});
