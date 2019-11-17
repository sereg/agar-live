//ln -s ../node_modules ./node_modules
//npx webpack -w
//https://icomoon.io/app/#/select
import * as ReactDOM from "react-dom";
import * as React from 'react';
import Go from './wasm_exec.js';
import {Animal, Plant, Size, Status} from './const/Const';

import {ControlPanel} from "./components/ControlPanel";

interface AppProps {
}

interface el {
    id: number,
    size: number,
    type: string
}

interface elInfo {
    Type: string,
    El: {
        ID: number,
        Size: number
    }
}

interface AppState {
    status: Status;
    countAnimal: number,
    countPlant: number,
    selectedElement: el,
    tmpElement: string,
}

class App extends React.Component<AppProps, AppState> {

    constructor(props: AppProps) {
        super(props);
        this.state = {
            status: Status.stop,
            countAnimal: 5,
            countPlant: 50,
            selectedElement: {
                id: -1,
                size: 0,
                type: ""
            },
            tmpElement: "",
        }
    }

    async componentDidMount() {
        const go = new Go();
        let {instance, module} = await WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject);
        await go.run(instance);
        await window.cycle();
    }

    cycle = () => {
        if (this.state.status === Status.stop) {
            return
        }
        window.cycle();
        window.requestAnimationFrame(() => {
            this.cycle()
        });
    };

    changeState = () => {
        let newState = Status.playing;
        if (newState === this.state.status) {
            newState = Status.stop
        }
        this.setState({
            status: newState
        });
        window.requestAnimationFrame(() => {
            this.cycle()
        });
    };

    async restart() {
        await window.restart();
    };

    async export() {
        let text = await window.export()
        const element = document.createElement('a');
        element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
        element.setAttribute('download', "export.json");
        element.style.display = 'none';
        document.body.appendChild(element);
        element.click();
        document.body.removeChild(element);
    }

    async import(e: any) {
        const reader = new FileReader()
        reader.onload = event => {
            const text = reader.result
            window.import(text)
        }
        reader.onerror = (e) => {
            console.error(e)
        }
        reader.readAsText(e.target.files[0])
    }

    changeCount = (e: any) => {
        const target = e.target
        const name = target.name
        let val = target.value * 1
        if (name == Animal) {
            this.setState({
                countAnimal: val
            })
        }
        if (name == Plant) {
            this.setState({
                countPlant: val
            })
        }
        if (name == Size) {
            this.setState({
                selectedElement: {
                    id: this.state.selectedElement.id,
                    type: this.state.selectedElement.type,
                    size: val
                }
            })
        }
    };

    generate = () => {
        window.generate(this.state.countAnimal, this.state.countPlant);
    };

    backward = () => {
        window.backward();
    };

    setSize = async (e: any) => {
        console.log(JSON.stringify(this.state.selectedElement))
        await window.setSize(JSON.stringify(this.state.selectedElement))
    }

    moveStart = async (e: any) => {
        let el = await window.changePosition(e.nativeEvent.offsetX, e.nativeEvent.offsetY)
        this.setState({
            tmpElement: el,
        });
        if (el != "") {
            let selectedEl:elInfo = JSON.parse(el)
            this.setState({
                selectedElement: {
                    type: selectedEl.Type,
                    id: selectedEl.El.ID,
                    size: selectedEl.El.Size
                }
            });
        }
    }

    moveEnd = async (e: any) => {
        const data = this.state.tmpElement
        if (data == "") {
            return
        }
        let el = await window.addFromJSON(data, e.nativeEvent.offsetX, e.nativeEvent.offsetY)
        if (el != "") {
            let selectedEl:elInfo = JSON.parse(el)
            this.setState({
                selectedElement: {
                    type: selectedEl.Type,
                    id: selectedEl.El.ID,
                    size: selectedEl.El.Size
                }
            });
        }
    }

    render() {
        return (
            <div className="row">
                <div className="col-3">
                    <ControlPanel
                        changeCount={this.changeCount}
                        generate={this.generate}
                        changes={this.changeState}
                        restart={this.restart}
                        export={this.export}
                        import={this.import}
                        backward={this.backward}
                        setSize={this.setSize}
                        status={this.state.status}
                        countAnimal={this.state.countAnimal}
                        countPlant={this.state.countPlant}
                        sizeElement={this.state.selectedElement.size}
                    />
                </div>
                <div
                    className="col-9"
                    onMouseDown={this.moveStart}
                    onMouseUp={this.moveEnd}
                    id="box"
                />
            </div>
        )
    };
}

ReactDOM.render(<App/>, document.querySelector('#app'));