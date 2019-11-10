//ln -s ../node_modules ./node_modules
//npx webpack -w
//https://icomoon.io/app/#/select
import * as ReactDOM from "react-dom";
import * as React from 'react';
import './scss/main.scss';
import Go from './wasm_exec.js';
import {Status} from './const/Const';

import {ControlPanel} from "./components/ControlPanel";

interface AppProps {
}

interface AppState {
    status: Status;
    mod: any;
    inst: any;
}

class App extends React.Component<AppProps, AppState> {

    constructor(props: AppProps) {
        super(props);
        this.state = {
            status: Status.stop,
            mod: "",
            inst: ""
        }
    }

    async componentDidMount() {
        console.log("componentDidMount")
        const go = new Go();
        let {instance, module} = await WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject)
        await go.run(instance)
        this.setState({
            mod: module,
            inst: instance
        })
        console.log("componentDidMount 2")
    }

    cycle = () => {
        if (this.state.status === Status.stop) {
            return
        }
        window.cycle()
        window.requestAnimationFrame(() => {
            this.cycle()
        });
    }

    changeState = () => {
        let newState = Status.playing
        if (newState === this.state.status) {
            newState = Status.stop
        }
        this.setState({
            status: newState
        })
        window.requestAnimationFrame(() => {
            this.cycle()
        });
    }

    render() {
        console.log("render")
        return (
            <div className="row">
                <div className="col-3">
                    <ControlPanel changes={this.changeState} status={this.state.status}/>
                </div>
                <div className="col-9" id="box"/>
            </div>
        )
    }
}

ReactDOM.render(<App/>, document.querySelector('#app'));