/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, { enumerable: true, get: getter });
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 			Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 		}
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// create a fake namespace object
/******/ 	// mode & 1: value is a module id, require it
/******/ 	// mode & 2: merge all properties of value into the ns
/******/ 	// mode & 4: return value when already ns object
/******/ 	// mode & 8|1: behave like require
/******/ 	__webpack_require__.t = function(value, mode) {
/******/ 		if(mode & 1) value = __webpack_require__(value);
/******/ 		if(mode & 8) return value;
/******/ 		if((mode & 4) && typeof value === 'object' && value && value.__esModule) return value;
/******/ 		var ns = Object.create(null);
/******/ 		__webpack_require__.r(ns);
/******/ 		Object.defineProperty(ns, 'default', { enumerable: true, value: value });
/******/ 		if(mode & 2 && typeof value != 'string') for(var key in value) __webpack_require__.d(ns, key, function(key) { return value[key]; }.bind(null, key));
/******/ 		return ns;
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "/public/";
/******/
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = "./src/index.tsx");
/******/ })
/************************************************************************/
/******/ ({

/***/ "./src/components/ControlPanel.tsx":
/*!*****************************************!*\
  !*** ./src/components/ControlPanel.tsx ***!
  \*****************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const React = __importStar(__webpack_require__(/*! react */ "react"));
const Const_1 = __webpack_require__(/*! ../const/Const */ "./src/const/Const.ts");
const icon_1 = __webpack_require__(/*! ./icon/icon */ "./src/components/icon/icon.tsx");
exports.ControlPanel = (props) => {
    return (React.createElement("div", null,
        React.createElement("div", { className: "row" },
            React.createElement("div", { className: "col-sm", onClick: props.backward },
                React.createElement(icon_1.Backward, { width: "50px" })),
            React.createElement("div", { className: "col-sm", onClick: props.changes }, props.status === 1 /* stop */ ? (React.createElement(icon_1.Play, { width: "50px" })) : (React.createElement(icon_1.Stop, { width: "50px" }))),
            React.createElement("div", { className: "col-sm" },
                React.createElement(icon_1.Forward, { width: "50px" }))),
        React.createElement("div", { className: "row" },
            React.createElement("div", { className: "col-sm", onClick: props.export },
                React.createElement(icon_1.Save, { width: "50px" })),
            React.createElement("div", { className: "col-sm" },
                React.createElement(icon_1.Load, { width: "50px" }),
                React.createElement("input", { id: "file", type: "file", onChange: props.import })),
            React.createElement("div", { className: "col-sm", onClick: props.restart },
                React.createElement(icon_1.Refresh, { width: "50px" }))),
        React.createElement("div", null,
            React.createElement("fieldset", null,
                React.createElement("legend", null, "Generate"),
                React.createElement("p", null,
                    React.createElement("label", { htmlFor: "input" }, "Count animal"),
                    React.createElement("input", { value: props.countAnimal, onChange: props.changeCount, type: "text", placeholder: "0", name: Const_1.Animal })),
                React.createElement("p", null,
                    React.createElement("label", { htmlFor: "input" }, "Count plant"),
                    React.createElement("input", { value: props.countPlant, onChange: props.changeCount, type: "text", placeholder: "0", name: Const_1.Plant })),
                React.createElement("p", null,
                    React.createElement("button", { type: "submit", onClick: props.generate }, "Generate")))),
        React.createElement("div", null,
            React.createElement("fieldset", null,
                React.createElement("legend", null, "Set size"),
                React.createElement("p", null,
                    React.createElement("label", { htmlFor: "input" }, "Size element"),
                    React.createElement("input", { value: props.sizeElement, onChange: props.changeCount, type: "text", placeholder: "0", name: Const_1.Size })),
                React.createElement("p", null,
                    React.createElement("button", { type: "submit", onClick: props.setSize }, "Set"))))));
};


/***/ }),

/***/ "./src/components/icon/icon.tsx":
/*!**************************************!*\
  !*** ./src/components/icon/icon.tsx ***!
  \**************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const react_1 = __importDefault(__webpack_require__(/*! react */ "react"));
