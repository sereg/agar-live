//ln -s ../node_modules ./node_modules
//npx webpack -w
import * as ReactDOM from "react-dom";
import * as React from 'react';
import './scss/main.scss';
import  Go from "./wasm_exec.js";

import {ControlPanel} from "./components/ControlPanel";

class App extends React.Component {

    constructor(props: any) {
        super(props);
    }

    async componentDidMount() {
        console.log("componentDidMount")
        const go = new Go();
        let { instance, module } = await WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject)
        await go.run(instance)
        this.setState({
            mod: module,
            inst: instance
        })
    }


    render() {
        console.log("render")
        return (
            <div className="row">
                <div className="col-3">
                    <ControlPanel compiler="TypeScript" framework="React"/>
                </div>
                <div className="col-9" id="box"/>
            </div>
        )
    }
}

ReactDOM.render(<App/>, document.getElementById("example"));