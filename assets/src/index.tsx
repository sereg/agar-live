//ln -s ../node_modules ./node_modules
//npx webpack -w
import * as React from "react";
import * as ReactDOM from "react-dom";
import './scss/main.scss';

import { Hello } from "./components/Hello";

ReactDOM.render(
    <Hello compiler="TypeScript" framework="React" />,
    document.getElementById("example")
);