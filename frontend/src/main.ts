import "./style.css";
import App from "./App.svelte";
import "./tailwind.css";
import { mount } from "svelte";

const app = mount(App, {
    target: document.getElementById("app"),
});

export default app;
