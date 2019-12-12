//ln -s ../node_modules ./node_modules
//npx webpack -w
//https://icomoon.io/app/#/select
import * as ReactDOM from "react-dom";
import * as React from 'react';
import {init, Universe} from './wasm.js';
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
    action: Universe,
}

class App extends React.Component<AppProps, AppState> {

    constructor(props: AppProps) {
        super(props);
        this.state = {
            tmpElement: "",
            action: new Universe(),
            status: Status.stop,
            countAnimal: 5,
            countPlant: 50,
            selectedElement: {
                id: 0,
                size: 0,
                type: "string"
            }
        }
    }

    async componentDidMount() {
        await init();
        await this.state.action.cycle();
    }

    cycle = () => {
        if (this.state.status === Status.stop) {
            return
        }
        this.state.action.cycle();
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
        await this.state.action.restart();
    };

    async export() {
        let text = await this.state.action.export();
        const element = document.createElement('a');
        element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
        element.setAttribute('download', "export.json");
        element.style.display = 'none';
        document.body.appendChild(element);
        element.click();
        document.body.removeChild(element);
    }

    async import(e: any) {
        const reader = new FileReader();
        reader.onload = event => {
            // @ts-ignore
            const text: string = reader.result;
            this.state.action.import(text);
        };
        reader.onerror = (e) => {
            console.error(e)
        };
        reader.readAsText(e.target.files[0])
    }

    changeCount = (e: any) => {
        const target = e.target;
        const name = target.name;
        let val = target.value * 1;
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

    generate = async () => {
        await this.state.action.generate(this.state.countAnimal, this.state.countPlant);
    };

    backward = async () => {
        await this.state.action.backward();
    };

    setSize = async (e: any) => {
        console.log(JSON.stringify(this.state.selectedElement));
        await this.state.action.setSize(JSON.stringify(this.state.selectedElement));
    };

    moveStart = async (e: any) => {
        let el = await this.state.action.changePosition(e.nativeEvent.offsetX, e.nativeEvent.offsetY);
        this.setState({
            tmpElement: el
        });
        // @ts-ignore
        if (el != "") {
            // @ts-ignore
            let selectedEl: elInfo = JSON.parse(el);
            this.setState({
                selectedElement: {
                    type: selectedEl.Type,
                    id: selectedEl.El.ID,
                    size: selectedEl.El.Size
                }
            });
        }
    };

    moveEnd = async (e: any) => {
        const data = this.state.tmpElement;
        if (data == "") {
            return
        }
        let el = await this.state.action.addFromJSON(data, e.nativeEvent.offsetX, e.nativeEvent.offsetY);
        // @ts-ignore
        if (el != "") {
            let selectedEl: elInfo = JSON.parse(el);
            this.setState({
                selectedElement: {
                    type: selectedEl.Type,
                    id: selectedEl.El.ID,
                    size: selectedEl.El.Size
                }
            });
        }
    };

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