exports.Play = (props) => {
    let width = props.width || "20px";
    return (react_1.default.createElement("svg", { style: { width: width }, viewBox: "0 0 512 512" },
        react_1.default.createElement("path", { d: "M96 64l320 192-320 192z" })));
};
exports.Stop = (props) => {
    let width = props.width || "20px";
    return (react_1.default.createElement("svg", { style: { width: width }, viewBox: "0 0 512 512" },
        react_1.default.createElement("path", { d: "M64 64h160v384h-160zM288 64h160v384h-160z" })));
};
exports.Refresh = (props) => {
    let width = props.width || "20px";
    return (react_1.default.createElement("svg", { style: { width: width }, viewBox: "0 0 512 512" },
        react_1.default.createElement("path", { d: "M512 192h-192l71.765-71.765c-36.265-36.263-84.48-56.235-135.765-56.235s-99.5 19.972-135.765 56.235c-36.263 36.265-56.235 84.48-56.235 135.765s19.972 99.5 56.235 135.765c36.265 36.263 84.48 56.235 135.765 56.235s99.5-19.972 135.764-56.236c3.028-3.027 5.93-6.146 8.728-9.334l48.16 42.141c-46.923 53.583-115.832 87.429-192.652 87.429-141.385 0-256-114.615-256-256s114.615-256 256-256c70.693 0 134.684 28.663 181.008 74.992l74.992-74.992v192z" })));
};
exports.Backward = (props) => {
    let width = props.width || "20px";
    return (react_1.default.createElement("svg", { style: { width: width }, viewBox: "0 0 512 512" },
        react_1.default.createElement("path", { d: "M288 80v160l160-160v352l-160-160v160l-176-176z" })));
};
exports.Forward = (props) => {
    let width = props.width || "20px";
    return (react_1.default.createElement("svg", { style: { width: width }, viewBox: "0 0 512 512" },
        react_1.default.createElement("path", { d: "M256 432v-160l-160 160v-352l160 160v-160l176 176z" })));
};
exports.Save = (props) => {
    let width = props.width || "20px";
    return (react_1.default.createElement("svg", { style: { width: width }, viewBox: "0 0 512 512" },
        react_1.default.createElement("path", { d: "M224 288h64v-128h96l-128-128-128 128h96zM320 216v49.356l146.533 54.644-210.533 78.509-210.533-78.509 146.533-54.644v-49.356l-192 72v128l256 96 256-96v-128z" })));
};
exports.Load = (props) => {
    let width = props.width || "20px";
    return (react_1.default.createElement("svg", { style: { width: width }, viewBox: "0 0 512 512" },
        react_1.default.createElement("path", { d: "M256 288l128-128h-96v-128h-64v128h-96zM372.363 235.636l-35.87 35.871 130.040 48.493-210.533 78.509-210.533-78.509 130.040-48.493-35.871-35.871-139.636 52.364v128l256 96 256-96v-128z" })));
};


/***/ }),

/***/ "./src/const/Const.ts":
/*!****************************!*\
  !*** ./src/const/Const.ts ***!
  \****************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

Object.defineProperty(exports, "__esModule", { value: true });
exports.Animal = "animal";
exports.Plant = "plant";
exports.Size = "size";


/***/ }),

/***/ "./src/index.tsx":
/*!***********************!*\
  !*** ./src/index.tsx ***!
  \***********************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
//ln -s ../node_modules ./node_modules
//npx webpack -w
//https://icomoon.io/app/#/select
const ReactDOM = __importStar(__webpack_require__(/*! react-dom */ "react-dom"));
const React = __importStar(__webpack_require__(/*! react */ "react"));
const wasm_js_1 = __webpack_require__(/*! ./wasm.js */ "./src/wasm.js");
const Const_1 = __webpack_require__(/*! ./const/Const */ "./src/const/Const.ts");
const ControlPanel_1 = __webpack_require__(/*! ./components/ControlPanel */ "./src/components/ControlPanel.tsx");
class App extends React.Component {
    constructor(props) {
        super(props);
        this.cycle = () => {
            if (this.state.status === 1 /* stop */) {
                return;
            }
            this.state.action.cycle();
            window.requestAnimationFrame(() => {
                this.cycle();
            });
        };
        this.changeState = () => {
            let newState = 0 /* playing */;
            if (newState === this.state.status) {
                newState = 1 /* stop */;
            }
            this.setState({
                status: newState
            });
            window.requestAnimationFrame(() => {
                this.cycle();
            });
        };
        this.changeCount = (e) => {
            const target = e.target;
            const name = target.name;
            let val = target.value * 1;
            if (name == Const_1.Animal) {
                this.setState({
                    countAnimal: val
                });
            }
            if (name == Const_1.Plant) {
                this.setState({
                    countPlant: val
                });
            }
            if (name == Const_1.Size) {
                this.setState({
                    selectedElement: {
                        id: this.state.selectedElement.id,
                        type: this.state.selectedElement.type,
                        size: val
                    }
                });
            }
        };
        this.generate = () => __awaiter(this, void 0, void 0, function* () {
            yield this.state.action.generate(this.state.countAnimal, this.state.countPlant);
        });
        this.backward = () => __awaiter(this, void 0, void 0, function* () {
            yield this.state.action.backward();
        });
        this.setSize = (e) => __awaiter(this, void 0, void 0, function* () {
            console.log(JSON.stringify(this.state.selectedElement));
            yield this.state.action.setSize(JSON.stringify(this.state.selectedElement));
        });
        this.moveStart = (e) => __awaiter(this, void 0, void 0, function* () {
            let el = yield this.state.action.changePosition(e.nativeEvent.offsetX, e.nativeEvent.offsetY);
            this.setState({
                tmpElement: el
            });
            // @ts-ignore
            if (el != "") {
                // @ts-ignore
                let selectedEl = JSON.parse(el);
                this.setState({
                    selectedElement: {
                        type: selectedEl.Type,
                        id: selectedEl.El.ID,
                        size: selectedEl.El.Size
                    }
                });
            }
        });
        this.moveEnd = (e) => __awaiter(this, void 0, void 0, function* () {
            const data = this.state.tmpElement;
            if (data == "") {
                return;
            }
            let el = yield this.state.action.addFromJSON(data, e.nativeEvent.offsetX, e.nativeEvent.offsetY);
            // @ts-ignore
            if (el != "") {
                let selectedEl = JSON.parse(el);
                this.setState({
                    selectedElement: {
                        type: selectedEl.Type,
                        id: selectedEl.El.ID,
                        size: selectedEl.El.Size
                    }
                });
            }
        });
        this.state = {
            tmpElement: "",
            action: new wasm_js_1.Universe(),
            status: 1 /* stop */,
            countAnimal: 5,
            countPlant: 50,
            selectedElement: {
                id: 0,
                size: 0,
                type: "string"
            }
        };
    }
    componentDidMount() {
        return __awaiter(this, void 0, void 0, function* () {
            yield wasm_js_1.init();
            yield this.state.action.cycle();
        });
    }
    restart() {
        return __awaiter(this, void 0, void 0, function* () {
            yield this.state.action.restart();
        });
    }
    ;
    export() {
        return __awaiter(this, void 0, void 0, function* () {
            let text = yield this.state.action.export();
            const element = document.createElement('a');
            element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
            element.setAttribute('download', "export.json");
            element.style.display = 'none';
            document.body.appendChild(element);
            element.click();
            document.body.removeChild(element);
        });
    }
    import(e) {
        return __awaiter(this, void 0, void 0, function* () {
            const reader = new FileReader();
            reader.onload = event => {
                // @ts-ignore
                const text = reader.result;
                this.state.action.import(text);
            };
            reader.onerror = (e) => {
                console.error(e);
            };
            reader.readAsText(e.target.files[0]);
        });
    }
    render() {
        return (React.createElement("div", { className: "row" },
            React.createElement("div", { className: "col-3" },
                React.createElement(ControlPanel_1.ControlPanel, { changeCount: this.changeCount, generate: this.generate, changes: this.changeState, restart: this.restart, export: this.export, import: this.import, backward: this.backward, setSize: this.setSize, status: this.state.status, countAnimal: this.state.countAnimal, countPlant: this.state.countPlant, sizeElement: this.state.selectedElement.size })),
            React.createElement("div", { className: "col-9", onMouseDown: this.moveStart, onMouseUp: this.moveEnd, id: "box" })));
    }
    ;
}
ReactDOM.render(React.createElement(App, null), document.querySelector('#app'));


/***/ }),

/***/ "./src/wasm.js":
/*!*********************!*\
  !*** ./src/wasm.js ***!
  \*********************/
/*! exports provided: init, Universe */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "init", function() { return init; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "Universe", function() { return Universe; });
// import Go from './wasm_exec.js';


async function  init() {
    const go = new Go();
    // let result = await WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject);
    // await go.run(result.instance);
    let {instance, module} = await WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject);
    await go.run(instance);
}

class Universe {

    async cycle() {
        await window.cycle();
    }

    async restart() {
        await window.restart();
    }
    /**
     * @returns {string}
     */
    async export() {
        return await window.export();
    }
    /**
     * @param {string} text
     */
    async import(text) {
        await window.import(text);
    }
    /**
     * @param {number} countAnimal
     * @param {number} countPlant
     */
    async generate(countAnimal, countPlant) {
        await window.generate(countAnimal, countPlant);
    }
    async backward() {
        await window.backward();
    }
    /**
     * @param {string} params
     */
    async setSize(params) {
        await window.setSize(JSON.stringify(this.state.selectedElement));
    }
    /**
     * @param {number} x
     * @param {number} y
     * @returns {string}
     */
    async changePosition(x, y) {
        return window.changePosition(x, y);
    }
    /**
     * @param {string} data
     * @param {number} x
     * @param {number} y
     * @returns {string}
     */
    async addFromJSON(data, x, y) {
        await window.addFromJSON(data, x, y);
    }
}

/***/ }),

/***/ "react":
/*!************************!*\
  !*** external "React" ***!
  \************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = React;

/***/ }),

/***/ "react-dom":
/*!***************************!*\
  !*** external "ReactDOM" ***!
  \***************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = ReactDOM;

/***/ })

/******/ });
//# sourceMappingURL=main.js.map